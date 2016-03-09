package models

import "time"

const TABLE_NAME_JOBSUBMIT = "job_submits"

type JobSubmit struct {
	Id          uint   `orm:"pk;auto"`
	Job         *Job   `orm:"rel(fk)"`
	ProduceUser *User  `orm:"rel(fk)"`
	Status      uint8  `orm:"default(0)"`
	Remark      string `orm:"size(255)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (p *JobSubmit) TableName() string {
	return TABLE_NAME_JOBSUBMIT
}
