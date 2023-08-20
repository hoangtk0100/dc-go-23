package router

import "github.com/gin-gonic/gin"

func addUserRoutes(server Server, rg *gin.RouterGroup) {
	router := rg.Group("/users")

	router.POST("/register", server.Register)
}
