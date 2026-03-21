package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album 表示有关专辑的数据.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// 专辑切片以填充专辑数据记录.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// AsciiJSON()
	// HttpMethod()
	ParamInPath()
}

/* func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
} */

// getAlbums 以 JSON 格式响应所有专辑的列表.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums 从请求体中收到的JSON中添加一个专辑 .
func postAlbums(c *gin.Context) {
	var newAlbum album

	// 调用 BindJSON 将收到的 JSON 绑定到
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// 将新专辑添加到切片.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID 查找 ID 值与客户端发送的 id
// 参数匹配的专辑，然后返回该专辑作为响应.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// 循环浏览专辑列表，查找
	// ID 值与参数匹配的专辑.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
