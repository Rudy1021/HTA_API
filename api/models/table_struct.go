package models

type Attendee struct {
	User_id     int  `json:"user_id"`
	Customer_id int  `json:"customer_id"`
	A_id        *int `gorm:"<-:update;primary_key" json:"a_id"`
}
type Attendee_u struct {
	A_id        int
	User_id     int
	Customer_id int
}
type Sysuser struct {
	Logonid string
	Name    string
}
type Auth struct {
}
