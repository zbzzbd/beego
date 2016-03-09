package models

import "time"

const TABLE_NAME_JOBVALID = "job_valids"

//作业审核
type JobValid struct {
	Id            uint      `orm:"pk;auto"`
	Job           *Job      `orm:"rel(fk)"`
	OperationUser *User     `orm:"rel(fk)"`
	Status        uint8     `orm:"default(0)"`
	FinishTime    time.Time `orm:"type(datetime);null"` //任务完成时间
	Message       string    `orm:"size(255)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (j *JobValid) TableName() string {
	return TABLE_NAME_JOBVALID
}

func (jv *JobValid) Insert() error {
	lastInsertId, err := GetDB().Insert(jv)
	if err != nil {
		return err
	}
	jv.Id = uint(lastInsertId)
	return nil
}
