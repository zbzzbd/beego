package models

import "time"

const TABLE_NAME_JOBASSIGN = "job_assign"

type JobAssign struct {
	Id       uint   `orm:"pk;auto"`
	Job      *Job   `orm:"rel(fk)"`
	FromUser *User  `orm:"rel(fk)"`
	ToUser   *User  `orm:"rel(fk)"`
	Remark   string `orm:"size(255)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (p *JobAssign) TableName() string {
	return TABLE_NAME_JOBASSIGN
}

func (ja *JobAssign) Insert() error {
	lastInsertId, err := GetDB().Insert(ja)
	if err != nil {
		return err
	}
	ja.Id = uint(lastInsertId)
	return nil
}
