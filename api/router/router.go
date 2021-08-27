package router

import (
	. "HTA_api/api/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/SelectAttendee", Attendee_r)

	router.POST("/UploadAttendee", Attendee_c)

	router.PUT("/UpdateAttendee", Attendee_u)

	router.DELETE("/DeleteAttendee", Attendee_d)

	router.GET("/test", Test)
	return router
}
