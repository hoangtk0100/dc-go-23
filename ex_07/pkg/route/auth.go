package router

import "github.com/gin-gonic/gin"

func addAuthRoutes(server Server, rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/login", server.Login)
}
