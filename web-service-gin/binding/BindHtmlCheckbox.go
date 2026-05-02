package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	具有相同 name 属性的 HTML 复选框在被选中时会提交多个值。Gin 可以通过使用带有 [] 后缀的 form 结构体标签（匹配 HTML 的 name 属性）将这些值直接绑定到结构体的 []string 切片中
	这对于用户选择一个或多个选项的表单非常有用——例如颜色选择器、权限选择器或多选过滤器
	colors[] 中的 [] 后缀是 HTML 的约定，不是 Go 的要求。结构体标签必须与 HTML 的 name 属性完全匹配。如果你的 HTML 使用 name="colors"（不带方括号），你的结构体标签应该是 form:"colors"
*/

type FakeForm struct {
	Colors []string `form:"colors[]"`
}

func BindHtmlCheckbox(c *gin.Context) {
	var fakeForm FakeForm
	if err := c.ShouldBind(&fakeForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}
