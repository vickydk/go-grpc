package omsApi

const (
	TblOrderPayment string = "order_payment"
)

type Order_payment struct {
	Id               int    `json:"id"`
	MasterPaymentId  int    `json:"master_payment_id"`
	Amount           int64  `json:"amount"`
	TransactionDate  string `json:"transaction_date"`
	Currency         int    `json:"currency"`
	Status           int    `json:"status"`
	Active           int    `json:"active"`
	BinNumber        string `json:"bin_number"`
	ConfirmTrfBy     string `json:"confirm_trf_by"`
	ConfirmTrfBank   string `json:"confirm_trf_bank"`
	ConfirmTrfAmount string `json:"confirm_trf_amount"`
	CreatedBy        int    `json:"created_by"`
	CreatedDate      string `json:"created_date"`
	LastModified     string `json:"last_modified"`
	ReferenceNumber  string `json:"reference_number"`
}

type OrderPaymentReq struct {
	ReferenceNumber  string `json:"reference_number,omitempty"`
	BinNumber        string `json:"bin_number,omitempty"`
	ConfirmTrfBy     string `json:"confirm_trf_by,omitempty"`
	ConfirmTrfBank   string `json:"confirm_trf_bank,omitempty"`
	ConfirmTrfAmount string `json:"confirm_trf_amount,omitempty"`
	PaymentStatus    int    `json:"payment_status,omitempty"`
	Action           int    `json:"action,omitempty"`
}

type OrderPaymentResp struct {
	Id                int     `json:"id"`
	ReferenceNumber   string  `json:"reference_number"`
	MasterPaymentName string  `json:"master_payment_name"`
	Amount            float64 `json:"amount"`
	DiscAmount        float64 `json:"disc_amount"`
	PaidAmount        float64 `json:"paid_amount"`
	TransactionDate   string  `json:"transaction_date"`
	Currency          int     `json:"currency"`
	CurrencyName      string  `json:"currency_name"`
	Status            int     `json:"status"`
	StatusName        string  `json:"status_name"`
	Active            int     `json:"active"`
	BinNumber         string  `json:"bin_number"`
	ConfirmTrfBy      string  `json:"confirm_trf_by"`
	ConfirmTrfBank    string  `json:"confirm_trf_bank"`
	ConfirmTrfAmount  string  `json:"confirm_trf_amount"`
}

type CreateReqOrderPayment struct {
	ReferenceNumber string `json:"reference_number"`
	MasterPaymentId int    `json:"master_payment_id" validate:"required"`
	Amount          int64  `json:"amount" validate:"required"`
	Currency        int    `json:"currency" validate:"required"`
}
