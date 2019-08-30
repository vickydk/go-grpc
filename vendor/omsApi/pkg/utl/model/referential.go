package omsApi

import "time"

type Referential struct {
	TableName struct{} `sql:"referential"`

	ID          int       `sql:",pk" json:"id"`
	Code        int       `sql:",notnull" json:"code"`
	SubCode     int       `sql:",notnull" json:"sub_code"`
	TblName     string    `sql:"table_name, type:varchar(100)" json:"tbl_name"`
	FieldName   string    `sql:"type:varchar(100)" json:"field_name"`
	Name        string    `sql:"type:varchar(100)" json:"name"`
	Status      string    `sql:"type:varchar(200)" json:"status"`
	CreatedBy   int       `json:"created_by"`
	CreatedDate time.Time `sql:"type:timestamp, default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedBy   int       ` json:"updated_by"`
	UpdatedDate time.Time `sql:"type:timestamp, default:CURRENT_TIMESTAMP" json:"updated_date"`
	DeletedBy   int       `json:"deletedBy"`
	DeletedDate time.Time `sql:"type:timestamp" json:"deletedDate"`
}

type ReferentialResp struct {
	Code    int32  `json:"code"`
	SubCode int32  `json:"sub_code"`
	Key     string `json:"key"`
	Name    string `json:"name"`
}

type ReferentialReq struct {
	TblName string `json:"flowName"`
	Code    int32  `json:"code"`
	Key     string `json:"key"`
	Name    string `json:"name"`
}
