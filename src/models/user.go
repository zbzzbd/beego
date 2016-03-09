package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

const TABLE_NAME_USER = "users"

type User struct {
	Id            uint     `orm:"pk;auto"`
	Name          string   `orm:"size(32)"`
	Email         string   `orm:"size(128)"`
	Mobile        string   `orm:"size(16);null"` //添加手机号码用于将来短信通知
	Password      string   `orm:"size(32)"`
	Salt          string   `orm:"size(6)"`
	Avatar        string   `orm:"size(255)"`
	Company       *Company `orm:"rel(fk)"`
	Roles         string   `orm:"size(255)"`
	LastLoginTime time.Time

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Deleted time.Time `orm:"type(datetime);null"`
}

func (u *User) TableName() string {
	return TABLE_NAME_USER
}

func (u *User) FindBy(conditions map[string]interface{}) error {
	qs := GetDB().QueryTable(u)
	for column, value := range conditions {
		qs = qs.Filter(column, value)
	}
	err := qs.One(u)
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		return fmt.Errorf("user not found: condition = %v", conditions)
	}
	return nil
}

func (u *User) Update(columns ...string) error {
	affectedRowNum, err := GetDB().Update(u, columns...)
	if affectedRowNum < 1 || err != nil {
		return err
	}
	return nil
}

func (u *User) Insert() error {
	lastInsertId, err := GetDB().Insert(u)
	if err != nil {
		return err
	}
	u.Id = uint(lastInsertId)
	return nil
}

func (u *User) IsDeleted() bool {
	return !u.Deleted.IsZero()
}
