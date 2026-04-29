package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	ShouldBindUri 使用 uri 结构体标签将 URI 路径参数直接绑定到结构体中。结合 binding 验证标签，这让你可以通过一次调用来验证路径参数（例如要求有效的 UUID）。
*/

type PersonUri struct {
	Email string `uri:"email" binding:"required,email"`
	ID    string `uri:"id" binding:"required,uuid"`
	Name  string `uri:"name" binding:"required"`
}

func BindUri(c *gin.Context) {
	var person_uri PersonUri
	if err := c.ShouldBindUri(&person_uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Email": person_uri.Email,
		"ID":    person_uri.ID,
		"Name":  person_uri.Name,
	})
}
