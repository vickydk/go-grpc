package omsApi

import "time"

type Order_return struct {
	TableName struct{} `sql:"order_return"`

	ReturnNumber   string    `sql:"type:varchar(100),pk" orm:"column(return_number)" json:"ReturnNumber"`
	LogisticNumber string    `sql:"type:varchar(100)" orm:"colun(logistic_number);null" json:"LogisticNumber"`
	Quantity       int       `sql:",notnull" orm:"column(quantity)" json:"Quantity"`
	Reason         string    `orm:"column(reason)" json:"Reason"`
	Status         int       `sql:",notnull" orm:"column(status)" json:"Status"`
	CreatedBy      int       `orm:"column(created_by);null" json:"CreatedBy"`
	CreatedDate    time.Time `sql:"type:timestamp" orm:"column(created_date);null" json:"CreatedDate"`
	LastModified   time.Time `sql:"type:timestamp, default:CURRENT_TIMESTAMP" orm:"column(last_modified);null" json:"LastModified"`
}

type CreateReqOrderReturn struct {
	OrderId         string               `json:"return_id"`
	ReferenceNumber string               `json:"order_id" validate:"required"`
	OrderItems      []CreateReqOrderItem `json:"order_items" validate:"required"`
}
