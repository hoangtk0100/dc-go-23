package router

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	GetRouter() *gin.Engine
	Start()

	// Product
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProductByID(ctx *gin.Context)
	GetProductByID(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
}

func SetupRoutes(server Server) {
	v1 := server.GetRouter().Group("/v1")

	addProductRoutes(server, v1)
}
