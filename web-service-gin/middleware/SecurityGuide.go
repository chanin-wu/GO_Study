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
			永远不要将 AllowOrigins: []string{"*"} 与 AllowCredentials: true 一起使用。这会告诉浏览器任何站点都可以向你的 API 发送认证请求。

		CSRF 保护
			跨站请求伪造会诱骗已认证用户的浏览器向你的应用发送不需要的请求。
			任何依赖 cookie 进行认证的状态变更端点（POST、PUT、DELETE）都需要 CSRF 保护。
			CSRF 保护对于使用基于 cookie 认证的应用至关重要。
			仅依赖 Authorization 头（例如 Bearer 令牌）的 API 不会受到 CSRF 攻击，因为浏览器不会自动附加这些头。

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
	type client struct {
		limiter *rate.Limiter
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		if _, exists := clients[ip]; !exists {
			// Allow 10 requests per second with a burst of 20
			clients[ip] = &client{limiter: rate.NewLimiter(10, 20)}
		}
		cl := clients[ip]
		mu.Unlock()

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
