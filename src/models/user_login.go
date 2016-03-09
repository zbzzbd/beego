package models

import "time"

type UserLogin struct {
	Id      uint      `orm:"pk;auto"`
	User    *User     `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (u *UserLogin) TableName() string {
	return "user_logins"
}

func (u *UserLogin) Insert() error {
	lastInsertId, err := GetDB().Insert(u)
	if err != nil {
		return err
	}
	u.Id = uint(lastInsertId)
	return nil
}
