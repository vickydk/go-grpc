package omsApi

import "time"

type Order_header_log struct {
	TableName struct{} `sql:"order_header_log"`

	Base
	OrderHeader     *Order_header `json:"-"`
	OrderHeaderID   int           `json:"order_header_id"`
	PurchaseCode    string        `sql:"type:varchar(100)" json:"purchase_code"`
	ReferenceNumber string        `sql:"type:varchar(100)" json:"reference_number"`
	CustomerId      int           `json:"customer_id"`
	SourceLocation  int           `json:"source_location"`
	OrderDate       time.Time     `sql:"type:timestamp" json:"order_date"`
	Status          int           `sql:"default:1" json:"status"`
}
