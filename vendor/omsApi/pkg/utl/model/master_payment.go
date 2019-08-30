package omsApi

import "time"

type Master_payment struct {
	TableName struct{} `sql:"master_payment"`

	Base
	Name          string    `sql:"type:varchar(100)" orm:"column(name)" json:"name"`
	Description   string    `orm:"column(description)" json:"Description"`
	Type          int       `orm:"column(type)" json:"Type"`
	MaxGrandTotal int64     `orm:"column(maximum_grand_total)" json:"MaxGrandTotal"`
	MinGrandTotal int64     `orm:"column(minimum_grand_total)" json:"MinGrandTotal"`
	Status        int       `orm:"column(status)" json:"Status"`
	ShowOrder     int       `orm:"column(show_order)" json:"ShowOrder"`
	Group         int       `orm:"column(group)" json:"Group"`
	CreatedDate   time.Time `sql:"type:timestamp" orm:"column(created_date);null" json:"CreatedDate"`
	LastModified  time.Time `sql:"type:timestamp, default:CURRENT_TIMESTAMP" orm:"column(last_modified);null" json:"LastModified"`
}
