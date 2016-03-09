package models

import "time"

const TABLE_NAME_JOBPROGRESS = "job_progresses"

type JobProgress struct {
	Id             uint `orm:"pk;auto"`
	JobId          uint `orm:"index"`
	ProgressType   uint
	Desc           string `orm:"size(32)"`
	EventTableName string `orm:"size(32)"` //table name
	PrimaryKey     uint

	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (p *JobProgress) TableName() string {
	return TABLE_NAME_JOBPROGRESS
}

func (p *JobProgress) Insert() error {
	lastInsertId, err := GetDB().Insert(p)
	if err != nil {
		return err
	}
	p.Id = uint(lastInsertId)
	return nil
}
