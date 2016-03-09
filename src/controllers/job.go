package controllers

import (
	"models"
	"services/export"
	"services/job"
	"services/progress"
	"services/project"
	"services/require"
	"services/upload"
	"services/user"
	"strconv"
	"utils"
)

const (
	FileSaveKey = "job"
)

type JobController struct {
	BaseController
}

func (c *JobController) Prepare() {
	c.BaseController.Prepare()
}

func (c *JobController) Finish() {
	defer c.BaseController.Finish()
}

func (c *JobController) RecoverJob() {
	data := make(map[string]interface{})
	data["Job_id"] = c.Ctx.Input.Param(":id")
	_, err := job.NewJobList(1, -1).UpdateJobDelStatus(data, models.Job_DelStatus_effective)

	if err != nil {
		c.SetJsonResponse("error", err.Error())
	}
	c.GetJsonResponse().ServeJson()
}
func (c *JobController) DelJob() {
	data := make(map[string]interface{})
	data["Job_id"] = c.Ctx.Input.Param(":id")
	_, err := job.NewJobList(1, -1).UpdateJobDelStatus(data, models.Job_DelStatus_invalid)

	if err != nil {
		c.SetJsonResponse("error", err.Error())
	}
	c.GetJsonResponse().ServeJson()
}

func (c *JobController) Export() {
	c.getInitData()
	p, _ := c.GetInt("p")
	del_status, _ := c.GetInt("delstatus")
	jobList := job.NewJobList(p, -1).SetCode(c.GetString("code")).SetProjectName(c.GetString("project")).
		SetDepartment(c.GetString("department")).SetType(c.GetString("type")).SetEmployeeId(c.GetString("employee_id")).
		SetFinishTimeFrom(c.GetString("finish_time_from")).SetFinishTimeTo(c.GetString("finish_time_to")).
		SetCondition("submit_time__gte", c.GetString("submit_time_from")).SetCondition("submit_time__lte", c.GetString("submit_time_to")).
		SetCondition("create_user_id", c.GetString("create_user_id")).SetCondition("del_status", del_status)

	err := jobList.GetList()
	if err != nil {
		c.Data["error"] = err.Error()
	}
	records := jobList.GetExportData()
	c.Ctx.Output.Download(export.NewExporteFile(jobList).MakeCsvFile(records).GetFilePath(), jobList.GetOutFileName())
}

func (c *JobController) Get() {
	createUserId, _ := c.GetInt("create_user_id", c.userId)
	conditions := map[string]interface{}{
		"create_user_id": createUserId,
		"del_status":     models.Job_DelStatus_effective,
	}
	jl, _ := c.findJobList(conditions)
	c.makePaginator(jl)
	c.setFindConditions()

	c.TplNames = "job/index.tpl"
}

func (c *JobController) GetDelJobList() {
	c.TplNames = "project/del_list.tpl"
	conditions := map[string]interface{}{}
	conditions["del_status"] = models.Job_DelStatus_invalid
	jl, _ := c.findJobList(conditions)
	c.makePaginator(jl)
	c.setFindConditions()
}
func (c *JobController) GetJobList() {
	c.TplNames = "job_valid/job_list.tpl"
	conditions := map[string]interface{}{}
	conditions["del_status"] = models.Job_DelStatus_effective
	jl, _ := c.findJobList(conditions)
	c.makePaginator(jl)
	c.setFindConditions()
}

func (c *JobController) ViewCreateJob() {
	c.getInitData()

	c.Data["post_url"] = "/job/create"
	c.TplNames = "job/create.tpl"
}

func (c *JobController) PostCreateJob() {
	id, _ := c.GetInt("project_id")
	employee_id, _ := c.GetInt("employee_id")
	jid, historyId, err := job.NewJobCreator(c.currentUser.GetUser()).SetCondition("ProjectId", id).
		SetCondition("EmployeeId", employee_id).SetCondition("Type", c.GetString("type")).
		SetCondition("Department", c.GetString("department")).SetCondition("Target", c.GetString("target")).
		SetCondition("TargetUrl", c.GetString("target_url")).SetCondition("Desc", c.GetString("desc")).
		SetCondition("Message", c.GetString("message")).SetCondition("FinishTime", c.GetString("finish_time")).Do()
	if err != nil {
		c.SetJsonResponse("error", err.Error())
	} else {
		c.SetJsonResponse("id", uint(jid))
		c.uploadFiles(historyId, upload.FT_Create)
	}
	c.GetJsonResponse().ServeJson()
}

func (c *JobController) uploadFiles(relId uint, uft upload.UploadFileType) {
	upload.NewUploadFile(c.Ctx.Request, "files[]", FileSaveKey).SetFileType(uft).SetRelId(relId).Do()
}

func (c *JobController) ViewEditJob() {
	c.ViewCreateJob()
	c.Data["Job"], _ = job.NewJobList(1, 1).SetCondition("id", c.Ctx.Input.Param(":id")).GetOne()
	//c.Data["JobFiles"], _ = upload.NewQueryFile().SetCondition("rel_id", c.Ctx.Input.Param(":id")).GetJobAllFiles()

	c.Data["post_url"] = "/job/edit/" + c.Ctx.Input.Param(":id")

	c.Data["is_edit"] = true
}

func (c *JobController) PostEditJob() {
	id := c.Ctx.Input.Param(":id")
	historyId, err := job.NewJobModify(c.currentUser.GetUser(), job.JobCode_BussinessModify).
		SetCondition("Id", id).SetCondition("ProjectId", c.GetString("project_id")).SetCondition("EmployeeId", c.GetString("employee_id")).
		SetCondition("Type", c.GetString("type")).SetCondition("Department", c.GetString("department")).SetCondition("Target", c.GetString("target")).
		SetCondition("TargetUrl", c.GetString("target_url")).SetCondition("Desc", c.GetString("desc")).SetCondition("Message", c.GetString("message")).
		SetCondition("FinishTime", c.GetString("finish_time")).Do()

	if err != nil {
		c.SetJsonResponse("error", err.Error())
	} else {
		jid, _ := strconv.Atoi(id)
		c.SetJsonResponse("id", uint(jid))
		c.uploadFiles(historyId, upload.FT_Modify)
	}
	c.GetJsonResponse().ServeJson()
}

func (c *JobController) ViewJob() {
	id := c.Ctx.Input.Param(":id")
	c.Data["error"] = ""
	c.TplNames = "job/view.tpl"

	c.getJobCards(id)

	filter := make(map[string]interface{})
	filter["id"] = id
	c.Data["JobFiles"], _ = upload.NewQueryFile().SetCondition("rel_id", id).GetJobAllFiles()
}

func (c *JobController) DelJobFile() {
	postData := make(map[string]interface{})
	postData["id"] = c.GetString("id")
	err := upload.NewQueryFile().SetCondition("id", c.GetString("id")).DelFile()

	c.SetJsonResponse("error", err.Error())
	c.GetJsonResponse().ServeJson()
}

func (c *JobController) getJobCards(jobId string) {
	c.Data["jobId"] = jobId
	c.Data["cards"] = ""
	progressList := progress.NewProgressList(jobId)
	err := progressList.Do()
	if err != nil {
		return
	}
	c.Data["cards"] = progressList.RenderedTemplates()
}

func (c *JobController) findJobList(params map[string]interface{}) (jl *job.JobList, err error) {

	p, _ := c.GetInt("p")
	status := c.GetString("status")
	jl = job.NewJobList(p, 10).SetCode(c.GetString("code")).SetProjectName(c.GetString("project")).
		SetDepartment(c.GetString("department")).SetType(c.GetString("type")).SetEmployeeId(c.GetString("employee_id")).
		SetFinishTimeFrom(c.GetString("finish_time_from")).SetFinishTimeTo(c.GetString("finish_time_to")).
		SetCondition("submit_time__gte", c.GetString("submit_time_from")).SetCondition("submit_time__lte", c.GetString("submit_time_to")).
		SetCondition("create_user_id", c.GetString("create_user_id"))

	if len(params) > 0 {
		jl.SetConditions(params)
	}

	if status != "" {
		jl.GetStatusValueByDesc(status)
	}

	err = jl.GetList()
	if err != nil {
		c.Data["error"] = err.Error()
	}
	return
}

func (c *JobController) setFindConditions() {
	c.getInitData()

	c.Data["code"] = c.GetString("code")
	c.Data["create_user_id"] = c.GetString("create_user_id")
	c.Data["project__name"] = c.GetString("project")
	c.Data["department"] = c.GetString("department")
	c.Data["type"] = c.GetString("type")
	c.Data["employee_id"] = c.GetString("employee_id")
	c.Data["finish_time__gte"] = c.GetString("finish_time_from")
	c.Data["finish_time__lte"] = c.GetString("finish_time_to")
	c.Data["submit_time__gte"] = c.GetString("submit_time_from")
	c.Data["submit_time__lte"] = c.GetString("submit_time_to")
	c.Data["status"] = c.GetString("status")
}

func (c *JobController) makePaginator(jl *job.JobList) {
	c.Data["jobs"] = jl.GetJobs()
	c.Data["count"] = jl.GetCount()
	c.Data["paginator"] = utils.NewPaginator(c.Ctx.Request, int(10), jl.GetCount())
	c.Data["role"] = c.currentUser.GetRole()
}

func (c *JobController) getInitData() {
	c.Data["Types"] = c.getRequires()
	c.Data["Departments"] = departments
	c.Data["is_edit"] = false

	c.Data["ProjectNames"], _ = project.NewProjectList().GetAllProjectNames()
	c.Data["Status"] = c.getinitStatusData()

	users, err := user.GetList()
	if err == nil {
		c.Data["Employees"] = users
	}
	c.Data["BussinessMen"] = user.BussinessMen
}

func (c *JobController) getRequires() []string {
	return require.NewRequireList().GetRequireNames()
}

func (c *JobController) getDepartments() {

}

func (c *JobController) getinitStatusData() []string {
	status := []string{
		"待审核",
		"待认领制作",
		"审核未通过待修改",
		"任务取消",

		"制作中",
		"认领被拒绝",

		"待验收",
		"验收通过",
		"验收未通过",
	}
	return status
}

var departments = []struct {
	Department string
	Role       user.RoleType
}{
	{Department: user.TechGuy.Desc(), Role: user.TechGuy},
	{Department: user.ArtGuy.Desc(), Role: user.ArtGuy},
	{Department: user.Customer.Desc(), Role: user.Customer},
	{Department: user.Interactive.Desc(), Role: user.Interactive},
}
