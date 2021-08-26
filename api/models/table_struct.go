package models

type Attendee struct {
	User_id     int
	Customer_id int
	A_id        *int
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
