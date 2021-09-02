package models

import (
	"time"
)

type Attendee struct {
	User_id     int `json:"user_id"`
	Customer_id int `json:"customer_id"`
	A_id        int `gorm:"<-:update;primary_key" json:"a_id"`
}

type Countersign struct {
	C_id        int
	Dep_id      int
	Feedback    string
	Create_time time.Time
	Status      bool
}

// type CountSign struct {
// 	C_id        int
// 	Dep_id      int
// 	Status      bool
// 	Feedback    string
// 	Create_time time.Time
// }

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

type Department struct {
	D_id         int
	Name         string
	Eng_name     string
	Introduction string
	Manager      int
	Tel          string
	Fax          string
}

type Files struct {
	F_id        int
	Name        string
	Path        string
	Create_time time.Time
}

type Interview struct {
	I_id          int
	Time          time.Time
	Tpye          string
	Content       string
	Cus_demand_id int
	Create_time   time.Time
	Creater       int
}

type Jig_demand struct {
	J_id                      int
	Kind                      string
	Type                      string
	Urgent                    bool
	Title                     string
	Quantity                  int
	Date_for_notify           time.Time
	Information               time.Time
	Expect_shipment_day       time.Time
	Po_date                   time.Time
	Order_name                string
	Project_id                int
	Item                      string
	Standard                  string
	Summary                   string
	Remark                    string
	Liaison                   string
	Liaison_phone             string
	Date_for_demand           time.Time
	Date_for_respond          time.Time
	Date_for_respond_of_limit time.Time
	Customer_id               int
	Create_time               time.Time
	Creater                   int
}

type Logs struct {
	L_id        int
	Tpye        string
	Tablename   string
	Sql_code    string
	Content     string
	User        string
	Create_time time.Time
}

type Machine_type struct {
	M_id        int
	Code        string
	Name        string
	Description string
	Create_time time.Time
}

type Machine_work_place struct {
	M_id            int
	Enable          bool
	Place           string
	Machine_type_id int
	Pn_code         string
	Sn_code         string
	Customer_id     int
	Project_code    string
	Remark          string
	Create_time     time.Time
	Creater         int
	Task_id         []int
	Meeting_id      []int
	Countersign_id  []int
}

type Manufacture_order struct {
	M_id                int
	Customer_id         string
	Order_name          string
	Amount              string
	Shipment_location   string
	Open_date           time.Time
	Close_date          time.Time
	Expect_shipment_day time.Time
	Sales_assistant     string
	recipient           string
	Contact_person      string
	Remark              string
	Create_time         time.Time
	Project_id          int
	Copy                []int
	Status              string
	Creater             int
}

type Meeting struct {
	M_id        int
	Name        string
	Room        string
	Start_date  time.Time
	End_date    time.Time
	Create_time time.Time
	Chairman    int
	Attendee    int
}

type Notice_time struct {
	N_id      int
	Meet_id   int
	Meet_time time.Time
}

type Project struct {
	P_id                   int
	Code                   string
	Name                   string
	Customer_id            int
	Date_for_start         time.Time
	Date_for_end           time.Time
	Salesman_id            int
	Serviceman_id          int
	Projectman_id          int
	Status                 string
	Create_time            time.Time
	Type                   string
	Project_member         []string
	Meeting_id             []string
	File                   []string
	Task_id                []string
	Manufactrue_order_list []string
}

type Sysuser struct {
	Logonid    string
	Name       string
	Name1      string
	Title      string
	Email      string
	Gender     string
	Empno      string
	Logenabled string
	Labuser    string
	Gboss      string
	Sysboss    string
	Gkind      string
	Gkindboss  string
	Rdopt      string
}

type Task struct {
	T_id            int
	Type            string
	Name            string
	Description     string
	Principal       string
	Before_id       string
	Time_for_start  time.Time
	Time_for_done   time.Time
	Create_time     time.Time
	Include_weekend bool
	Labor_hour      int
	Finish          bool
	Creater         int
}

type Work_item struct {
	W_id           int
	Name           string
	Type           string
	Required       bool
	Provide_amount bool
	Create_time    time.Time
}

type Worker_set struct {
	W_id              int
	Work_item         int
	User_id           string
	Date_for_expected time.Time
	Date_for_done     time.Time
	Amount            string
	Remark            string
	Create_time       time.Time
}

type Auth struct {
	A_id        int `gorm:"<-:update;primary_key" json:"a_id"`
	Name        string
	Path        string
	File        string
	Create_time *time.Time
}
