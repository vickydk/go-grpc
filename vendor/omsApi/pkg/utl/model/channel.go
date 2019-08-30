package omsApi

type Channel struct {
	TableName struct{} `sql:"channel"`

	Base
	Name        string `sql:"type:varchar(100)" orm:"column(name)" json:"Name"`
	Description string `orm:"column(description)" json:"Description"`
}
