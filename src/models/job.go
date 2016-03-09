package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

const TABLE_NAME_JOB = "jobs"

const (
	JOB_EMPLOYEE = "employee"

	JOB_VALID_STATUS  = "valid_status"
	JOB_CLAIM_STATUS  = "claim_status"
	JOB_SUBMIT_STATUS = "submit_status"
	JOB_VALID_TIME    = "valid_time"
	JOB_SUBMIT_TIME   = "submit_time"
	JOB_CLAIM_TIME    = "claim_time"
	JOB_FINISH_TIME   = "finish_time"

	Job_Valid_None   = 0
	Job_Valid_OK     = 1
	Job_Valid_Refuse = 2
	Job_Valid_Cancel = 3

	Job_Claim_None   = 0
	Job_Claim_Ok     = 1
	Job_Claim_Refuse = 2

	Job_Submit_None              = 0
	Job_Submit_Wait_Acceptance   = 1
	Job_Submit_Acceptance_Refuse = 2
	Job_Submit_Acceptance_OK     = 3

	Job_DelStatus_effective = 0 //有效的job
	Job_DelStatus_invalid   = 1 //无效，代表已经删除的job
)

type Job struct {
	Id             uint      `orm:"pk;auto"`
	Code           string    `orm:"size(32)"`
	Project        *Project  `orm:"rel(fk)"`
	Employee       *User     `orm:"rel(fk)"` //作业执行者
	CreateUser     *User     `orm:"rel(fk)"` //作业创建者
	Type           string    `orm:"size(16)"`
	Department     string    `orm:"size(16)"`
	Target         string    `orm:"size(64)"`
	TargetUrl      string    `orm:"size(64)"`
	ValidTime      time.Time `orm:"type(datetime);null"`
	ClaimTime      time.Time `orm:"type(datetime);null"`
	SubmitTime     time.Time `orm:"type(datetime);null"` //实际完成时间
	FinishTime     time.Time `orm:"type(datetime);null"` //要求完成时间
	AcceptanceTime time.Time `orm:"type(datetime);null"` //验收时间
	Desc           string    `orm:"size(4096)"`
	Message        string    `orm:"size(16)"`
	ValidStatus    uint8     `orm:"default(0)"` //审核状态0表示没有被处理过, 1表示审核通过, 2表示审核不通过,重新修改 3表示任务取消
	ClaimStatus    uint8     `orm:"default(0)"` //认领状态0表示没有被认领过, 1表示认领, 2表示拒绝认领
	SubmitStatus   uint8     `orm:"default(0)"` //提交状态0表示没有被提交过, 1表示制作单元提交，等待验收 2 表示验收不通过，重新制作 3表示验收通过
	DelStatus      uint      `orm:"default(0)"` //状态0 表示没有被删除的job   1 表示已经被删除的job

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (j *Job) TableName() string {
	return TABLE_NAME_JOB
}

func (j *Job) Update(columns ...string) error {
	affectedRowNum, err := GetDB().Update(j, columns...)
	if affectedRowNum < 1 || err != nil {
		return err
	}
	return nil
}

func (j *Job) FindBy(conditions map[string]interface{}) error {
	qs := GetDB().QueryTable(j)
	for column, value := range conditions {
		qs = qs.Filter(column, value)
	}
	err := qs.One(j)
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		return fmt.Errorf("job not found: condition = %v", conditions)
	}
	return nil
}

func (j *Job) IsPassedValidation() bool {
	return j.ValidStatus == Job_Valid_OK
}
