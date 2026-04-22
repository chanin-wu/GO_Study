package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/*
	ShouldBind 和 ShouldBindBodyWith 都用于将 HTTP 请求中的数据绑定到 Go 语言的结构体上
	ShouldBind 是一个通用的数据绑定方法，它会根据请求的 Content - Type 头信息自动选择合适的绑定器来处理请求数据。它支持多种常见的数据格式，如 JSON、XML、表单数据等。
	ShouldBind 根据请求的 Content - Type 来决定如何将请求体数据绑定到 RequestData 结构体。
	如果是 application/json 类型，它会按照 JSON 格式解析；如果是 application/x-www-form-urlencoded，则按表单数据解析。

	ShouldBindBodyWith 允许你显式指定使用哪种绑定器来处理请求体数据。
	ShouldBindBodyWith 明确指定使用 binding.JSON 绑定器，所以无论请求的 Content - Type 是什么，都会尝试按照 JSON 格式来绑定数据。
*/

/*
	区别总结

	绑定方式：
		ShouldBind 自动根据 Content - Type 选择绑定器，更加灵活通用，适用于多种数据格式的自动处理场景。
		ShouldBindBodyWith 则需要手动指定绑定器，更精确地控制数据绑定过程，适用于你确定数据格式且不希望受 Content - Type 影响的场景。
	应用场景：
		如果你的 API 需要处理多种不同格式（JSON、XML、表单等）的请求，并且希望根据请求的 Content - Type 自动处理，ShouldBind 是个好选择。
		当你确定请求数据总是某种特定格式（例如总是 JSON），并且想要避免 Content - Type 头信息错误导致的绑定问题时，ShouldBindBodyWith 更为合适。同时，在一些需要多次尝试不同绑定方式的场景（如先尝试 JSON 绑定，再尝试 XML 绑定），ShouldBindBodyWith 能更方便地实现这种逻辑。
*/

type A1 struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type B1 struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

// 要想多次绑定，可以使用 c.ShouldBindBodyWith.
// ShouldBindBodyWith 方法通常用于处理包含请求体的请求，如 POST、PUT 等
func ShouldBindBodyWith(c *gin.Context) {
	objA1 := A1{}
	objB1 := B1{}

	if errA := c.ShouldBindBodyWith(&objA1, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be form A`)
	} else if errB := c.ShouldBindBodyWith(&objB1, binding.JSON); errB == nil {
		c.String(http.StatusOK, "the body should be form B")
	} else if errB2 := c.ShouldBindBodyWith(&objB1, binding.XML); errB2 == nil {
		c.String(http.StatusOK, "the body should be form B2")
	}
}
