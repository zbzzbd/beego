package models

import (
	"time"
)

const TABLE_NAME_JOBCOMPLAINTREPLY = "job_complaint_replies"

type JobComplaintReply struct {
	Id        uint          `orm:"pk;auto"`
	Message   string        `orm:"size(255)"`
	Complaint *JobComplaint `orm:"rel(fk)"`
	User      *User         `orm:"rel(fk)"`
	Created   time.Time     `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time     `orm:"auto_now;type(datetime)"`
}

func (r *JobComplaintReply) TableName() string {
	return TABLE_NAME_JOBCOMPLAINTREPLY
}
