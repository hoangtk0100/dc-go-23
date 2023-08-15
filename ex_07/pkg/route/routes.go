package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/middleware"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/repository"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/token"
)

type Server interface {
	GetRouter() *gin.Engine
	GetRepository() repository.Repository
	GetTokenMaker() token.TokenMaker
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
	v1.Use(middleware.Recovery())

	addAuthRoutes(server, v1)
	addUserRoutes(server, v1)
	addProductRoutes(server, v1)
	addCartRoutes(server, v1)
}
