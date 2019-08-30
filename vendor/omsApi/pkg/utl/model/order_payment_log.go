package omsApi

import "time"

type Order_payment_log struct {
	TableName struct{} `sql:"order_payment_log"`

	Base
	OrderPayment    *Order_payment `json:"-"`
	OrderPaymentID  int            `json:"order_payment_id"`
	PurchaseCode    string         `sql:"type:varchar(100)" json:"purchase_code"`
	MasterPaymentId int            `json:"master_payment_id"`
	TransactionDate time.Time      `sql:"type:timestamp" json:"transaction_date"`
	Status          int            `sql:"default:0" json:"status"`
}
