package models

import "time"

const TABLE_NAME_JOBCOMPLAINT = "job_complaints"

type JobComplaint struct {
	Id          uint      `orm:"pk;auto"`
	Complain    string    `orm:"size(2000)"`
	Source      string    `orm:"size(50);null"`
	Job         *Job      `orm:"rel(fk)"`
	Project     *Project  `orm:"rel(fk)"`
	Employee    *User     `orm:"rel(fk)"`
	User        *User     `orm:"column(complain_user_id);rel(fk)"`
	Type        string    `orm:"size(16)"`
	Response    uint      `orm:"size(5)"`
	ReplyStatus uint      `orm:"size(5);null"`
	DelStatus   uint      `orm:"default(0)"`
	EditStatus  uint      `orm:"default(0)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

func (c *JobComplaint) TableName() string {

	return TABLE_NAME_JOBCOMPLAINT
}
