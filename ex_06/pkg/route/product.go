package router

import (
	"github.com/gin-gonic/gin"
)

func addProductRoutes(server Server, rg *gin.RouterGroup) {
	router := rg.Group("/products")

	router.POST("/", server.CreateProduct)
	router.PUT("/:id", server.UpdateProduct)
	router.DELETE("/:id", server.DeleteProductByID)
	router.GET("/:id", server.GetProductByID)
	router.GET("/", server.GetProducts)
}
