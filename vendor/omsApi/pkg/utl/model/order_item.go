package omsApi

import "time"

const (
	TblOrderItem string = "order_item"
)

type Order_item struct {
	ID            int     `json:"id"`
	OrderHeaderID int     `json:"order_header_id"`
	SKU           string  `json:"sku"`
	Quantity      int     `json:"quantity"`
	ItemPrice     int64   `json:"item_price"`
	ItemDiscPrice float64 `json:"item_disc_price"`
	ItemPaidPrice float64 `json:"item_paid_price"`
	Active        int     `json:"active"`
	CreatedBy     int     `json:"created_by"`
	CreatedDate   string  `json:"created_date"`
	LastModified  string  `json:"last_modified"`
}

type OrderItemReq struct {
	Id            int    `json:"id,omitempty"`
	OrderHeaderID string `json:"order_header_id,omitempty"`
	Status        int    `json:"status,omitempty"`
	Action        int    `json:"action,omitempty"`
}

type OrderItemResp struct {
	Id               int       `json:"id"`
	OrderId          string    `json:"order_id"`
	SKU              string    `json:"sku"`
	Quantity         int       `json:"quantity"`
	ItemPrice        int64     `json:"item_price"`
	ItemDiscPrice    float64   `json:"item_disc_price"`
	ItemPaidPrice    float64   `json:"item_paid_price"`
	Status           int       `json:"status"`
	StatusName       string    `json:"status_name"`
	LogisticNumber   string    `json:"logistic_number"`
	LogisticCourier  int       `json:"logistic_courier"`
	LogisticShipDate time.Time `json:"logistic_ship_date"`
	LogisticStatus   int       `json:"logistic_status"`
}

type CreateReqOrderItem struct {
	OrderHeaderID int64   `json:"order_header_id" validate:"required"`
	SKU           string  `json:"sku" validate:"required"`
	Quantity      int     `json:"quantity" validate:"required"`
	ItemPrice     int64   `json:"item_price" validate:"required"`
	ItemDiscPrice float64 `json:"item_disc_price,omitempty"`
	ItemPaidPrice float64 `json:"item_paid_price" validate:"required"`
}
