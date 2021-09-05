package router

import (
	. "HTA_api/api/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	/*-----------------------Attendee-----------------------*/
	router.GET("/Attendee/Select", Attendee_r)

	router.GET("/Attendee/Select/:id", Attendee_one)

	router.POST("/Attendee/Upload", Attendee_c)

	router.PUT("/Attendee/Update", Attendee_u)

	router.DELETE("/Attendee/Delete", Attendee_d)
	/*-----------------------Attendee-----------------------*/

	/*-----------------------Auth-----------------------*/
	router.GET("/Auth/Select", Auth_r)

	router.GET("/Auth/Select/:id", Auth_one)

	router.POST("/Auth/Upload", Auth_c)

	router.PUT("/Auth/Update", Auth_u)

<<<<<<< HEAD
	router.DELETE("/DeleteAuth", Auth_d)
	/*
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

		router.DELETE("/DeleteAuth")*/
=======
	router.DELETE("/Auth/Delete", Auth_d)
	/*-----------------------Auth-----------------------*/

	/*-----------------------Countersign-----------------------*/
	router.GET("/Countersign/Select", Countersign_r)

	router.GET("/Countersign/Select/:id", Countersign_one)

	router.POST("/Countersign/Upload", Countersign_c)

	router.PUT("/Countersign/Update", Countersign_u)

	router.DELETE("/Countersign/Delete", Countersign_d)
	/*-----------------------Countersign-----------------------*/

	/*-----------------------Customer_demand-----------------------*/
	router.GET("/Customer_demand/Select", Customer_demand_r)

	router.GET("/Customer_demand/Select/:id", Customer_demand_one)

	router.POST("/Customer_demand/Upload", Customer_demand_c)

	router.PUT("/Customer_demand/Update", Customer_demand_u)

	router.DELETE("/Customer_demand/Delete", Customer_demand_d)
	/*-----------------------Customer_demand-----------------------*/

	/*-----------------------Customer-----------------------*/
	router.GET("/Customer/Select", Customer_r)

	router.GET("/Customer/Select/:id", Customer_one)

	router.POST("/Customer/Upload", Customer_c)

	router.PUT("/Customer/Update", Customer_u)

	router.DELETE("/Customer/Delete", Customer_d)
	/*-----------------------Customer-----------------------*/

	/*-----------------------Department-----------------------*/
	router.GET("/Department/Select", Department_r)

	router.GET("/Department/Select/:id", Department_one)

	router.POST("/Department/Upload", Department_c)

	router.PUT("/Department/Update", Department_u)

	router.DELETE("/Department/Delete", Department_d)
	/*-----------------------Department-----------------------*/

	/*-----------------------Files-----------------------*/
	router.GET("/Files/Select", Files_r)

	router.GET("/Files/Select/:id", Files_one)

	router.POST("/Files/Upload", Files_c)

	router.PUT("/Files/Update", Files_u)

	router.DELETE("/Files/Delete", Files_d)
	/*-----------------------Files-----------------------*/

	/*-----------------------Interview-----------------------*/
	router.GET("/Interview/Select", Interview_r)

	router.GET("/Interview/Select/:id", Interview_one)

	router.POST("/Interview/Upload", Interview_c)

	router.PUT("/Interview/Update", Interview_u)

	router.DELETE("/Interview/Delete", Interview_d)
	/*-----------------------Interview-----------------------*/

	/*-----------------------Jig_demand-----------------------*/
	router.GET("/Jig_demand/Select", Jig_demand_r)

	router.GET("/Jig_demand/Select/:id", Jig_demand_one)

	router.POST("/Jig_demand/Upload", Jig_demand_c)

	router.PUT("/Jig_demand/Update", Jig_demand_u)

	router.DELETE("/Jig_demand/Delete", Jig_demand_d)
	/*-----------------------Jig_demand-----------------------*/

	/*-----------------------Log-----------------------*/
	router.GET("/Log/Select", Log_r)

	router.GET("/Log/Select/:id", Log_one)

	router.POST("/Log/Upload", Log_c)

	router.PUT("/Log/Update", Log_u)

	router.DELETE("/Log/Delete", Log_d)
	/*-----------------------Log-----------------------*/

	/*-----------------------Machine_type-----------------------*/
	router.GET("/Machine_type/Select", Machine_type_r)

	router.GET("/Machine_type/Select/:id", Machine_type_one)

	router.POST("/Machine_type/Upload", Machine_type_c)

	router.PUT("/Machine_type/Update", Machine_type_u)

	router.DELETE("/Machine_type/Delete", Machine_type_d)
	/*-----------------------Machine_type-----------------------*/

	/*-----------------------Machine_work_place-----------------------*/
	router.GET("/Machine_work_place/Select", Machine_work_place_r)

	router.GET("/Machine_work_place/Select/:id", Machine_work_place_one)

	router.POST("/Machine_work_place/Upload", Machine_work_place_c)

	router.PUT("/Machine_work_place/Update", Machine_work_place_u)

	router.DELETE("/Machine_work_place/Delete", Machine_work_place_d)
	/*-----------------------Machine_work_place-----------------------*/

	/*-----------------------Manufacture_order-----------------------*/
	router.GET("/Manufacture_order/Select", Manufacture_order_r)

	router.GET("/Manufacture_order/Select/:id", Manufacture_order_one)

	router.POST("/Manufacture_order/Upload", Manufacture_order_c)

	router.PUT("/Manufacture_order/Update", Manufacture_order_u)

	router.DELETE("/Manufacture_order/Delete", Manufacture_order_d)
	/*-----------------------Manufacture_order-----------------------*/

	/*-----------------------Meeting-----------------------*/
	router.GET("/Meeting/Select", Meeting_r)

	router.GET("/Meeting/Select/:id", Meeting_one)

	router.POST("/Meeting/Upload", Meeting_c)

	router.PUT("/Meeting/Update", Meeting_u)

	router.DELETE("/Meeting/Delete", Meeting_d)
	/*-----------------------Meeting-----------------------*/

	/*-----------------------Notice_time-----------------------*/
	router.GET("/Notice_time/Select", Notice_time_r)

	router.GET("/Notice_time/Select/:id", Notice_time_one)

	router.POST("/Notice_time/Upload", Notice_time_c)

	router.PUT("/Notice_time/Update", Notice_time_u)

	router.DELETE("/Notice_time/Delete", Notice_time_d)
	/*-----------------------Notice_time-----------------------*/

	/*-----------------------Project-----------------------*/
	router.GET("/Project/Select", Project_r)

	router.GET("/Project/Select/:id", Project_one)

	router.POST("/Project/Upload", Project_c)

	router.PUT("/Project/Update", Project_u)

	router.DELETE("/Project/Delete", Project_d)
	/*-----------------------Project-----------------------*/

	/*-----------------------Sysuser-----------------------*/
	router.GET("/Sysuser/Select", Sysuser_r)

	router.GET("/Sysuser/Select/:id", Sysuser_one)

	router.POST("/Sysuser/Upload", Sysuser_c)

	router.PUT("/Sysuser/Update", Sysuser_u)

	router.DELETE("/Sysuser/Delete", Sysuser_d)
	/*-----------------------Sysuser-----------------------*/

	/*-----------------------Task-----------------------*/
	router.GET("/Task/Select", Task_r)

	router.GET("/Task/Select/:id", Task_one)

	router.POST("/Task/Upload", Task_c)

	router.PUT("/Task/Update", Task_u)

	router.DELETE("/Task/Delete", Task_d)
	/*-----------------------Task-----------------------*/

	/*-----------------------Work_item-----------------------*/
	router.GET("/Work_item/Select", Work_item_r)

	router.GET("/Work_item/Select/:id", Work_item_one)

	router.POST("/Work_item/Upload", Work_item_c)

	router.PUT("/Work_item/Update", Work_item_u)

	router.DELETE("/Work_item/Delete", Work_item_d)
	/*-----------------------Work_item-----------------------*/

	/*-----------------------Work_set-----------------------*/
	router.GET("/Worker_set/Select", Worker_set_r)

	router.GET("/Worker_set/Select/:id", Worker_set_one)

	router.POST("/Worker_set/Upload", Worker_set_c)

	router.PUT("/Worker_set/Update", Worker_set_u)

	router.DELETE("/Worker_set/Delete", Worker_set_d)
	/*-----------------------Work_set-----------------------*/
>>>>>>> 5188e78607fb41897e5d42faff41f74375f4bad6
	return router
}
