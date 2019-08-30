package omsApi

import "time"

type Order_cancel struct {
	TableName struct{} `sql:"order_cancel"`

	CancelNumber    string    `sql:"type:varchar(100),pk" orm:"pk;column(cancel_number)" json:"CancelNumber"`
	PurchaseCode    string    `sql:"type:varchar(100)" orm:"column(purchase_code);null" json:"PurchaseCode"`
	ReferenceNumber string    `sql:"type:varchar(100),notnull" orm:"column(reference_number);null" json:"ReferenceNumber"`
	Quantity        int       `sql:",notnull" orm:"column(quantity)" json:"Quantity"`
	Reason          string    `orm:"column(reason)" json:"Reason"`
	Status          int       `sql:",notnull" orm:"column(status)" json:"Status"`
	CreatedBy       int       `orm:"column(created_by);null" json:"CreatedBy"`
	CreatedDate     time.Time `sql:"type:timestamp, default:CURRENT_TIMESTAMP" orm:"column(created_date);null" json:"CreatedDate"`
	LastModified    time.Time `sql:"type:timestamp, default:CURRENT_TIMESTAMP" orm:"column(last_modified);null" json:"LastModified"`
}
