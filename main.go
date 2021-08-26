package main

import (
	_ "HTA_api/api/database"
	orm "HTA_api/api/database"
	"HTA_api/api/router"
)

func main() {
	defer orm.Db.Close()
	router := router.InitRouter()
	router.Run(":5555")
}
