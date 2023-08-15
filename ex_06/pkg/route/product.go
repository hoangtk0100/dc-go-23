package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/middleware"
)

func addProductRoutes(server Server, rg *gin.RouterGroup) {
	router := rg.Group("/products")
	router.Use(middleware.RequireAuth(server.GetRepository().User(), server.GetTokenMaker()))

	router.POST("/", server.CreateProduct)
	router.PUT("/:id", server.UpdateProduct)
	router.DELETE("/:id", server.DeleteProductByID)
	router.GET("/:id", server.GetProductByID)
	router.GET("/", server.GetProducts)
}
