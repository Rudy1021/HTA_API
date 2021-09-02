package router

import (
	. "HTA_api/api/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	/*-----------------------Attendee-----------------------*/
	router.GET("/SelectAttendees", Attendee_r)

	router.GET("/SelectAttendee/:id", Attendee_one)

	router.POST("/UploadAttendee", Attendee_c)

	router.PUT("/UpdateAttendee", Attendee_u)

	router.DELETE("/DeleteAttendee", Attendee_d)
	/*-----------------------Attendee-----------------------*/

	/*-----------------------Auth-----------------------*/
	router.GET("/SelectAuths", Auth_r)

	router.GET("/SelectAuth/:id", Auth_one)

	router.POST("/UploadAuth", Auth_c)

	router.PUT("/UpdateAuth", Auth_u)

	router.DELETE("/DeleteAuth", Auth_d)
	/*-----------------------Auth-----------------------*/

	/*-----------------------Countersign-----------------------*/
	router.GET("/SelectCountersigns", Countersign_r)

	router.GET("/SelectCountersign/:id", Countersign_one)

	router.POST("/UploadCountersign", Countersign_c)

	router.PUT("/UpdateCountersign", Countersign_u)

	router.DELETE("/DeleteCountersign", Countersign_d)
	/*-----------------------Countersign-----------------------*/

	/*-----------------------Customer_demand-----------------------*/
	router.GET("/SelectCustomer_demands", Customer_demand_r)

	router.GET("/SelectCustomer_demand/:id", Customer_demand_one)

	router.POST("/UploadCustomer_demand", Customer_demand_c)

	router.PUT("/UpdateCustomer_demand", Customer_demand_u)

	router.DELETE("/DeleteCustomer_demand", Customer_demand_d)
	/*-----------------------Customer_demand-----------------------*/

	/*-----------------------Customer-----------------------*/
	router.GET("/SelectCustomers", Customer_r)

	router.GET("/SelectCustomer/:id", Customer_one)

	router.POST("/UploadCustomer", Customer_c)

	router.PUT("/UpdateCustomer", Customer_u)

	router.DELETE("/DeleteCustomer", Customer_d)
	/*-----------------------Customer-----------------------*/

	/*-----------------------Department-----------------------*/
	router.GET("/SelectDepartments", Department_r)

	router.GET("/SelectDepartment/:id", Department_one)

	router.POST("/UploadDepartment", Department_c)

	router.PUT("/UpdateDepartment", Department_u)

	router.DELETE("/DeleteDepartment", Department_d)
	/*-----------------------Department-----------------------*/

	/*-----------------------Files-----------------------*/
	router.GET("/SelectFiles", Files_r)

	router.GET("/SelectFiles/:id", Files_one)

	router.POST("/UploadFiles", Files_c)

	router.PUT("/UpdateFiles", Files_u)

	router.DELETE("/DeleteFiles", Files_d)
	/*-----------------------Files-----------------------*/

	/*-----------------------Interview-----------------------*/
	router.GET("/SelectInterviews", Interview_r)

	router.GET("/SelectInterview/:id", Interview_one)

	router.POST("/UploadInterview", Interview_c)

	router.PUT("/UpdateInterview", Interview_u)

	router.DELETE("/DeleteInterview", Interview_d)
	/*-----------------------Interview-----------------------*/

	/*-----------------------Jig_demand-----------------------*/
	router.GET("/SelectJig_demands", Jig_demand_r)

	router.GET("/SelectJig_demand/:id", Jig_demand_one)

	router.POST("/UploadJig_demand", Jig_demand_c)

	router.PUT("/UpdateJig_demand", Jig_demand_u)

	router.DELETE("/DeleteJig_demand", Jig_demand_d)
	/*-----------------------Jig_demand-----------------------*/

	/*-----------------------Log-----------------------*/
	router.GET("/SelectLogs", Log_r)

	router.GET("/SelectLog/:id", Log_one)

	router.POST("/UploadLog", Log_c)

	router.PUT("/UpdateLog", Log_u)

	router.DELETE("/DeleteLog", Log_d)
	/*-----------------------Log-----------------------*/

	/*-----------------------Machine_type-----------------------*/
	router.GET("/SelectMachine_types", Machine_type_r)

	router.GET("/SelectMachine_type/:id", Machine_type_one)

	router.POST("/UploadMachine_type", Machine_type_c)

	router.PUT("/UpdateMachine_type", Machine_type_u)

	router.DELETE("/DeleteMachine_type", Machine_type_d)
	/*-----------------------Machine_type-----------------------*/

	/*-----------------------Machine_work_place-----------------------*/
	router.GET("/SelectMachine_work_places", Machine_work_place_r)

	router.GET("/SelectMachine_work_place/:id", Machine_work_place_one)

	router.POST("/UploadMachine_work_place", Machine_work_place_c)

	router.PUT("/UpdateMachine_work_place", Machine_work_place_u)

	router.DELETE("/DeleteMachine_work_place", Machine_work_place_d)
	/*-----------------------Machine_work_place-----------------------*/

	/*-----------------------Manufacture_order-----------------------*/
	router.GET("/SelectManufacture_orders", Manufacture_order_r)

	router.GET("/SelectManufacture_order/:id", Manufacture_order_one)

	router.POST("/UploadManufacture_order", Manufacture_order_c)

	router.PUT("/UpdateManufacture_order", Manufacture_order_u)

	router.DELETE("/DeleteManufacture_order", Manufacture_order_d)
	/*-----------------------Manufacture_order-----------------------*/

	/*-----------------------Meeting-----------------------*/
	router.GET("/SelectMeetings", Meeting_r)

	router.GET("/SelectMeeting/:id", Meeting_one)

	router.POST("/UploadMeeting", Meeting_c)

	router.PUT("/UpdateMeeting", Meeting_u)

	router.DELETE("/DeleteMeeting", Meeting_d)
	/*-----------------------Meeting-----------------------*/

	/*-----------------------Notice_time-----------------------*/
	router.GET("/SelectNotice_times", Notice_time_r)

	router.GET("/SelectNotice_time/:id", Notice_time_one)

	router.POST("/UploadNotice_time", Notice_time_c)

	router.PUT("/UpdateNotice_time", Notice_time_u)

	router.DELETE("/DeleteNotice_time", Notice_time_d)
	/*-----------------------Notice_time-----------------------*/

	/*-----------------------Project-----------------------*/
	router.GET("/SelectProjects", Project_r)

	router.GET("/SelectProject/:id", Project_one)

	router.POST("/UploadProject", Project_c)

	router.PUT("/UpdateProject", Project_u)

	router.DELETE("/DeleteProject", Project_d)
	/*-----------------------Project-----------------------*/

	/*-----------------------Sysuser-----------------------*/
	router.GET("/SelectSysusers", Sysuser_r)

	router.GET("/SelectSysuser/:id", Sysuser_one)

	router.POST("/UploadSysuser", Sysuser_c)

	router.PUT("/UpdateSysuser", Sysuser_u)

	router.DELETE("/DeleteSysuser", Sysuser_d)
	/*-----------------------Sysuser-----------------------*/

	/*-----------------------Task-----------------------*/
	router.GET("/SelectTasks", Task_r)

	router.GET("/SelectTask/:id", Task_one)

	router.POST("/UploadTask", Task_c)

	router.PUT("/UpdateTask", Task_u)

	router.DELETE("/DeleteTask", Task_d)
	/*-----------------------Task-----------------------*/

	/*-----------------------Work_item-----------------------*/
	router.GET("/SelectWork_items", Work_item_r)

	router.GET("/SelectWork_item/:id", Work_item_one)

	router.POST("/UploadWork_item", Work_item_c)

	router.PUT("/UpdateWork_item", Work_item_u)

	router.DELETE("/DeleteWork_item", Work_item_d)
	/*-----------------------Work_item-----------------------*/

	/*-----------------------Work_set-----------------------*/
	router.GET("/SelectWorker_sets", Worker_set_r)

	router.GET("/SelectWorker_set/:id", Worker_set_one)

	router.POST("/UploadWorker_set", Worker_set_c)

	router.PUT("/UpdateWorker_set", Worker_set_u)

	router.DELETE("/DeleteWorker_set", Worker_set_d)
	/*-----------------------Work_set-----------------------*/
	return router
}
