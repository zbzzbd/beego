package jobclaim

import (
	"models"
	"time"
	"utils"

	"services/progress"

	"services/job"

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
	progress.Register(progress.PT_Claim)
}

type JobClaim struct {
	jobId    uint
	mapJob   map[string]interface{}
	user     *models.User
	job      *models.Job
	jobClaim *models.JobClaim
}

func NewJobClaim(jobId uint, user *models.User) *JobClaim {
	jobClaim := JobClaim{
		jobId:    jobId,
		mapJob:   make(map[string]interface{}),
		user:     user,
		job:      &models.Job{},
		jobClaim: &models.JobClaim{},
	}
	return &jobClaim
}

func (jv *JobClaim) SetCondition(k string, v interface{}) *JobClaim {
	jv.mapJob[k] = v

	return jv
}

func (jv *JobClaim) getJob() (err error) {
	err = models.GetDB().QueryTable(models.TABLE_NAME_JOB).Filter("id", jv.jobId).RelatedSel("CreateUser").One(jv.job)
	if err != nil {
		return
	}
	return
}

func (jv *JobClaim) Do() (id uint, err error) {
	if err = jv.getJob(); err != nil {
		return
	}

	if err = jv.insertJobClaim(); err != nil {
		return
	}

	if err = jv.updateJobStatus(); err != nil {
		return
	}

	if err = jv.addJobProgress(); err != nil {
		return
	}

	if err = jv.sendEmailNotifier(); err != nil {
		return
	}

	return
}

func (jv *JobClaim) insertJobClaim() (err error) {
	jv.jobClaim.Status = uint8(jv.mapJob["Status"].(int))
	jv.jobClaim.Remark = jv.mapJob["Remark"].(string)
	jv.jobClaim.ClaimUser = jv.user
	jv.jobClaim.Job = &models.Job{Id: jv.jobId}

	_, err = models.GetDB().Insert(jv.jobClaim)
	if err != nil {
		utils.GetLog().Debug("services.job.Claim : debug : jobClaim=%s", utils.Sdump(jv.jobClaim))
		return
	}
	return
}

func (jv *JobClaim) updateJobStatus() (err error) {
	q := models.GetDB().QueryTable(models.TABLE_NAME_JOB).Filter("id", jv.jobId)
	newData := orm.Params{
		"claim_time":    time.Now(),
		"claim_status":  jv.jobClaim.Status,
		"submit_status": models.Job_Submit_None,
		"Code":          job.NewJobCode(jv.job.Code, jv.user.Company.Code, job.JobCode_NoneModify).GetCode(),
	}
	_, err = q.Update(newData)
	return err
}

func (jv *JobClaim) addJobProgress() (err error) {
	p := progress.NewProgress(*jv)
	err = p.AddOne()
	return
}

//############################################################
// implements JobProgresser
//############################################################
func (jv JobClaim) GetJobId() uint {
	return jv.jobId
}

func (jv JobClaim) GetTableName() string {
	j := models.JobClaim{}
	return j.TableName()
}

func (jv JobClaim) GetType() progress.ProgressType {
	return progress.PT_Claim
}

func (jv JobClaim) GetDesc() string {
	var desc string = ""
	status := jv.mapJob["Status"].(int)
	switch status {
	case 1:
		desc = "认领作业"
	case 2:
		desc = "拒绝认领作业"
	default:
		desc = "作业未被认领"
	}
	return desc
}

func (jv JobClaim) GetPrimaryKey() uint {
	return jv.jobClaim.Id
}

//############################################################
// implements JobProgresser
//############################################################

func (jc JobClaim) sendEmailNotifier() (err error) {
	n := notify.NewEmail(jc)
	err = n.Send()
	return
}

//############################################################
// implements EmailNotifier
//############################################################
func (jc JobClaim) GetTo() []string {
	return []string{
		jc.job.CreateUser.Email,
	}
}

func (jc JobClaim) GetTemplateName() notify.EmailTemplate {
	return notify.PM_Notify
}

func (jc JobClaim) GetVariables() []map[string]interface{} {
	data := []map[string]interface{}{}
	data = append(data, map[string]interface{}{
		"name":            jc.job.CreateUser.Name,
		"job_name":        jc.job.Type,
		"job_link":        fmt.Sprintf("%s/job/view/%d", utils.BaseUrl(), jc.job.Id),
		"job_status_desc": jc.GetDesc(),
		"job_created_at":  jc.job.Created.Format(utils.YMDHIS),
		"datetime":        time.Now().Format(utils.YMDHIS),
	})
	return data
}

//############################################################
// implements EmailNotifier
//############################################################
