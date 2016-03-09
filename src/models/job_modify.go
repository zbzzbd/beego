package models

import "time"

const TABLE_NAME_JOBMODIFY = "job_modifys"

type JobModify struct {
	Id       uint        `orm:"pk;auto"`
	Job      *Job        `orm:"rel(fk)"`
	EditUser *User       `orm:"rel(fk)"`
	History  *JobHistory `orm:"rel(fk)"`
	Remark   string      `orm:"size(32)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (j *JobModify) TableName() string {
	return TABLE_NAME_JOBMODIFY
}
