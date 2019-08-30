package omsApi

import "time"

type Order_cancel_log struct {
	TableName struct{} `sql:"order_cancel_log"`

	CancelNumber string    `sql:"type:varchar(100),pk" orm:"pk;column(cancel_number)" json:"CancelNumber"`
	OrderItemId  int       `sql:",notnull" orm:"column(order_item_id)" json:"OrderItemId"`
	SKU          string    `sql:"type:varchar(255),notnull" orm:"column(SKU);null" json:"SKU"`
	Quantity     int       `sql:",notnull" orm:"column(quantity)" json:"Quantity"`
	Status       int       `sql:",notnull" orm:"column(status)" json:"Status"`
	CreatedBy    int       `orm:"column(created_by);null" json:"CreatedBy"`
	CreatedDate  time.Time `sql:"type:timestamp, default:CURRENT_TIMESTAMP" orm:"column(created_date);null" json:"CreatedDate"`
	LastModified time.Time `sql:"type:timestamp, default:CURRENT_TIMESTAMP" orm:"column(last_modified);null" json:"LastModified"`
}
