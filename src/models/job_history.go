package models

import "time"

const TABLE_NAME_JOBHISTORY = "job_histories"

type JobHistory struct {
	Id           uint      `orm:"pk;auto"`
	Job          *Job      `orm:"rel(fk)"`
	IsCreate     bool      `orm:"size(1)"`
	Code         string    `orm:"size(32)"`
	Project      *Project  `orm:"rel(fk)"`
	Employee     *User     `orm:"rel(fk)"`
	CreateUser   *User     `orm:"rel(fk)"`
	Type         string    `orm:"size(16)"`
	Department   string    `orm:"size(16)"`
	Target       string    `orm:"size(64)"`
	TargetUrl    string    `orm:"size(64)"`
	ValidTime    time.Time `orm:"type(datetime);null"`
	ClaimTime    time.Time `orm:"type(datetime);null"`
	FinishTime   time.Time `orm:"type(datetime;null)"`
	Desc         string    `orm:"size(4096)"`
	Message      string    `orm:"size(16)"`
	ValidStatus  uint8     `orm:"null"` //审核状态null表示没有被处理过, 0表示审核不通过, 1表示审核通过
	ClaimStatus  uint8     `orm:"null"` //认领状态null表示没有被认领过, 0表示拒绝认领, 1表示认领
	SubmitStatus uint8     `orm"null"`  //提交状态null表示没有被提交过,0表示提交后没被通过, 1表示提交后被接受了

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (j *JobHistory) TableName() string {
	return TABLE_NAME_JOBHISTORY
}

func (j *JobHistory) Update(columns ...string) error {
	affectedRowNum, err := GetDB().Update(j, columns...)
	if affectedRowNum < 1 || err != nil {
		return err
	}
	return nil
}
