package job

import (
	"errors"
	"fmt"
	"helpers"
	"models"
	"services/export"
	"strconv"
	"utils"

	"github.com/astaxie/beego/orm"
)

type JobList struct {
	jobs     []*models.Job
	page     int
	pageSize int
	count    int64
	filter   map[string]interface{}
}

func NewJobList(page, pageSize int) *JobList {
	jobList := JobList{}
	jobList.filter = make(map[string]interface{})
	jobList.page = page
	if jobList.page < 1 {
		jobList.page = 1
	}
	if pageSize > 0 {
		jobList.pageSize = pageSize
		jobList.filter["limit"] = jobList.pageSize
		jobList.filter["offset"] = (jobList.page - 1) * jobList.pageSize
	}
	return &jobList
}

func (jobList *JobList) GetJobs() []*models.Job {
	return jobList.jobs
}

func (jobList *JobList) GetCount() int64 {
	return jobList.count
}

func (jobList *JobList) SetCode(code string) *JobList {
	jobList.filter["code__icontains"] = code
	return jobList
}

func (jobList *JobList) SetProjectName(project string) *JobList {
	jobList.filter["project__name__icontains"] = project
	return jobList
}

func (jobList *JobList) SetDepartment(department string) *JobList {
	jobList.filter["department"] = department
	return jobList
}

func (jobList *JobList) SetType(t string) *JobList {
	jobList.filter["type"] = t
	return jobList
}

func (jobList *JobList) SetEmployeeId(employee_id string) *JobList {
	jobList.filter["employee_id"] = employee_id
	return jobList
}

func (jobList *JobList) SetFinishTimeFrom(finish_time string) *JobList {
	jobList.filter["finish_time__gte"] = finish_time
	return jobList
}

func (jobList *JobList) SetFinishTimeTo(finish_time string) *JobList {
	jobList.filter["finish_time__lte"] = finish_time
	return jobList
}
func (jobList *JobList) SetValidStatus(valid_status string) *JobList {
	jobList.filter["valid_status"] = valid_status
	return jobList
}
func (jobList *JobList) SetClaimStatus(claim_status string) *JobList {
	jobList.filter["claim_status"] = claim_status
	return jobList
}
func (jobList *JobList) SetSubmitStatus(submit_status string) *JobList {
	jobList.filter["submit_status"] = submit_status
	return jobList
}

func (jobList *JobList) SetConditions(condition map[string]interface{}) *JobList {
	for k, v := range condition {
		jobList.filter[k] = v
	}

	return jobList
}

func (jobList *JobList) SetCondition(k string, v interface{}) *JobList {
	jobList.filter[k] = v

	return jobList
}

//更新删除状态
func (jobList *JobList) UpdateJobDelStatus(filter map[string]interface{}, Delstatus uint) (num int64, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.job.updateJobDelStatus: error : %s", err.Error())
		} else {
			utils.GetLog().Debug("services.job.update: debug: filter=%s,job=%v", utils.Sdump(filter))
		}
	}()

	q := models.GetDB().QueryTable(models.TABLE_NAME_JOB)
	for k, v := range filter {
		if k == "Job_id" {
			num, err = q.Filter("Id", v).Update(orm.Params{"del_status": Delstatus})
		}
	}
	return
}
func (jobList *JobList) GetStatusValueByDesc(status string) *JobList {
	statusConditions := map[string]func(){
		"待审核": func() {
			jobList = jobList.SetCondition(models.JOB_VALID_STATUS, models.Job_Valid_None)
		},
		"待认领制作": func() {
			jobList = jobList.SetCondition(models.JOB_VALID_STATUS, models.Job_Valid_OK).SetCondition(models.JOB_CLAIM_STATUS, models.Job_Claim_None)
		},
		"审核未通过待修改": func() {
			jobList = jobList.SetCondition(models.JOB_VALID_STATUS, models.Job_Valid_Refuse)
		},
		"任务取消": func() {
			jobList = jobList.SetCondition(models.JOB_VALID_STATUS, models.Job_Valid_Cancel)
		},

		"制作中": func() {
			jobList = jobList.SetCondition(models.JOB_VALID_STATUS, models.Job_Valid_OK).SetCondition(models.JOB_CLAIM_STATUS, models.Job_Claim_Ok).SetCondition(models.JOB_SUBMIT_STATUS, models.Job_Submit_None)
		},
		"认领被拒绝": func() {
			jobList = jobList.SetCondition(models.JOB_VALID_STATUS, models.Job_Valid_OK).SetCondition(models.JOB_CLAIM_STATUS, models.Job_Claim_Refuse)
		},

		"待验收": func() {
			jobList = jobList.SetCondition(models.JOB_VALID_STATUS, models.Job_Valid_OK).SetCondition(models.JOB_CLAIM_STATUS, models.Job_Claim_Ok).SetCondition(models.JOB_SUBMIT_STATUS, models.Job_Submit_Wait_Acceptance)
		},
		"验收通过": func() {
			jobList = jobList.SetCondition(models.JOB_VALID_STATUS, models.Job_Valid_OK).SetCondition(models.JOB_CLAIM_STATUS, models.Job_Claim_Ok).SetCondition(models.JOB_SUBMIT_STATUS, models.Job_Submit_Acceptance_OK)
		},
		"验收未通过": func() {
			jobList = jobList.SetCondition(models.JOB_VALID_STATUS, models.Job_Valid_OK).SetCondition(models.JOB_CLAIM_STATUS, models.Job_Claim_Ok).SetCondition(models.JOB_SUBMIT_STATUS, models.Job_Submit_Acceptance_Refuse)
		},
	}

	if f, ok := statusConditions[status]; ok {
		f()
	}
	return jobList
}

func (jobList *JobList) GetList() (err error) {
	q := models.GetDB().QueryTable(models.TABLE_NAME_JOB).OrderBy("-created").RelatedSel("Project", "Employee", "CreateUser")
	for k, v := range jobList.filter {
		if k != "limit" && k != "offset" && v != "" {
			q = q.Filter(k, v)
		}
	}
	jobList.count, _ = q.Count()

	if jobList.filter["limit"] != nil {
		q = q.Limit(jobList.filter["limit"])
	}

	if jobList.filter["offset"] != nil {
		q = q.Offset(jobList.filter["offset"])
	}
	_, err = q.All(&jobList.jobs)

	return
}

func (jobList *JobList) GetOne() (*models.Job, error) {
	var retJob models.Job
	var err error
	q := models.GetDB().QueryTable(models.TABLE_NAME_JOB).RelatedSel("Project", "Employee")
	for k, v := range jobList.filter {
		if k != "limit" && k != "offset" && v != "" {
			q = q.Filter(k, v)
		}
	}
	err = q.One(&retJob)
	if err == orm.ErrNoRows {
		err = errors.New("查找作业不存在")
	}

	return &retJob, err
}

func (joblist *JobList) GetUserName(employId uint) (*models.User, error) {
	user := &models.User{}
	err := models.GetDB().QueryTable(models.TABLE_NAME_USER).Filter("id", employId).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//###############################
// implements  exporter
//#############################

func (jobList *JobList) GetExportData() [][]string {
	var res []*models.Job
	records := [][]string{{"序号", "作业状态", "作业编号", "项目名称", "业务单元", "作业要求", "作业对象", "作业部门", "作业单元", "发布时间", "审核时间", "要求完成时间", "作业认领时间", "实际完成时间", "任务用时", "任务延时"}}
	res = jobList.GetJobs()
	for i := 0; i < len(res); i++ {
		data := []string{strconv.Itoa(i + 1), helpers.GetStatusDesc(*res[i]), res[i].Code, res[i].Project.Name, res[i].CreateUser.Name, res[i].Type, res[i].Target,
			res[i].Department, res[i].Employee.Name, helpers.TimeFormat(res[i].Created), helpers.TimeFormat(res[i].ValidTime), helpers.TimeFormat(res[i].FinishTime), helpers.TimeFormat(res[i].ClaimTime),
			helpers.TimeFormat(res[i].SubmitTime), helpers.GetTimeDiff(res[i].SubmitTime, res[i].ClaimTime), helpers.GetTimeDiff(res[i].SubmitTime, res[i].FinishTime),
		}
		records = append(records, data)
	}
	return records
}

func (jobList *JobList) GetOutFileName() string {
	fileName := "作业总单"
	mapData := jobList.filter
	style := jobList.GetOutStyleFormat()
	for k, v := range mapData {
		if k != "employee_id" && k != "create_user_id" && v != "" {
			fileName = fileName + "-" + fmt.Sprint(v)
		} else if v != "" {
			tmp, _ := strconv.Atoi(fmt.Sprint(v))
			employee, _ := jobList.GetUserName(uint(tmp))
			fileName = fileName + "-" + employee.Name
		}
	}
	return fileName + style

}

func (jobList *JobList) GetOutStyleFormat() string {
	return export.CSV_FORMAT
}
