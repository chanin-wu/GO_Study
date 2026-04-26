package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	binding := router.Group("/binding")

	{
		bindingAndValidation := binding.Group("/bindingAndValidation")
		bindingAndValidation.POST("/loginJSON", LoginJSON)
		bindingAndValidation.POST("/loginXML", LoginXML)
		bindingAndValidation.POST("/loginForm", LoginForm)
	}

	{
		binding.GET("/customValidators", CustomValidators)
	}

	router.Run(":8080")
}
