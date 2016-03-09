package complaint

import (
	"fmt"
	"models"
	"services/notify"
	"services/progress"
	"strconv"
	"time"
	"utils"
)

func (cp *ComplaintCreator) CreateComplian(mapComplain map[string]interface{}, user *models.User) (complainId, jobId uint, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.complain.Create : error : %s, ", err.Error())
		} else {
			utils.GetLog().Debug("services.complain.Create : debug : mapComplain=%s", utils.Sdump(mapComplain))
		}
	}()

	var complian models.JobComplaint
	pid, _ := mapComplain["JobId"].(string)

	filter := make(map[string]interface{})
	filter["id"] = pid

	var job *models.Job
	job, err = GetOneJob(filter)
	complian.Job = &models.Job{Id: job.Id}
	complian.Project = &models.Project{Id: job.Project.Id}

	if mapComplain["employee_id"] != nil {
		employId, _ := strconv.Atoi(mapComplain["employee_id"].(string))
		complian.Employee = &models.User{Id: uint(employId)}
	}
	complian.Complain = mapComplain["Complain"].(string)
	complian.Source = mapComplain["source"].(string)
	complian.Type = mapComplain["Type"].(string)
	response, _ := mapComplain["response"].(int)
	complian.Response = uint(response)
	if mapComplain["editstatus"] != nil {
		editStatus, _ := strconv.Atoi(mapComplain["editstatus"].(string))
		complian.EditStatus = uint(editStatus)
	}

	complian.User = user

	//cp.compliant = complian
	var insertId int64
	insertId, err = models.GetDB().Insert(&complian)
	if err != nil {
		utils.GetLog().Debug("services.complian.Create : debug : job=%s", utils.Sdump(complian))
		return
	}
	complainId = uint(insertId)
	jobId = uint(job.Id)
	return

}

type ComplaintCreator struct {
	id           uint
	jobId        uint
	mapComplaint map[string]interface{}
	user         *models.User
	compliant    *models.JobComplaint
	employee     *models.User
}

func NewcomplainCreator(mapComplaint map[string]interface{}, user *models.User) *ComplaintCreator {
	return &ComplaintCreator{
		mapComplaint: mapComplaint,
		user:         user,
		compliant:    &models.JobComplaint{},
		employee:     &models.User{},
	}
}

func (cp *ComplaintCreator) Do() (complaintId, jid uint, err error) {
	cp.id, cp.jobId, err = cp.CreateComplian(cp.mapComplaint, cp.user)
	jid = cp.jobId
	complaintId = cp.id
	cp.compliant, _ = cp.GetComplaint(complaintId)
	cp.employee, _ = cp.GetToUser()
	if err == nil {
		cp.addJobProgress()
		cp.sendEmailNotifier()
	}

	return
}

func (cp ComplaintCreator) GetToUser() (*models.User, error) {
	user := &models.User{}
	err := models.GetDB().QueryTable(models.TABLE_NAME_USER).Filter("id", cp.compliant.Employee.Id).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (cp ComplaintCreator) GetComplaint(complaintId uint) (*models.JobComplaint, error) {
	complain := &models.JobComplaint{}
	err := models.GetDB().QueryTable(models.TABLE_NAME_JOBCOMPLAINT).Filter("id", complaintId).One(complain)
	if err != nil {
		return nil, err
	}
	return complain, nil
}

func (cp *ComplaintCreator) addJobProgress() (err error) {
	p := progress.NewProgress(*cp)
	err = p.AddOne()
	return
}

//############################################################
// implements JobProgresser
//############################################################

func (cp ComplaintCreator) GetJobId() uint {
	return cp.jobId
}

func (cp ComplaintCreator) GetTableName() string {
	j := models.JobComplaint{}
	return j.TableName()
}

func (cp ComplaintCreator) GetType() progress.ProgressType {
	return progress.PT_Complain
}

func (cp ComplaintCreator) GetDesc() string {
	return "创建投诉"
}

func (cp ComplaintCreator) GetPrimaryKey() uint {
	return cp.id
}

//############################################################
// implements JobProgresser
//############################################################

func (cp ComplaintCreator) sendEmailNotifier() (err error) {
	n := notify.NewEmail(cp)
	err = n.Send()
	return
}

//############################################################
// implements EmailNotifier
//############################################################

func (cp ComplaintCreator) GetTo() []string {
	return []string{
		cp.employee.Email,
	}
}

func (cp ComplaintCreator) GetTemplateName() notify.EmailTemplate {
	return notify.PM_Notify
}

func (cp ComplaintCreator) GetVariables() []map[string]interface{} {
	complain := cp.compliant
	data := []map[string]interface{}{}
	data = append(data, map[string]interface{}{
		"name":            cp.employee.Name,
		"job_name":        cp.GetDesc(),
		"job_link":        fmt.Sprintf("%s/job/view/%d", utils.BaseUrl(), cp.jobId),
		"job_status_desc": complain.Complain,
		"job_created_at":  complain.Created.Format(utils.YMDHIS),
		"datetime":        time.Now().Format(utils.YMDHIS),
	})
	return data
}

//############################################################
// implements EmailNotifier
//############################################################
