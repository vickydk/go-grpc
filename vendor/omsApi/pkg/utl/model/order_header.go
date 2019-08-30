package omsApi

const (
	TblOrderHeader string = "order_header"
)

type Order_header struct {
	ID                  int64   `json:"id"`
	OrderId             string  `json:"order_id"`
	ReferenceNumber     string  `json:"reference_number"`
	CustomerId          int     `json:"customer_id"`
	CustomerEmail       string  `json:"customer_email"`
	CustomerName        string  `json:"customer_name"`
	CustomerPhone       string  `json:"customer_phone"`
	ShippingAddress     string  `json:"shipping_address"`
	BillingAddress      string  `json:"billing_address"`
	OrderDate           string  `json:"order_date"`
	SourceChannel       int     `json:"source_channel"`
	SourceOrder         int     `json:"source_order"`
	SourceLocation      int     `json:"source_location"`
	CouponCode          string  `json:"coupon_code"`
	CouponValue         float64 `json:"coupon_value"`
	FulfillmentLocation int     `json:"fulfillment_location"`
	Awb                 string  `json:"awb"`
	LogisticCourier     int     `json:"logistic_courier"`
	LogisticCourierType int     `json:"logistic_courier_type"`
	LogisticShipDate    string  `json:"logistic_ship_date"`
	LogisticStatus      int     `json:"logistic_status"`
	DeliveredDate       string  `json:"delivered_date"`
	CreatedBy           int     `json:"created_by"`
	CreatedDate         string  `json:"created_date"`
	LastModified        string  `json:"last_modified"`
	Status              int     `json:"status"`
	SubStatus           int     `json:"sub_status"`
	OrderType           int     `json:"order_type"`
	OrderPaymentId      int     `json:"order_payment_id"`
}

type OrderHeaderUpStatReq struct {
	OrderId             string `json:"order_id" validate:"required"`
	CustomerId          int    `json:"customer_id,omitempty"`
	Status              int    `json:"status" validate:"required"`
	SubStatus           int    `json:"sub_status,omitempty"`
	Action              int    `json:"action,omitempty"`
	SubAction           int    `json:"sub_action,omitempty"`
	TransferOrder       int    `json:"transfer_order_location,omitempty"`
	LogisticNumber      string `json:"logistic_number,omitempty"`
	LogisticCourier     int    `json:"logistic_courier,omitempty"`
	LogisticCourierType int    `json:"logistic_courier_type,omitempty"`
	LogisticCourierName string `json:"logistic_courier_name,omitempty"`
	LogisticShipDate    string `json:"logistic_ship_date,omitempty"`
	LogisticStatus      int    `json:"logistic_status,omitempty"`
}

type RespondOrderHeader struct {
	Id      int64  `json:"id"`
	OrderId string `json:"order_id"`
	//Age               float64   `json:"age"`
	OrderType           int               `json:"order_type"`
	OrderDate           string            `json:"order_date"`
	ReferenceNumber     string            `json:"reference_number"`
	CustomerId          int               `json:"customer_id"`
	CustomerEmail       string            `json:"customer_email"`
	CustomerName        string            `json:"customer_name"`
	CustomerPhone       string            `json:"customer_phone"`
	ShippingAddress     string            `json:"shipping_address"`
	BillingAddress      string            `json:"billing_address"`
	SourceChannel       int               `json:"source_channel_id"`
	SourceChannelName   string            `json:"source_channel_name"`
	SourceOrder         int               `json:"source_order"`
	SourceOrderName     string            `json:"source_order_name"`
	SourceLocation      int               `json:"source_location"`
	CouponCode          string            `json:"coupon_code"`
	CouponValue         float64           `json:"coupon_value"`
	TotalDisc           float64           `json:"total_disc"`
	FulfillmentLocation int               `json:"fulfillment_location"`
	Awb                 string            `json:"awb"`
	LogisticCourier     int               `json:"logistic_courier"`
	LogisticCourierType int               `json:"logistic_courier_type"`
	LogisticShipDate    string            `json:"logistic_ship_date"`
	LogisticStatus      int               `json:"logistic_status"`
	Quantity            int               `json:"quantity"`
	Amount              float64           `json:"amount"`
	Status              int               `json:"status"`
	StatusName          string            `json:"status_name"`
	SubStatus           int               `json:"sub_status"`
	SubStatusName       string            `json:"sub_status_name"`
	OrderItems          *[]Order_item     `json:"order_items"`
	Payment             *OrderPaymentResp `json:"payment"`
	Logs                *[]OrderLogResp   `json:"logs"`
}

type CreateReqOrder struct {
	OrderId         string                `json:"order_id" validate:"required"`
	ReferenceNumber string                `json:"reference_number" validate:"required"`
	OrderType       int                   `json:"order_type" validate:"required"`
	CustomerId      int                   `json:"customer_id" validate:"required"`
	CustomerEmail   string                `json:"customer_email" validate:"required"`
	CustomerName    string                `json:"customer_name" validate:"required"`
	CustomerPhone   string                `json:"customer_phone" validate:"required"`
	ShippingAddress string                `json:"shipping_address" validate:"required"`
	BillingAddress  string                `json:"billing_address" validate:"required"`
	GrandTotal      int64                 `json:"grand_total" validate:"required"`
	SourceChannel   int                   `json:"source_channel" validate:"required"`
	SourceOrder     int                   `json:"source_order" validate:"required"`
	SourceLocation  int                   `json:"source_location" validate:"required"`
	CouponCode      string                `json:"coupon_code"`
	CouponValue     int                   `json:"coupon_value"`
	PaymentId       int64                 `json:"payment_id"`
	OrderItems      []CreateReqOrderItem  `json:"order_items" validate:"required"`
	Payment         CreateReqOrderPayment `json:"payment" validate:"required"`
}
