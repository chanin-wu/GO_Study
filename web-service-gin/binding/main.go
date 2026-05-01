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

	{
		// binding.Any("/onlyBindQueryString", OnlyBindQueryString)
		binding.GET("/onlyBindQueryString", OnlyBindQueryString)
	}

	{
		router.POST("/binding", Binding)
	}

	{
		binding.POST("/bindQueryOrPost", BindQueryOrPost)
	}

	{
		binding.POST("/bindDefaultValues", BindDefaultValues)
	}

	{
		binding.GET("/collectionFormatForArrays", CollectionFormatForArrays)
	}

	{
		binding.GET("/bindUri/:name/:id/:email", BindUri)
	}

	{
		bindCustomUnmarshaler := binding.Group("/bindCustomUnmarshaler")
		bindCustomUnmarshaler.GET("/textUnmarshaler", TextUnmarshaler)
		bindCustomUnmarshaler.GET("/bindUnmarshaler", BindUnmarshaler)
	}

	{
		binding.POST("/bindHeader", BindHeader)
	}

	router.Run(":8080")
}
