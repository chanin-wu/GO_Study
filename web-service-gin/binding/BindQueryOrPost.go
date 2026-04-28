package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonInfo struct {
	Name     string `form:"name"`
	Address  string `form:"address"`
	Birthday string `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func BindQueryOrPost(c *gin.Context) {
	var person_info PersonInfo

	if err := c.ShouldBind(&person_info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Name: %s, Address: %s, Birthday: %s\n", person_info.Name, person_info.Address, person_info.Birthday)
	c.JSON(http.StatusOK, gin.H{
		"name":     person_info.Name,
		"address":  person_info.Address,
		"birthday": person_info.Birthday,
	})
}
