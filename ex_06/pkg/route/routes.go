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

	// Cart
	GetCartDetails(ctx *gin.Context)
	AddCartItem(ctx *gin.Context)
	RemoveCartItem(ctx *gin.Context)
	Checkout(ctx *gin.Context)

	// User
	Register(ctx *gin.Context)

	// Auth
	Login(ctx *gin.Context)
}

func SetupRoutes(server Server) {
	v1 := server.GetRouter().Group("/v1")

	addProductRoutes(server, v1)
	addCartRoutes(server, v1)
	addUserRoutes(server, v1)
	addAuthRoutes(server, v1)
}
