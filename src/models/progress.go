package models

import "time"

const TABLE_NAME_PROGRESS = "progress"

type Progress struct {
	Id      uint      `orm:"pk;auto"`
	Name    string    `orm:"size(128);unique"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (p *Progress) TableName() string {
	return TABLE_NAME_PROGRESS
}

func GetProgressList() (err error, progressList []*Progress) {
	_, err = GetDB().QueryTable("progress").All(&progressList)
	return
}
func (p *Progress) GetProgressById(id string) (err error) {
	err = GetDB().QueryTable("progress").Filter("id", id).One(p)
	return err

}
