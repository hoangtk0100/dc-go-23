package router

import (
	"github.com/gin-gonic/gin"
)

func addCartRoutes(server Server, rg *gin.RouterGroup) {
	router := rg.Group("/cart")

	router.GET("/", server.GetCartDetails)
	router.POST("/add", server.AddCartItem)
	router.DELETE("/remove", server.RemoveCartItem)
	router.POST("/checkout", server.Checkout)
}
