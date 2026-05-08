package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"golang.org/x/time/rate"
)

/*
	安全最佳实践
		Web 应用是攻击者的主要目标。处理用户输入、存储数据或在反向代理后运行的 Gin 应用在进入生产环境之前需要有意的安全配置。
		本指南涵盖了最重要的防御措施，并展示如何使用 Gin 中间件和标准 Go 库来应用每项措施。

		安全是分层的。本列表中没有任何单一技术本身就足够。应用所有相关部分以构建纵深防御。

		CORS 配置
			跨源资源共享（CORS）控制哪些外部域可以向你的 API 发出请求。配置错误的 CORS 可以允许恶意网站代表已认证用户读取你服务器的响应。
				AllowCredentials：
					含义：设置为 true 表示允许在跨源请求中携带身份凭证（如 cookies、HTTP 认证及客户端 SSL 证书等）。如果设置为 false，即使源站在 AllowOrigins 列表中，带有身份凭证的请求也会被阻止。需要注意的是，
					当 AllowCredentials 为 true 时，AllowOrigins 不能设置为 *，必须指定具体的源站。
			永远不要将 AllowOrigins: []string{"*"} 与 AllowCredentials: true 一起使用。这会告诉浏览器任何站点都可以向你的 API 发送认证请求。

		CSRF 保护
			跨站请求伪造会诱骗已认证用户的浏览器向你的应用发送不需要的请求。
			任何依赖 cookie 进行认证的状态变更端点（POST、PUT、DELETE）都需要 CSRF 保护。
			CSRF 保护对于使用基于 cookie 认证的应用至关重要。
			仅依赖 Authorization 头（例如 Bearer 令牌）的 API 不会受到 CSRF 攻击，因为浏览器不会自动附加这些头。

				创建了一个基于 Cookie 的会话存储。cookie.NewStore 函数接受一个字节切片作为密钥，用于加密和解密存储在 Cookie 中的会话数据。这个密钥必须保密且足够复杂，以防止攻击者破解会话数据。
				将刚刚创建的会话存储与名为 "mysession" 的会话关联起来，并将这个会话中间件应用到整个 Gin 路由器 router 上。每个请求经过这个中间件时，都会初始化或恢复与 "mysession" 相关的会话。会话对于 CSRF 防护很重要，因为 CSRF 令牌通常与会话关联。
				Secret：这是 CSRF 令牌生成和验证所使用的密钥。与会话密钥类似，这个密钥也必须保密且唯一。不同的应用程序应该使用不同的 CSRF 密钥，以防止一个应用程序的 CSRF 漏洞影响到其他应用程序。
				ErrorFunc：这是一个回调函数，当 CSRF 令牌验证失败时会被调用。在这个函数中，它向客户端返回一个 HTTP 403 状态码（表示禁止访问）和一条简单的错误信息 "CSRF token mismatch"，然后调用 c.Abort() 停止请求的进一步处理。
				这确保了如果检测到 CSRF 攻击，服务器不会继续处理该请求，从而保护了服务器资源。

		限流
			限流可以防止滥用、暴力攻击和资源耗尽。
			下面的示例将限流器存储在内存映射中。在生产环境中，你应该添加定期清理过期条目的机制，如果你运行多个应用实例，请考虑使用分布式限流器（例如基于 Redis）。

		输入验证和 SQL 注入防护
			始终使用 Gin 的模型绑定和结构体标签来验证和绑定输入。永远不要通过拼接用户输入来构造 SQL 查询。

		使用参数化查询
			row := db.QueryRow("SELECT id FROM users WHERE username = $1", username)
			这适用于每个数据库库。无论你使用 database/sql、GORM、sqlx 还是其他 ORM，都要使用参数占位符而不是字符串拼接。

		输入验证和参数化查询是你对抗注入攻击的两个最重要的防御手段。单独使用任何一个都不够——请同时使用。

		XSS 防护
			跨站脚本（XSS）发生在攻击者注入恶意脚本并在其他用户的浏览器中执行时。在多个层面进行防御。

		转义 HTML 输出
			渲染 HTML 模板时，Go 的 html/template 包默认转义输出。如果你将用户提供的数据作为 JSON 返回，请确保正确设置 Content-Type 头，以便浏览器不会将 JSON 解释为 HTML。

		安全头中间件

		可信代理
			当你的应用在反向代理或负载均衡器后运行时，你必须告诉 Gin 信任哪些代理。
			没有此配置，攻击者可以伪造 X-Forwarded-For 头以绕过基于 IP 的访问控制和限流。
			Trust only your known proxy addresses
			router.SetTrustedProxies([]string{"10.0.0.1", "192.168.1.0/24"})

		HTTPS 和 TLS
			所有生产环境的 Gin 应用都应该通过 HTTPS 提供流量。Gin 支持通过 Let’s Encrypt 自动获取 TLS 证书。
			始终将 HTTPS 与 Strict-Transport-Security 头（HSTS）结合使用，以防止协议降级攻击。一旦设置了 HSTS 头，浏览器将拒绝通过纯 HTTP 连接。
*/

func RateLimiter() gin.HandlerFunc {
	// rate.Limiter 是 Go 标准库 golang.org/x/time/rate 包中的一个类型，用于实现速率限制逻辑
	type client struct {
		limiter *rate.Limiter
	}

	// mu 是一个 sync.Mutex 类型的互斥锁，用于保护对 clients 映射的并发访问。因为多个请求可能同时尝试访问和修改 clients 映射，所以需要使用互斥锁来确保数据的一致性和线程安全
	// clients 是一个映射，键是客户端的 IP 地址（字符串类型），值是指向 client 结构体的指针。这个映射用于存储每个客户端的速率限制器。
	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	return func(c *gin.Context) {
		// 获取客户端 IP 地址
		ip := c.ClientIP()

		// 首先，通过 mu.Lock() 锁定互斥锁，以确保对 clients 映射的操作是线程安全的。
		// 然后检查 clients 映射中是否已经存在该 IP 地址对应的客户端记录。
		// 如果不存在，则使用 rate.NewLimiter(10, 20) 创建一个新的速率限制器。
		// 这里的 10 表示每秒允许 10 个请求，20 表示允许的突发请求数（即瞬间可以处理的最大请求数）。
		// 最后，解锁互斥锁 mu.Unlock()
		mu.Lock()
		if _, exists := clients[ip]; !exists {
			// Allow 10 requests per second with a burst of 20
			clients[ip] = &client{limiter: rate.NewLimiter(10, 20)}
		}
		cl := clients[ip]
		mu.Unlock()

		// 调用 cl.limiter.Allow() 方法来检查当前请求是否被允许通过速率限制。
		// 如果不允许（即请求频率超过了设定的限制），则使用 c.AbortWithStatusJSON 方法终止请求的处理，并向客户端返回一个 HTTP 429 状态码（表示请求过多）和一个包含错误信息的 JSON 响应
		if !cl.limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			return
		}

		c.Next()
	}
}

func SecurityGuide(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "SecurityGuide",
	})
}
