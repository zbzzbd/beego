package job

import (
	"models"
	"utils"

	"fmt"
	"services/notify"
	"services/progress"
	"time"
)

type JobCreator struct {
	id     uint
	jobId  uint
	mapJob map[string]interface{}
	user   *models.User
	job    *models.Job
}

func NewJobCreator(user *models.User) *JobCreator {
	return &JobCreator{
		mapJob: make(map[string]interface{}),
		user:   user,
	}
}

func (jv *JobCreator) Do() (jid, historyId uint, err error) {
	jv.jobId, jv.id, err = jv.create()
	jid = jv.jobId
	historyId = jv.id
	if err == nil {
		jv.addJobProgress()
	}

	err = jv.sendEmailNotifier()

	return
}

func (jv *JobCreator) create() (jid, historyId uint, err error) {
	var j models.Job
	j.Code = NewJobCode("", jv.user.Company.Code, JobCode_NoneModify).GetCode()
	j.CreateUser = jv.user
	pid, _ := jv.mapJob["ProjectId"].(int)
	j.Project = &models.Project{Id: uint(pid)}
	uid, _ := jv.mapJob["EmployeeId"].(int)
	if uid != 0 {
		j.Employee = &models.User{Id: uint(uid)}
	}
	j.Type = jv.mapJob["Type"].(string)
	j.Department = jv.mapJob["Department"].(string)
	j.Target = jv.mapJob["Target"].(string)
	j.TargetUrl = jv.mapJob["TargetUrl"].(string)
	j.Desc = jv.mapJob["Desc"].(string)
	j.Message = jv.mapJob["Message"].(string)
	finishTime, _ := utils.FormatTime(jv.mapJob["FinishTime"].(string))
	j.FinishTime = finishTime

	var insertId int64
	insertId, err = models.GetDB().Insert(&j)
	if err != nil {
		utils.GetLog().Debug("services.job.Create : debug : job=%s", utils.Sdump(j))
		return
	}
	historyId, _ = AddHistory(&j, jv.user, true)

	jid = uint(insertId)

	jv.job = &j

	return
}
func (jv *JobCreator) SetCondition(k string, v interface{}) *JobCreator {
	jv.mapJob[k] = v

	return jv
}

func (jc *JobCreator) getJobProjectRegisterUser() (err error) {
	if jc.job.Project.Name != "" {
		return nil
	}

	projectId := jc.job.Project
	err = models.GetDB().QueryTable(models.TABLE_NAME_PROJECT).Filter("id", projectId).RelatedSel("Registrant").One(jc.job.Project)
	return
}

func (jv *JobCreator) addJobProgress() (err error) {
	p := progress.NewProgress(*jv)
	err = p.AddOne()
	return
}

//############################################################
// implements JobProgresser
//############################################################
func (jv JobCreator) GetJobId() uint {
	return jv.jobId
}

func (jv JobCreator) GetTableName() string {
	j := models.JobHistory{}
	return j.TableName()
}

func (jv JobCreator) GetType() progress.ProgressType {
	return progress.PT_Create
}

func (jv JobCreator) GetDesc() string {
	return "作业登记"
}

func (jv JobCreator) GetPrimaryKey() uint {
	return jv.id
}

//############################################################
// implements JobProgresser
//############################################################

func (jc JobCreator) sendEmailNotifier() (err error) {
	n := notify.NewEmail(jc)
	err = n.Send()
	return
}

//############################################################
// implements EmailNotifier
//############################################################
func (jc JobCreator) GetTo() []string {
	err := (&jc).getJobProjectRegisterUser()
	if err != nil {
		//handle error
	}

	return []string{
		(&jc).job.Project.Registrant.Email,
	}
}

func (jc JobCreator) GetTemplateName() notify.EmailTemplate {
	return notify.PM_Notify
}

func (jc JobCreator) GetVariables() []map[string]interface{} {
	err := (&jc).getJobProjectRegisterUser()
	if err != nil {
		//handle error
	}

	data := []map[string]interface{}{}
	data = append(data, map[string]interface{}{
		"name":            (&jc).job.Project.Registrant.Name,
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
