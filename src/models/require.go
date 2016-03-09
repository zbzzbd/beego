package models

import "time"

const TABLE_NAME_REQUIRE = "requires"

type Require struct {
	Id             uint      `orm:"pk;auto"`
	Name           string    `orm:"size(128)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (j *Require) TableName() string {
	return TABLE_NAME_REQUIRE
}
