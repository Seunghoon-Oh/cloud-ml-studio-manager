package main

import (
	"net/http"

	v1 "github.com/Seunghoon-Oh/cloud-ml-studio-manager/controller/v1"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "홈페이지 입니다.",
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.GET("/studios", v1.GetStudios)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8082")
}
