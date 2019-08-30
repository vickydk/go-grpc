package omsApi

type Order_item_log struct {
	TableName struct{} `sql:"order_item_log"`

	Base
	OrderItem      *Order_item `json:"-"`
	OrderItemID    int         `json:"order_item_id"`
	PurchaseCode   string      `sql:"type:varchar(100)" json:"purchase_code"`
	SKU            string      `sql:"type:varchar(255)" json:"sku"`
	Status         int         `sql:"default:1" json:"status"`
	LogisticStatus int         `json:"logistic_status"`
}
