package models

import "time"

const TABLE_NAME_COMPANY = "companies"

type Company struct {
	Id      uint   `orm:"pk;auto"`
	Code    string `orm:"unique"`
	Name    string `orm:"size(16)"`
	Address string `orm:"size(128)"`
	City    string `orm:"size(16)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (c *Company) TableName() string {
	return TABLE_NAME_COMPANY
}
