package router

import (
	. "HTA_api/api/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/SelectAttendees", Attendee_r)

	router.GET("/SelectAttendee/:id", Attendee_one)

	router.POST("/UploadAttendee", Attendee_c)

	router.PUT("/UpdateAttendee", Attendee_u)

	router.DELETE("/DeleteAttendee", Attendee_d)

	router.GET("/SelectAuth", Auth_r)

	router.POST("/UploadAuth", Auth_c)

	router.PUT("/UpdateAuth")

	router.DELETE("/DeleteAuth", Auth_d)

	//Countersign
	router.GET("/SelectCountersign")

	router.POST("/UploadCountersign")

	router.PUT("/UpdateCountersign")

	router.DELETE("/DeleteCountersign")

	//Customer
	router.GET("/SelectCustomer")

	router.POST("/UploadCustomer")

	router.PUT("/UpdateCustomer")

	router.DELETE("/DeleteCustomer")

	//Countersign
	router.GET("/SelectCountersign")

	router.POST("/UploadCountersign")

	router.PUT("/UpdateCountersign")

	router.DELETE("/DeleteCountersign")

	//Customer_demand
	router.GET("/SelectCustomerDemand")

	router.POST("/UploadCustomerDemand")

	router.PUT("/UpdateCustomerDemand")

	router.DELETE("/DeleteCustomerDemand")

	//Department
	router.GET("/SelectDepartment")

	router.POST("/UploadDepartment")

	router.PUT("/UpdateDepartment")

	router.DELETE("/DeleteDepartment")

	//Files
	router.GET("/SelectFiles")

	router.POST("/UploadFiles")

	router.PUT("/UpdateFiles")

	router.DELETE("/DeleteFiles")

	//Interview
	router.GET("/SelectInterview")

	router.POST("/UploadInterview")

	router.PUT("/UpdateInterview")

	router.DELETE("/DeleteInterview")

	//Jid_demand
	router.GET("/SelectJidDemand")

	router.POST("/UploadJidDemand")

	router.PUT("/UpdateJidDemand")

	router.DELETE("/DeleteJidDemand")

	//Log
	router.GET("/SelectLog")

	router.POST("/UploadLog")

	router.PUT("/UpdateLog")

	router.DELETE("/DeleteLog")

	//Machine_type
	router.GET("/SelectMachineType")

	router.POST("/UploadMachineType")

	router.PUT("/UpdateMachineType")

	router.DELETE("/DeleteMachineType")

	//Machine_work_place
	router.GET("/SelectMachineWorkPlace")

	router.POST("/UploadMachineWorkPlace")

	router.PUT("/UpdateMachineWorkPlace")

	router.DELETE("/DeleteMachineWorkPlace")

	//Manufactrue_order
	router.GET("/SelectManufactrueOrder")

	router.POST("/UploadManufactrueOrder")

	router.PUT("/UpdateManufactrueOrder")

	router.DELETE("/DeleteManufactrueOrder")

	//Meeting
	router.GET("/SelectMeeting")

	router.POST("/UploadMeeting")

	router.PUT("/UpdateMeeting")

	router.DELETE("/DeleteMeeting")

	//NoticeTime
	router.GET("/SelectNoticeTime")

	router.POST("/UploadNoticeTime")

	router.PUT("/UpdateNoticeTime")

	router.DELETE("/DeleteNoticeTime")

	//Project
	router.GET("/SelectProject")

	router.POST("/UploadProject")

	router.PUT("/UpdateProjecte")

	router.DELETE("/DeleteProject")

	//Sysuser
	router.GET("/SelectSysuser")

	router.POST("/UploadSysuser")

	router.PUT("/UpdateSysuser")

	router.DELETE("/DeleteSysuser")

	//Task
	router.GET("/SelectTask")

	router.POST("/UploadTask")

	router.PUT("/UpdateTask")

	router.DELETE("/DeleteTask")

	//Work_item
	router.GET("/SelectWorkItem")

	router.POST("/UploadWorkItem")

	router.PUT("/UpdateWorkItem")

	router.DELETE("/DeleteWorkItem")

	//Work_set
	router.GET("/SelectWorkSet")

	router.POST("/UploadWorkSet")

	router.PUT("/UpdateWorkSet")

	router.DELETE("/DeleteWorkSet")

	//Auth
	router.GET("/SelectAuth")

	router.POST("/UploadAuth")

	router.PUT("/UpdateAuth")

	router.DELETE("/DeleteAuth")
	return router
}
