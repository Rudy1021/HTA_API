package main

import (
	_ "HTA_api/api/database"
	orm "HTA_api/api/database"
	"HTA_api/api/router"
)

func main() {
	/*
		var user []models.Sysuser
		db.Find(&user)
		fmt.Print(user)
		var attendees []models.Attendee
		db.Find(&attendees)
		fmt.Print(attendees)
	*/
	defer orm.Db.Close()
	router := router.InitRouter()
	router.Run(":5555")
}
