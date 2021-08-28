package models

import "time"

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
	A_id        int
	Name        string
	Path        string
	File        string
	Create_time time.Time
}
