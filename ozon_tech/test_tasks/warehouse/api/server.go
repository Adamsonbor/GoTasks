package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "products")
}

func runServer() {
	var router = gin.Default()

	router.GET("/products", getProducts)

	router.Run("localhost:8000")
}
