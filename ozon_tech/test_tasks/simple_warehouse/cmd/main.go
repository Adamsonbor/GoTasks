package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func reserveProduct(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}

func releaseProduct(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}

func remainProducts(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}

func main() {
	route := gin.Default()

	route.GET("/reserve", reserveProduct)
	route.GET("/release", releaseProduct)
	route.GET("/remain", remainProducts)

	route.Run()
}
