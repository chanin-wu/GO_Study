package rendering

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	rendering := router.Group("/rendering")

	{
		rendering.POST("/", Rendering)
	}

	router.Run(":8080")
}
