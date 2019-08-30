package omsApi

import "time"

type Customer struct {
	TableName struct{} `sql:"customer"`

	Base
	CustomerEmail    string    `sql:"type:varchar(100),unique" json:"CustomerEmail"`
	CustomerFName    string    `sql:"customer_fname, type:varchar(100)" json:"CustomerFName"`
	CustomerLName    string    `sql:"customer_lname, type:varchar(100)" json:"CustomerLName"`
	CustomerPassword string    `sql:"type:varchar(255)" json:"CustomerPassword"`
	CustomerPhone    string    `sql:"type:varchar(100)" json:"CustomerPhone"`
	CustomerGender   int       `json:"CustomerGender"`
	Status           int       `json:"Status"`
	CustomerDob      time.Time `sql:"type:timestamp" json:"CustomerDob"`
	CustomerRegDate  time.Time `sql:"type:timestamp" json:"CustomerRegDate"`
	ResetSentOn      time.Time `sql:"resetsenton, type:timestamp" json:"ResetSentOn"`
	LastLoginDate    time.Time `sql:"type:timestamp" json:"LastLoginDate"`
	UpdateDate       time.Time `sql:"type:timestamp" json:"UpdateDate"`
}
