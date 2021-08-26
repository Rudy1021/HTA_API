package router

import (
	. "HTA_api/api/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", HelloWorld)

	router.POST("/user")

	router.PUT("/user/:id")

	router.DELETE("/user/:id")

	return router
}
