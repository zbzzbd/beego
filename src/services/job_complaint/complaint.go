package complaint

import (
	"errors"
	"models"
	"services/export"
	"services/progress"
	"utils"

	"fmt"
	"helpers"
	"strconv"

	"github.com/astaxie/beego/orm"
)

func init() {
	registerJobProgress()
}

func registerJobProgress() {
	progress.Register(progress.PT_ComplainReplay)
	progress.Register(progress.PT_Complain)
}

//更新投诉状态
func UpdateComplaint(filter map[string]interface{}) (num int64, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.job.UpdateComplain : error : %s", err.Error())
		} else {
			utils.GetLog().Debug("services.job.UpdateComplain : debug : filter=%s, job=%v", utils.Sdump(filter))
		}
	}()

	q := models.GetDB().QueryTable("job_complaints")
	for k, v := range filter {
		if k == "Complain_id" {
			num, err = q.Filter("Id", v).Update(orm.Params{"ReplyStatus": 1})
		}
	}
	return
}

//根据投诉编号获取投诉信息
func GetOneComplain(filter map[string]interface{}) (*models.JobComplaint, error) {
	var complaint models.JobComplaint
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.job.GetOneComplain : error : %s", err.Error())
		} else {
			utils.GetLog().Debug("services.job.GetOneComplain : debug : filter=%s, job=%v", utils.Sdump(filter), utils.Sdump(complaint))
		}
	}()
	q := models.GetDB().QueryTable("job_complaints").RelatedSel("Job", "User", "Project")
	for k, v := range filter {
		q = q.Filter(k, v)
	}
	err = q.One(&complaint)
	if err == orm.ErrNoRows {
		err = errors.New("投诉找不到")
	}
	return &complaint, err

}

//获取投诉列表
func GetComplainList(filter map[string]interface{}) (compalins []*models.JobComplaint, cnt int64, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.job.GetList : error : %s", err.Error())
		} else {
			utils.GetLog().Debug("services.job.GetList : debug : filter=%s, compalints=%v", utils.Sdump(filter), utils.Sdump(compalins))
		}
	}()

	q := models.GetDB().QueryTable("job_complaints").OrderBy("-created").RelatedSel("Job", "User", "Project", "Employee")

	//按照传入的条件进行过滤查询结果
	for k, v := range filter {
		if k != "limit" && k != "offset" && v != "" {
			q = q.Filter(k, v)
		}
	}
	cnt, _ = q.Count()

	if filter["limit"] != nil {
		q = q.Limit(filter["limit"].(int))
	}

	if filter["offset"] != nil {
		q = q.Offset(filter["offset"].(int))
	}

	_, err = q.All(&compalins)

	return

}

//根据ID查询对应的job
func GetOneJob(filter map[string]interface{}) (*models.Job, error) {
	var retJob models.Job
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.job.GetOne : error : %s", err.Error())
		} else {
			utils.GetLog().Debug("services.job.GetOne : debug : filter=%s, job=%v", utils.Sdump(filter), utils.Sdump(retJob))
		}
	}()

	q := models.GetDB().QueryTable(models.TABLE_NAME_JOB).RelatedSel("Project", "Employee")
	for k, v := range filter {
		q = q.Filter(k, v)
	}
	err = q.One(&retJob)
	if err == orm.ErrNoRows {
		err = errors.New("查找作业不存在")
	}

	return &retJob, err
}

//创建一个投诉

func CreateComlian(mapComplain map[string]interface{}, user *models.User) (complainId uint, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.complain.Create : error : %s, ", err.Error())
		} else {
			utils.GetLog().Debug("services.complain.Create : debug : mapComplain=%s", utils.Sdump(mapComplain))
		}
	}()

	var complian models.JobComplaint
	pid, _ := mapComplain["Jobcode"].(string)

	filter := make(map[string]interface{})
	filter["code"] = pid

	var job *models.Job
	job, err = GetOneJob(filter)
	complian.Job = &models.Job{Id: job.Id}
	complian.Project = &models.Project{Id: job.Project.Id}
	complian.Employee = &models.User{Id: job.Employee.Id}

	complian.Complain = mapComplain["Complain"].(string)
	complian.Type = mapComplain["Type"].(string)
	response, _ := mapComplain["response"].(int)
	complian.Response = uint(response)

	complian.User = user

	var insertId int64
	insertId, err = models.GetDB().Insert(&complian)
	if err != nil {
		utils.GetLog().Debug("services.complian.Create : debug : job=%s", utils.Sdump(complian))
		return
	}
	complainId = uint(insertId)
	return

}

//更新删除状态
func (c *ComplainList) UpdateComplaintDelStatus(filter map[string]interface{}) (num int64, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.job.UpdateComplain : error : %s", err.Error())
		} else {
			utils.GetLog().Debug("services.job.UpdateComplain : debug : filter=%s, job=%v", utils.Sdump(filter))
		}
	}()

	q := models.GetDB().QueryTable("job_complaints")
	for k, v := range filter {
		if k == "Complain_id" {
			num, err = q.Filter("Id", v).Update(orm.Params{"del_status": 1})
		}
	}
	return
}

type ComplainList struct {
	complaints []*models.JobComplaint
	page       int
	pageSize   int
	count      int64
	filter     map[string]interface{}
}

//根据page,pagesize 返回ComplainList
func NewComplaintList(page, pageSize int) *ComplainList {
	complaintList := ComplainList{}
	complaintList.filter = make(map[string]interface{})
	complaintList.page = page
	complaintList.pageSize = pageSize

	if complaintList.page < 1 {
		complaintList.page = 1
	}

	if pageSize > 0 {
		complaintList.pageSize = pageSize
		complaintList.filter["limit"] = complaintList.pageSize
		complaintList.filter["offset"] = (complaintList.page - 1) * complaintList.pageSize
	}

	return &complaintList
}

//从列表中获取投诉
func (complainList *ComplainList) GetComplaints() []*models.JobComplaint {
	return complainList.complaints
}

func (complainList *ComplainList) GetCount() int64 {
	return complainList.count
}
func (complainList *ComplainList) SetCode(code string) *ComplainList {
	complainList.filter["job__code"] = code
	return complainList
}

func (complainList *ComplainList) SetProjectName(project string) *ComplainList {

	complainList.filter["project__name"] = project
	return complainList
}

func (complainList *ComplainList) SetDepartment(department string) *ComplainList {
	complainList.filter["job__department"] = department
	return complainList
}
func (complainList *ComplainList) SetResponse(response string) *ComplainList {

	complainList.filter["response"] = response
	return complainList
}

func (complainList *ComplainList) SetEmployeeId(employee_id string) *ComplainList {

	complainList.filter["employee_id"] = employee_id
	return complainList
}

func (complainList *ComplainList) SetFinishTimeFrom(finish_time string) *ComplainList {
	complainList.filter["created__gte"] = finish_time
	return complainList
}

func (complainList *ComplainList) SetFinishTimeTo(end_time string) *ComplainList {
	complainList.filter["created__lte"] = end_time
	return complainList
}

func (complainList *ComplainList) SetCondition(condition map[string]interface{}) *ComplainList {
	for k, v := range condition {
		complainList.filter[k] = v
	}
	return complainList
}

func (complainList *ComplainList) SetConditions(k string, v interface{}) *ComplainList {
	complainList.filter[k] = v

	return complainList
}
func (complainList *ComplainList) GetComplainList() (err error) {
	complainList.complaints, complainList.count, err = GetComplainList(complainList.filter)

	return
}

//########################
// implements  exporter
//########################

func (complainList *ComplainList) GetExportData() [][]string {
	var res []*models.JobComplaint
	records := [][]string{{"序号", "投诉日期", "投诉来源", "作业编号", "项目名称", "业务单元", "作业部门", "作业单元", "客诉事项", "客诉描述", "答复需求", "答复情况"}}
	res = complainList.GetComplaints()
	length := len(res)
	for i := 0; i < length; i++ {
		data := []string{strconv.Itoa(i + 1), helpers.TimeFormat(res[i].Created), res[i].Source, res[i].Job.Code, res[i].Project.Name, res[i].User.Name, res[i].Job.Department,
			res[i].Employee.Name, res[i].Type, res[i].Complain, strconv.Itoa(int(res[i].Response)), strconv.Itoa(int(res[i].ReplyStatus))}
		records = append(records, data)
	}
	return records

}

func (complainList *ComplainList) GetOutFileName() string {
	fileName := "投诉进程"
	style := complainList.GetOutStyleFormat()
	data := complainList.filter
	for k, v := range data {
		if k != "" && v != "" {
			fileName = fileName + "-" + fmt.Sprint(v)
		}
	}
	return fileName + style
}

func (complainList *ComplainList) GetOutStyleFormat() string {
	return export.CSV_FORMAT
}
