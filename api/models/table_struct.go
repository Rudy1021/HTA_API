package models

import (
	"time"
)

type Attendee struct {
	User_id     int `json:"user_id"`
	Customer_id int `json:"customer_id"`
	A_id        int `gorm:"<-:update;primary_key" json:"a_id"`
}
type Sysuser struct {
	Logonid string
	Name    string
}
type Auth struct {
	A_id        int `gorm:"<-:update;primary_key" json:"a_id"`
	Name        string
	Path        string
	File        string
	Create_time *time.Time
}
type CountSign struct {
	C_id        int
	Dep_id      int
	Status      bool
	Feedback    string
	Create_time time.Time
}
type Customer struct {
	C_id        int
	Short_name  string
	Eng_name    string
	Name        string
	Zip_code    string
	Address     string
	Tel         string
	Fax         string
	Email       string
	Map         string
	Creater     string
	Create_time time.Time
}
type Customer_demand struct {
	Id                int
	Customer_id       int
	Task_name         string
	Budget            string
	Remarks           string
	Extend_type_id    int
	Extend_rem        string
	Est_quantity      int
	Countersign_id    []int
	Meeting_id        []int
	Date_for_recive   time.Time
	Task_id           []int
	Accept            bool
	Date_for_devlop   time.Time
	Eva_report        bool
	Date_for_expected time.Time
	Date_for_done     time.Time
	Project_code      string
	Salesman_id       int
	File_id           int
	Creater           int
	Create_time       time.Time
}
