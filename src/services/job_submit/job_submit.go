package jobsubmit

import (
	"models"
	"time"
	"utils"

	"strconv"

	"services/job"
	"services/progress"

	"fmt"
	"services/notify"

	"github.com/astaxie/beego/orm"
)

func init() {
	//注册事件
	registerJobProgress()
}

func registerJobProgress() {
	//let it crash
	progress.Register(progress.PT_Submit)
}

type JobSubmit struct {
	id     uint
	jobId  uint
	mapJob map[string]interface{}
	user   *models.User
	job    *models.Job
}

func NewJobSubmit(user *models.User) *JobSubmit {
	jobSubmit := JobSubmit{
		mapJob: make(map[string]interface{}),
		user:   user,
		job:    &models.Job{},
	}

	return &jobSubmit
}

func (js *JobSubmit) Do() (id uint, err error) {
	id, err = js.submit()
	js.id = id
	if err == nil {
		js.addJobProgress()
	}
	err = js.sendEmailNotifier()

	return
}

func (js *JobSubmit) getJob() error {
	switch v := js.mapJob["Id"].(type) {
	case int:
		js.jobId = uint(v)
	case string:
		tmp, _ := strconv.Atoi(v)
		js.jobId = uint(tmp)
	}

	err := models.GetDB().QueryTable(models.TABLE_NAME_JOB).Filter("id", js.jobId).RelatedSel("Employee", "CreateUser").One(js.job)

	return err
}

func (js *JobSubmit) submit() (submitId uint, err error) {
	if js.getJob() != nil {
		return
	}

	var jc models.JobSubmit

	jc.Job = js.job
	jc.Status = uint8(js.mapJob["Status"].(int))
	jc.Remark = js.mapJob["Remark"].(string)
	jc.ProduceUser = js.user

	q := models.GetDB().QueryTable(models.TABLE_NAME_JOB).Filter("id", js.jobId)
	newData := orm.Params{
		"submit_status": jc.Status,
	}

	var codeType uint8 = job.JobCode_NoneModify
	if jc.Status == models.Job_Submit_Wait_Acceptance {
		newData["submit_time"] = time.Now()
		if js.job.SubmitStatus == models.Job_Submit_Acceptance_Refuse {
			codeType = job.JobCode_ProduceModify
		}
	} else if jc.Status == models.Job_Submit_Acceptance_OK {
		newData["acceptance_time"] = time.Now()
	}

	newData["Code"] = job.NewJobCode(js.job.Code, js.user.Company.Code, codeType).GetCode()
	_, err = q.Update(newData)

	var insertId int64
	insertId, err = models.GetDB().Insert(&jc)
	if err != nil {
		utils.GetLog().Debug("services.job.Submit : debug : jobSubmit=%s", utils.Sdump(jc))
		return
	}

	submitId = uint(insertId)

	return
}

func (js *JobSubmit) SetCondition(k string, v interface{}) *JobSubmit {
	js.mapJob[k] = v

	return js
}

func (js *JobSubmit) addJobProgress() (err error) {
	p := progress.NewProgress(*js)
	err = p.AddOne()
	return
}

//############################################################
// implements JobProgresser
//############################################################

func (js JobSubmit) GetJobId() uint {
	return js.jobId
}

func (js JobSubmit) GetTableName() string {
	j := models.JobSubmit{}
	return j.TableName()
}

func (js JobSubmit) GetType() progress.ProgressType {
	return progress.PT_Submit
}

func (js JobSubmit) GetDesc() string {
	var desc string = ""
	status := js.mapJob["Status"].(int)
	switch status {
	case models.Job_Submit_Wait_Acceptance:
		desc = "提交作业，等待验收"
	case models.Job_Submit_Acceptance_Refuse:
		desc = "验收作业未通过"
	case models.Job_Submit_Acceptance_OK:
		desc = "验收作业通过"
	default:
		desc = "作业未被提交"
	}
	return desc
}

func (js JobSubmit) GetPrimaryKey() uint {
	return js.id
}

//############################################################
// implements JobProgresser
//############################################################

func (js JobSubmit) sendEmailNotifier() (err error) {
	n := notify.NewEmail(js)
	err = n.Send()
	return
}

//############################################################
// implements EmailNotifier
//############################################################

func (js JobSubmit) GetTo() []string {
	status := js.mapJob["Status"].(int)
	if status > models.Job_Submit_Wait_Acceptance {
		return []string{
			js.job.Employee.Email,
		}
	}

	return []string{
		js.job.CreateUser.Email,
	}
}

func (js JobSubmit) GetTemplateName() notify.EmailTemplate {
	return notify.PM_Notify
}

func (js JobSubmit) GetVariables() []map[string]interface{} {
	var name string
	status := js.mapJob["Status"].(int)
	if status > models.Job_Submit_Wait_Acceptance {
		name = js.job.Employee.Name
	} else {
		name = js.job.CreateUser.Name
	}

	data := []map[string]interface{}{}
	data = append(data, map[string]interface{}{
		"name":            name,
		"job_name":        js.job.Type,
		"job_link":        fmt.Sprintf("%s/job/view/%d", utils.BaseUrl(), js.job.Id),
		"job_status_desc": js.GetDesc(),
		"job_created_at":  js.job.Created.Format(utils.YMDHIS),
		"datetime":        time.Now().Format(utils.YMDHIS),
	})
	return data
}

//############################################################
// implements EmailNotifier
//############################################################
