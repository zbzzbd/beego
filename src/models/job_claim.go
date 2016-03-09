package models

import "time"

const TABLE_NAME_JOBCLAIM = "job_claims"

type JobClaim struct {
	Id        uint   `orm:"pk;auto"`
	Job       *Job   `orm:"rel(fk)"`
	ClaimUser *User  `orm:"rel(fk)"`
	Status    uint8  `orm:"default(0)"`
	Remark    string `orm:"size(255)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (p *JobClaim) TableName() string {
	return TABLE_NAME_JOBCLAIM
}

func (p *JobClaim) Update(columns ...string) error {
	affectedRowNum, err := GetDB().Update(p, columns...)
	if affectedRowNum < 1 || err != nil {
		return err
	}
	return nil
}
