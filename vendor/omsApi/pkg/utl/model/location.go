package omsApi

// Location represents company location model
type Location struct {
	TableName struct{} `sql:"location"`

	Base
	Code      string  `sql:"type:varchar(100),notnull" orm:"column(code)" json:"Code"`
	Name      string  `sql:"type:varchar(100),notnull" orm:"column(name)" json:"Name"`
	Flag      int     `sql:",notnull" orm:"column(flag)" json:"Flag"`
	Parent    int     `sql:",default:0" orm:"column(parent)" json:"Parent"`
	Status    int     `sql:",notnull" orm:"column(status)" json:"Status"`
	Latitude  float64 `orm:"column(latitude);digits(15);decimals(10);null" json:"latitude"`
	Longitude float64 `orm:"column(longitude);digits(15);decimals(10);null" json:"longitude"`
}
