package omsApi

type Order_return_log struct {
	TableName struct{} `sql:"order_return_log"`

	ReturnNumber   string `sql:"type:varchar(100),pk" orm:"pk;column(return_number)" json:"ReturnNumber"`
	LogisticNumber string `sql:"type:varchar(100)" orm:"colun(logistic_number);null" json:"LogisticNumber"`
	Reason         string `orm:"column(reason)" json:"Reason"`
	Status         int    `sql:",notnull" orm:"column(status)" json:"Status"`
}
