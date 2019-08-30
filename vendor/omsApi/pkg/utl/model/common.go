package omsApi

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// Predefined model error codes.
const (
	ErrNone       = 200
	ErrDatabase   = 507
	ErrExtSystem  = 508
	ErrSystem     = 507
	ErrDupRows    = 406
	ErrNotFound   = 402
	ErrReqInvalid = 400
	ErrConflict   = 409
)

// response struct
type Response struct {
	Content interface{} `json:"content"`
	Message string      `json:"Message"`
	Code    int         `json:"code"`
}

type ListResponse struct {
	Body interface{} `json:"body"`
	Page interface{} `json:"pagination"`
}

type FilterOrder struct {
	OrderID         string `json:"order_id" query:"order_id"`
	ReferenceNumber string `json:"reference_number" query:"reference_number"`
	SourceLocation  int    `json:"source_location,omitempty" query:"source_location"`
	OrderDateFrom   string `json:"order_date_from,omitempty" query:"order_date_from"`
	OrderDateTo     string `json:"order_date_to,omitempty" query:"order_date_to"`
	SKU             string `json:"sku,omitempty" query:"sku"`
	CustomerEmail   string `json:"customer_email,omitempty" query:"customer_email"`
	CustomerPhone   string `json:"customer_phone,omitempty" query:"customer_phone"`
	CouponCode      string `json:"coupon_code,omitempty" query:"coupon_code"`
	SourceChannel   []int  `json:"source_channel,omitempty" query:"source_channel"`
	SourceOrder     []int  `json:"source_order,omitempty" query:"source_order"`
	Currency        int    `json:"currency,omitempty" query:"currency"`
	Status          int    `json:"status" query:"status"`
	SubStatus       int    `json:"status" query:"sub_status"`
	FilterBy        int    `json:"filter_by" query:"filter_by"`
	PaginationReq
}

type RespondViewOrder struct {
	Order_header RespondOrderHeader `json:"order_header"`
	Order_item   []Order_item       `json:"order_item"`
}

// @Title
// @Description
func Respond(err string, result interface{}, code int) *Response {
	fmt.Printf("err: %v", err)
	resp := &Response{
		Content: result,
		Code:    code,
		Message: err,
	}

	return resp
}

func ListRepond(body interface{}, pagination interface{}) *ListResponse {
	resp := &ListResponse{
		Body: body,
		Page: pagination,
	}

	return resp
}

func MappingError(err error) string {
	if err == pg.ErrNoRows {
		return "Data not found"
	} else if err == errors.New(string(http.StatusPaymentRequired)) {
		return "Payment Required"
	} else if err == errors.New(string(ErrReqInvalid)) {
		return "Missing request"
	}
	return fmt.Sprint("Belum di mapping: ", err)
}

func MapReferential(data []ReferentialResp, fieldNmae string, code int, sub_code int, name string) ReferentialResp {
	ret := ReferentialResp{}
	for _, eachrefs := range data {
		if eachrefs.Key == fieldNmae && eachrefs.Code == int32(code) && eachrefs.SubCode == int32(sub_code) {
			return eachrefs
			break
		}
		if eachrefs.Key == fieldNmae && strings.ToLower(strings.ReplaceAll(eachrefs.Name, " ", "")) == strings.ToLower(strings.ReplaceAll(name, " ", "")) {
			return eachrefs
			break
		}
	}
	return ret
}

func CheckJsonField(data interface{}, name string) bool {
	flag := false
	s := reflect.ValueOf(data)
	for i := 0; i < s.Elem().NumField(); i++ {
		v := s.Elem().Field(i).String()
		fieldName := s.Elem().Type().Field(i).Name
		if _, err := strconv.Atoi(v); err != nil {
			if len(v) > 0 && fieldName == name {
				flag = true
				break
			} else {
				continue
			}
		} else if fieldName == name {
			flag = true
			break
		}
	}
	return flag
}
