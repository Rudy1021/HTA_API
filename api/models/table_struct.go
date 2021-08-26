package models

import "time"

type Attendee struct {
	User_id     int
	Customer_id int
	A_id        int
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
