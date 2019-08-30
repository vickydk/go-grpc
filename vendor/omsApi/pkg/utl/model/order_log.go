package omsApi

import (
	"omsApi/pkg/utl/constants"
	"strings"
)

const (
	TblOrderLog string = "order_logs"
)

type Order_logs struct {
	ID            int    `json:"id"`
	CreatedDate   string `json:"created_date"`
	CreatedBy     int    `json:"created_by"`
	OrderHeaderId int64  `json:"order_header_id"`
	Log           string `json:"log"`
	Reason        string `json:"reason"`
	Comment       string `json:"comment"`
}

type CreateOrderLogReq struct {
	Status              int    `json:"status"`
	SubStatus           int    `json:"sub_status"`
	Action              string `json:"action"`
	OrderId             string `json:"order_id"`
	OrderHeaderId       int64  `json:"order_header_id"`
	FulfillmentLocation int    `json:"fulfillment_location"`
	CreatedBy           int    `json:"created_by"`
	Reason              string `json:"reason"`
	Comment             string `json:"comment"`
}

type OrderLogResp struct {
	CreatedDate string `json:"created_date"`
	CreatedBy   int    `json:"created_by"`
	OrderId     string `json:"order_id"`
	Log         string `json:"log"`
	Reason      string `json:"reason"`
	Comment     string `json:"comment"`
}

func ConvertLog(req *CreateOrderLogReq, refs []ReferentialResp) string {
	var log strings.Builder
	status := "status"
	if req.Action == "create" {
		log.WriteString(constants.CreateOrder)
	} else if req.Action == "return" {
		log.WriteString(constants.CreateOrderReturn)
	} else if req.Action == "reorder" {
		log.WriteString(constants.CreateReOrder)
	} else if req.Action == "pickbycustomer" {
		log.WriteString(constants.ReadyToPick)
	} else if req.Action == "assigncourier" {
		log.WriteString(constants.AssaignCourier)
	} else {
		log.WriteString(constants.MoveOrder)
		if req.Action == "outofdelivery" {
			log.WriteString(constants.SubsattusFrom)
		} else if req.Action == "intransti" {
			log.WriteString("- In transit ")
		} else if req.Action == "auth" || req.Action == "wfp" {
			log.WriteString(constants.SubstatusAt)
		}
	}

	logs := log.String()
	logs = strings.ReplaceAll(logs, "<Order_ID>", req.OrderId)
	if req.SubStatus > 0 {
		status = "sub_status"
	}
	logs = strings.ReplaceAll(logs, "<Status>", MapReferential(refs, status, req.Status, req.SubStatus, "").Name)
	logs = strings.ReplaceAll(logs, "<fulfillment_location>", string(req.FulfillmentLocation))

	return logs
}
