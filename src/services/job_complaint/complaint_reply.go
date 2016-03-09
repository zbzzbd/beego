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

//创建一个投诉的回复
func (cpr *ComplaintReplytor) CreateComplaintReply(mapComplainReply map[string]interface{}, user *models.User) (replyId, jobId uint, err error) {

	defer func() {
		if err != nil {
			utils.GetLog().Error("services.complainReply.Create : error : %s, ", err.Error())
		} else {
			utils.GetLog().Debug("services.complainReply.Create : debug : mapComplainReply=%s", utils.Sdump(mapComplainReply))
		}
	}()
	var complainReply models.JobComplaintReply
	complainReply.Message = mapComplainReply["Desc"].(string)

	var complainId uint
	switch v := mapComplainReply["Complain_id"].(type) {
	case int:
		complainId = uint(v)
	case string:
		tmp, _ := strconv.Atoi(v)
		complainId = uint(tmp)
	}
	filter := make(map[string]interface{})
	filter["Id"] = complainId
	complians, _ := GetOneComplain(filter)
	complainReply.Complaint = &models.JobComplaint{Id: complainId}
	complainReply.User = user

	cpr.complain = *complians
	var repId int64
	repId, err = models.GetDB().Insert(&complainReply)

	if err != nil {
		utils.GetLog().Debug("services.complianReply.Create : debug : job=%s", utils.Sdump(complainReply))
		return
	}
	replyId = uint(repId)
	jobId = uint(complians.Job.Id)
	return
}

type ComplaintReplytor struct {
	id                uint
	jobId             uint
	mapComplaintReply map[string]interface{}
	user              *models.User
	complainReply     *models.JobComplaintReply
	complain          models.JobComplaint
	touser            *models.User
}

func NewComplaintReplytor(mapComplaintReply map[string]interface{}, user *models.User) *ComplaintReplytor {
	return &ComplaintReplytor{
		mapComplaintReply: mapComplaintReply,
		user:              user,
	}
}

func (cpr *ComplaintReplytor) Do() (replyId, jobId uint, err error) {
	cpr.id, cpr.jobId, err = cpr.CreateComplaintReply(cpr.mapComplaintReply, cpr.user)
	replyId = cpr.id
	jobId = cpr.jobId
	cpr.touser, _ = cpr.GetToUser()
	cpr.complainReply, _ = cpr.GetComplaintReply(replyId)

	if err == nil {
		cpr.addJobProgress()
		cpr.sendEmailNotifier()
	}

	return
}

func (cpr ComplaintReplytor) GetToUser() (*models.User, error) {
	user := &models.User{}
	err := models.GetDB().QueryTable(models.TABLE_NAME_USER).Filter("id", cpr.complain.User.Id).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (cp ComplaintReplytor) GetComplaintReply(replyId uint) (*models.JobComplaintReply, error) {
	complainReply := &models.JobComplaintReply{}
	err := models.GetDB().QueryTable(models.TABLE_NAME_JOBCOMPLAINTREPLY).Filter("id", replyId).One(complainReply)
	if err != nil {
		return nil, err
	}
	return complainReply, nil
}

func (cpr *ComplaintReplytor) addJobProgress() (err error) {
	p := progress.NewProgress(*cpr)
	err = p.AddOne()
	return
}

//############################################################
// implements JobProgresser
//############################################################

func (cpr ComplaintReplytor) GetJobId() uint {
	return cpr.jobId
}

func (cpr ComplaintReplytor) GetTableName() string {
	j := models.JobComplaintReply{}
	return j.TableName()
}

func (cpr ComplaintReplytor) GetType() progress.ProgressType {
	return progress.PT_ComplainReplay
}

func (cpr ComplaintReplytor) GetDesc() string {
	return "回复投诉"
}

func (cpr ComplaintReplytor) GetPrimaryKey() uint {
	return cpr.id
}

//############################################################
// implements JobProgresser
//############################################################

func (cpr ComplaintReplytor) sendEmailNotifier() (err error) {
	n := notify.NewEmail(cpr)
	err = n.Send()
	return
}

//############################################################
// implements EmailNotifier
//############################################################
func (cpr ComplaintReplytor) GetTo() []string {
	return []string{
		cpr.touser.Email,
	}
}
func (cpr ComplaintReplytor) GetTemplateName() notify.EmailTemplate {
	return notify.PM_Notify
}
func (cpr ComplaintReplytor) GetVariables() []map[string]interface{} {
	complainReply := cpr.complainReply

	data := []map[string]interface{}{}
	data = append(data, map[string]interface{}{
		"name":            cpr.touser.Name,
		"job_name":        cpr.GetDesc(),
		"job_link":        fmt.Sprintf("%s/job/view/%d", utils.BaseUrl(), cpr.jobId),
		"job_status_desc": complainReply.Message,
		"job_created_at":  complainReply.Created.Format(utils.YMDHIS),
		"datetime":        time.Now().Format(utils.YMDHIS),
	})
	return data
}

//############################################################
// implements EmailNotifier
//############################################################
