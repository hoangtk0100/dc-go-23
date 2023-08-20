package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangtk0100/dc-go-23/ex_08/pkg/middleware"
)

func addCartRoutes(server Server, rg *gin.RouterGroup) {
	router := rg.Group("/cart")
	router.Use(middleware.RequireAuth(server.GetRepository().User(), server.GetTokenMaker()))

	router.GET("/", server.GetCartDetails)
	router.POST("/add", server.AddCartItem)
	router.DELETE("/remove", server.RemoveCartItem)
	router.POST("/checkout", server.Checkout)
}
