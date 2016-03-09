package job

import (
	"models"
	"services/progress"
	"strconv"
	"utils"

	"fmt"
	"services/notify"
	"time"

	"github.com/astaxie/beego/orm"
)

type JobModify struct {
	id       uint
	jobId    uint
	mapJob   map[string]interface{}
	codeType uint8
	user     *models.User
	job      *models.Job
}

func NewJobModify(user *models.User, codeType uint8) *JobModify {
	jobModify := JobModify{
		mapJob:   make(map[string]interface{}),
		user:     user,
		codeType: codeType,
		job:      &models.Job{},
	}

	return &jobModify
}

func (jv *JobModify) Do() (modifyId uint, err error) {
	modifyId, err = jv.update()
	jv.id = modifyId
	if err == nil {
		jv.addJobProgress()

	}

	err = jv.sendEmailNotifier()

	return
}

func (jv *JobModify) getJob() error {
	switch v := jv.mapJob["Id"].(type) {
	case int:
		jv.jobId = uint(v)
	case string:
		tmp, _ := strconv.Atoi(v)
		jv.jobId = uint(tmp)
	}

	err := models.GetDB().QueryTable(models.TABLE_NAME_JOB).Filter("id", jv.jobId).RelatedSel("Employee", "CreateUser").One(jv.job)

	return err
}

func (jv *JobModify) update() (modifyId uint, err error) {
	if jv.getJob() != nil {
		return
	}

	q := models.GetDB().QueryTable(models.TABLE_NAME_JOB).Filter("id", jv.jobId)

	newData := orm.Params{"claim_status": models.Job_Claim_None}

	// 审核未通过，编辑修改
	if jv.job.ValidStatus == models.Job_Valid_Refuse {
		newData["valid_status"] = models.Job_Valid_None
	}

	for k, v := range jv.mapJob {
		if k == "FinishTime" {
			finishTime, _ := utils.FormatTime(jv.mapJob["FinishTime"].(string))
			newData[utils.SepSplitName(k, "")] = finishTime
		} else {
			newData[utils.SepSplitName(k, "")] = v
		}
	}
	newData["Code"] = NewJobCode(jv.job.Code, jv.user.Company.Code, jv.codeType).GetCode()
	_, err = q.Update(newData)

	if err != nil {
		return
	}

	err = jv.getJob()
	if err == nil {

		modifyId, _ = AddHistory(jv.job, jv.user, false)
	}
	return
}

func (jv *JobModify) SetCondition(k string, v interface{}) *JobModify {
	jv.mapJob[k] = v

	return jv
}

func (jv *JobModify) addJobProgress() (err error) {
	p := progress.NewProgress(*jv)
	err = p.AddOne()
	return
}

//############################################################
// implements JobProgresser
//############################################################
func (jv JobModify) GetJobId() uint {
	return jv.jobId
}

func (jv JobModify) GetTableName() string {
	j := models.JobHistory{}
	return j.TableName()
}

func (jv JobModify) GetType() progress.ProgressType {
	return progress.PT_Modify
}

func (jv JobModify) GetDesc() string {
	return "修改作业"
}

func (jv JobModify) GetPrimaryKey() uint {
	return jv.id
}

//############################################################
// implements JobProgresser
//############################################################

func (self JobModify) sendEmailNotifier() (err error) {
	n := notify.NewEmail(self)
	err = n.Send()
	return
}

//############################################################
// implements EmailNotifier
//############################################################
func (jm JobModify) GetTo() []string {
	email := jm.job.CreateUser.Email

	if jm.job.ValidStatus == models.Job_Valid_OK {
		email = jm.job.Employee.Email
	}
	return []string{
		email,
	}
}

func (jm JobModify) GetTemplateName() notify.EmailTemplate {
	return notify.PM_Notify
}

func (jm JobModify) GetVariables() []map[string]interface{} {
	name := jm.job.CreateUser.Name
	if jm.job.ValidStatus == models.Job_Valid_OK {
		name = jm.job.Employee.Name
	}

	data := []map[string]interface{}{}
	data = append(data, map[string]interface{}{
		"name":            name,
		"job_name":        jm.job.Type,
		"job_link":        fmt.Sprintf("%s/job/view/%d", utils.BaseUrl(), jm.job.Id),
		"job_status_desc": jm.GetDesc(),
		"job_created_at":  jm.job.Created.Format(utils.YMDHIS),
		"datetime":        time.Now().Format(utils.YMDHIS),
	})
	return data
}

//############################################################
// implements EmailNotifier
//############################################################
