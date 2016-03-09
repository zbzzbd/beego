package jobassign

import (
	"fmt"
	"models"
	"services/notify"
	"services/progress"
	"time"
	"utils"
)

func init() {
	//注册事件
	registerJobProgress()
}

func registerJobProgress() {
	//let it crash
	progress.Register(progress.PT_Assign)
}

type JobAssignment struct {
	jobId      uint
	fromUserId uint
	toUserId   uint
	remark     string
	jobAssign  *models.JobAssign
	job        *models.Job
	fromUser   *models.User
	toUser     *models.User
}

func NewJobAssignment(jobId int, fromUserId int, toUserId int, remark string) *JobAssignment {
	return &JobAssignment{
		jobId:      uint(jobId),
		fromUserId: uint(fromUserId),
		toUserId:   uint(toUserId),
		remark:     remark,
		jobAssign:  &models.JobAssign{},
	}
}

func (ja *JobAssignment) getJob() (*models.Job, error) {
	if ja.job != nil {
		return ja.job, nil
	}

	ja.job = &models.Job{}

	err := models.GetDB().QueryTable(models.TABLE_NAME_JOB).Filter("id", ja.jobId).RelatedSel("CreateUser").One(ja.job)
	if err != nil {
		return nil, err
	}

	return ja.job, nil
}

func (ja *JobAssignment) getFromUser() (*models.User, error) {
	if ja.fromUser != nil {
		return ja.fromUser, nil
	}
	user, err := ja.getUser(ja.fromUserId)
	if err != nil {
		return nil, err
	}
	ja.fromUser = user
	return user, nil
}

func (ja *JobAssignment) getToUser() (*models.User, error) {
	if ja.toUser != nil {
		return ja.toUser, nil
	}
	user, err := ja.getUser(ja.toUserId)
	if err != nil {
		return nil, err
	}
	ja.toUser = user
	return user, err
}

func (ja *JobAssignment) getUser(userId uint) (*models.User, error) {
	user := &models.User{}
	err := models.GetDB().QueryTable(models.TABLE_NAME_USER).Filter("id", userId).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ja *JobAssignment) getRemark() (string, error) {
	return ja.remark, nil
}

func (ja *JobAssignment) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("job.JobAssignment.Do : error = %s, ojb= %s", err.Error(), utils.Sdump(ja))
		}
		utils.GetLog().Debug("job.JobAssignment.Do : ojb= %s", utils.Sdump(ja))
	}()

	if ja.jobAssign.Job, err = ja.getJob(); err != nil {
		return
	}

	if ja.jobAssign.FromUser, err = ja.getFromUser(); err != nil {
		return
	}

	if ja.jobAssign.ToUser, err = ja.getToUser(); err != nil {
		return
	}

	if ja.jobAssign.Remark, err = ja.getRemark(); err != nil {
		return
	}

	if err = ja.jobAssign.Insert(); err != nil {
		return
	}

	if err = ja.afterAssignment(); err != nil {
		return
	}

	return
}

func (ja *JobAssignment) afterAssignment() (err error) {
	//更新job状态
	ja.jobAssign.Job.Employee, _ = ja.getToUser()

	if err = ja.jobAssign.Job.Update(models.JOB_EMPLOYEE); err != nil {
		return
	}

	if err = ja.addJobProgress(); err != nil {
		return
	}

	if err = ja.sendEmailNotifier(); err != nil {
		return
	}

	return
}

func (ja *JobAssignment) addJobProgress() (err error) {
	p := progress.NewProgress(*ja)
	err = p.AddOne()
	return
}

//############################################################
// implements JobProgresser
//############################################################

func (ja JobAssignment) GetJobId() uint {
	return ja.jobId
}

func (ja JobAssignment) GetTableName() string {
	return ja.jobAssign.TableName()
}

func (ja JobAssignment) GetType() progress.ProgressType {
	return progress.PT_Assign
}

func (ja JobAssignment) GetDesc() string {
	return "作业转交"
}

func (ja JobAssignment) GetPrimaryKey() uint {
	return ja.jobAssign.Id
}

//############################################################
// implements JobProgresser
//############################################################

func (ja JobAssignment) sendEmailNotifier() (err error) {
	n := notify.NewEmail(ja)
	err = n.Send()
	return
}

//############################################################
// implements EmailNotifier
//############################################################
func (ja JobAssignment) GetTo() []string {
	user, _ := (&ja).getToUser()

	job, _ := (&ja).getJob()

	return []string{
		user.Email,
		job.CreateUser.Email,
	}
}

func (ja JobAssignment) GetTemplateName() notify.EmailTemplate {
	return notify.PM_Notify
}

func (ja JobAssignment) GetVariables() (data []map[string]interface{}) {
	toUser, _ := (&ja).getToUser()
	job, _ := (&ja).getJob()

	data = append(data, map[string]interface{}{
		"name":            toUser.Name,
		"job_name":        job.Type,
		"job_link":        fmt.Sprintf("%s/job/view/%d", utils.BaseUrl(), ja.job.Id),
		"job_status_desc": ja.GetDesc(),
		"job_created_at":  ja.jobAssign.Created.Format(utils.YMDHIS),
		"datetime":        time.Now().Format(utils.YMDHIS),
	})

	data = append(data, map[string]interface{}{
		"name":            job.CreateUser.Name,
		"job_name":        job.Type,
		"job_link":        fmt.Sprintf("%s/job/view/%d", utils.BaseUrl(), ja.job.Id),
		"job_status_desc": ja.GetDesc(),
		"job_created_at":  ja.jobAssign.Created.Format(utils.YMDHIS),
		"datetime":        time.Now().Format(utils.YMDHIS),
	})

	return data
}

//############################################################
// implements EmailNotifier
//############################################################
