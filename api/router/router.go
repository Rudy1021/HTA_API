package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users")

	router.POST("/user")

	router.PUT("/user/:id")

	router.DELETE("/user/:id")

	return router
}
