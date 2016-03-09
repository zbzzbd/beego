package jobvalid

import (
	"forms"
	"models"
	"services/progress"
	"time"
	"utils"

	"fmt"
	"services/notify"

	"github.com/astaxie/beego/orm"
)

var jobValidationResult map[int]string

func init() {
	jobValidationResult = map[int]string{
		1: "内容明确, 作业排产",
		2: "作业有误, 请重新修改",
		3: "作业取消",
	}

	//注册事件
	registerJobProgress()
}

func GetJobValidationResult() map[int]string {
	return jobValidationResult
}

func registerJobProgress() {
	//let it crash
	progress.Register(progress.PT_Valid)
}

type JobValidation struct {
	jobId         uint
	userId        uint
	result        uint8
	finishTime    string
	msg           string
	jobValidation *models.JobValid
}

func NewJobValidation(jvf *forms.JobValidForm) *JobValidation {
	return &JobValidation{
		jobId:         jvf.JobId,
		userId:        jvf.CurrentUserId,
		result:        jvf.Result,
		finishTime:    jvf.FinishTime,
		msg:           jvf.Message,
		jobValidation: &models.JobValid{},
	}
}

func (jv *JobValidation) getJob() (*models.Job, error) {
	//判断job的状态是否被修改过
	job := &models.Job{Id: jv.jobId}
	err := models.GetDB().QueryTable(models.TABLE_NAME_JOB).Filter("id", jv.jobId).RelatedSel("Employee", "CreateUser").One(job)
	return job, err
}

func (jv *JobValidation) getUser() (*models.User, error) {
	user := &models.User{Id: jv.userId}
	return user, nil
}

func (jv *JobValidation) getResult() (uint8, error) {
	return jv.result, nil
}

func (jv *JobValidation) getFinishTime() (time.Time, error) {
	return utils.FormatTime(jv.finishTime)
}

func (jv *JobValidation) getMessage() (string, error) {
	return jv.msg, nil
}

func (jv *JobValidation) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("job.JobValidation.Do : error = %s, ojb= %s", err.Error(), utils.Sdump(jv))
		}
		utils.GetLog().Debug("job.JobValidation.Do : ojb= %s", utils.Sdump(jv))
	}()

	if jv.jobValidation.Job, err = jv.getJob(); err != nil {
		return
	}

	if jv.jobValidation.OperationUser, err = jv.getUser(); err != nil {
		return
	}

	if jv.jobValidation.Status, err = jv.getResult(); err != nil {
		return
	}

	if jv.jobValidation.Message, err = jv.getMessage(); err != nil {
		return
	}

	jv.jobValidation.FinishTime, err = jv.getFinishTime()
	if jv.jobValidation.Status == 0 && err != nil {
		return
	}

	if err = jv.jobValidation.Insert(); err != nil {
		return
	}

	if err = jv.afterValidation(); err != nil {
		return
	}

	return
}

func (jv *JobValidation) afterValidation() (err error) {
	if err = jv.updateJobStatus(); err != nil {
		return
	}

	if err = jv.updateJobHistory(); err != nil {
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

func (jv *JobValidation) updateJobStatus() (err error) {
	//更新job状态
	jv.jobValidation.Job.ValidStatus, _ = jv.getResult()
	jv.jobValidation.Job.ValidTime = time.Now()

	//如果审核通过
	if jv.jobValidation.Job.ValidStatus == models.Job_Valid_OK {
		//修改作业完成时间
		jv.jobValidation.Job.FinishTime, _ = jv.getFinishTime()
	}

	if err = jv.jobValidation.Job.Update(models.JOB_VALID_STATUS, models.JOB_VALID_TIME, models.JOB_FINISH_TIME); err != nil {
		return
	}
	return
}

func (jv *JobValidation) updateJobHistory() (err error) {
	if jv.jobValidation.Job.ValidStatus != models.Job_Valid_OK {
		return
	}

	finishTime, _ := jv.getFinishTime()
	params := orm.Params{
		"finish_time": finishTime,
	}

	_, err = models.GetDB().QueryTable(models.TABLE_NAME_JOBHISTORY).Filter("job_id", jv.jobId).Update(params)

	return
}

func (jv *JobValidation) addJobProgress() (err error) {
	p := progress.NewProgress(*jv)
	err = p.AddOne()
	return
}

//############################################################
// implements JobProgresser
//############################################################
func (jv JobValidation) GetJobId() uint {
	return jv.jobId
}

func (jv JobValidation) GetTableName() string {
	return jv.jobValidation.TableName()
}

func (jv JobValidation) GetType() progress.ProgressType {
	return progress.PT_Valid
}

func (jv JobValidation) GetDesc() string {
	var desc string = "审核回复"
	switch jv.jobValidation.Job.ValidStatus {
	case models.Job_Valid_OK:
		desc += " - 通过"
	case models.Job_Valid_Refuse:
		desc += " - 拒绝"
	case models.Job_Valid_Cancel:
		desc += " - 作业取消"
	}
	return desc
}

func (jv JobValidation) GetPrimaryKey() uint {
	return jv.jobValidation.Id
}

//############################################################
// implements JobProgresser
//############################################################

func (jv JobValidation) sendEmailNotifier() (err error) {
	n := notify.NewEmail(jv)
	err = n.Send()
	return
}

//############################################################
// implements EmailNotifier
//############################################################
func (jv JobValidation) GetTo() (emailList []string) {
	bussinessUser := jv.jobValidation.Job.CreateUser
	emailList = append(emailList, bussinessUser.Email)

	if jv.jobValidation.Job.IsPassedValidation() {
		produceUser := jv.jobValidation.Job.Employee
		emailList = append(emailList, produceUser.Email)
	}
	return
}

func (jv JobValidation) GetTemplateName() notify.EmailTemplate {
	return notify.PM_Notify
}

func (jv JobValidation) GetVariables() []map[string]interface{} {
	job := jv.jobValidation.Job

	data := []map[string]interface{}{}
	data = append(data, map[string]interface{}{
		"name":            job.CreateUser.Name,
		"job_name":        job.Type,
		"job_link":        fmt.Sprintf("%s/job/view/%d", utils.BaseUrl(), job.Id),
		"job_status_desc": jv.GetDesc(),
		"job_created_at":  jv.jobValidation.Job.Created.Format(utils.YMDHIS),
		"datetime":        time.Now().Format(utils.YMDHIS),
	})

	if job.IsPassedValidation() {
		data = append(data, map[string]interface{}{
			"name":            job.Employee.Name,
			"job_name":        job.Type,
			"job_link":        fmt.Sprintf("%s/job/view/%d", utils.BaseUrl(), job.Id),
			"job_status_desc": jv.GetDesc(),
			"job_created_at":  jv.jobValidation.Job.Created.Format(utils.YMDHIS),
			"datetime":        time.Now().Format(utils.YMDHIS),
		})
	}

	return data
}

//############################################################
// implements EmailNotifier
//############################################################
