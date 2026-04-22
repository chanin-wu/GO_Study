package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	type A struct {：声明一个名为 A 的结构体。
	Foo string：在结构体 A 中定义了一个名为 Foo 的字段，其类型为 string
	json:"foo"：这是一个结构体标签
	binding:"required"：表示 Foo 字段是必需的
*/

// 比如说，你有个 A 盒子，里面 Foo 装着 “苹果” 。当把 A 盒子变成 JSON 格式数据后，看起来就像 {"foo":"苹果"} ，这里 Foo 就变成了 foo ，这就是这个神奇纸条 “json:"foo"” 的作用。
// 变成 XML 格式后，可能就像 <A><foo>苹果</foo></A> ，Foo 被 <foo> 标签包起来了，这就是 “xml:"foo"” 这个纸条的作用。
type A struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type B struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

// 一般通过调用 c.Request.Body 方法绑定数据，但不能多次调用这个方法。
func SomeHandler(c *gin.Context) {
	objA := A{}
	objB := B{}

	// &objA 表示获取变量 objA 的内存地址，返回一个指向 objA 的指针。
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be form A`)
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, "the body should be form B")
	}
}
