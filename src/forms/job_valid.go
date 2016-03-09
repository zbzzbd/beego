package forms

import (
	"fmt"
	"net/http"
	"time"
	"utils"

	"github.com/astaxie/beego/validation"
	"github.com/mholt/binding"
)

type JobValidForm struct {
	JobId              uint
	Result             uint8  //审核结果 0, 1, 2
	RequiredFinishTime string //作业要求的完成时间
	FinishTime         string
	Message            string
	CurrentUserId      uint

	valid *validation.Validation
	Base
}

func NewJobValidForm(req *http.Request) (*JobValidForm, error) {
	jv := new(JobValidForm)

	//binding
	errs := binding.Bind(req, jv)
	if errs.Len() > 0 {
		return nil, fmt.Errorf("%s", errs.Error())
	}

	//validator
	if err := jv.Valid(); err != nil {
		return nil, err
	}

	return jv, nil
}

func (jv *JobValidForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&jv.JobId:              "job_id",
		&jv.Result:             "result",
		&jv.FinishTime:         "finish_time",
		&jv.RequiredFinishTime: "job_required_finish_time",
		&jv.Message:            "message",
	}
}

func (jv *JobValidForm) Valid() (err error) {
	jv.valid = &validation.Validation{}

	if err = jv.validJobId(); err != nil {
		return
	}

	if err = jv.validResult(); err != nil {
		return
	}

	if err = jv.validFinishTime(); err != nil {
		return
	}

	if err = jv.validMessage(); err != nil {
		return
	}

	return
}

func (jv *JobValidForm) validJobId() error {
	if jv.JobId == 0 {
		return fmt.Errorf("jobId错误")
	}
	return nil
}

func (jv *JobValidForm) validResult() error {
	results := [...]uint8{1, 2, 3}
	for _, r := range results {
		if jv.Result == r {
			return nil
		}
	}
	return fmt.Errorf("审核结果选项不正确, %d", jv.Result)
}

func (jv *JobValidForm) validFinishTime() error {
	//如果审核不通过, 则不作判断
	if jv.Result != 1 {
		return nil
	}

	finishTime, err := utils.FormatTime(jv.FinishTime)
	if err != nil {
		return fmt.Errorf("finisht_time数据错误, %s", jv.FinishTime)
	}

	requiredFinishTime, err := utils.FormatTime(jv.RequiredFinishTime)
	if err != nil {
		return fmt.Errorf("job_required_finisht_time数据错误, %s", jv.RequiredFinishTime)
	}

	if finishTime.After(requiredFinishTime) {
		return fmt.Errorf("要求完成时间必须小于作业完成时间, %s, %s", jv.FinishTime, jv.RequiredFinishTime)
	}

	if finishTime.Before(time.Now()) {
		return fmt.Errorf("要求完成时间必须大于当前时间, %s", jv.FinishTime)
	}

	return nil
}

func (jv *JobValidForm) validMessage() error {
	if len(jv.Message) > 255 {
		return fmt.Errorf("补充说明过长: %s", jv.Message)
	}
	return nil
}
