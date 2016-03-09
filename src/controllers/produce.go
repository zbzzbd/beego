package controllers

import (
	"models"
	"services/job"
	"services/project"
	"services/user"
	"strconv"
	"utils"
)

type ProduceController struct {
	JobController
}

func (c *ProduceController) Prepare() {
	c.BaseController.Prepare()
}

func (c *ProduceController) Finish() {
	defer c.BaseController.Finish()
}

func (c *ProduceController) Get() {
	c.GetInitData()

	c.TplNames = "job/index.tpl"

	p, _ := c.GetInt("p")
	jobList := job.NewJobList(p, 10).SetCode(c.GetString("code")).SetProjectName(c.GetString("project")).
		SetDepartment(c.GetString("department")).SetType(c.GetString("type")).SetEmployeeId(c.GetString("employee_id")).
		SetFinishTimeFrom(c.GetString("finish_time_from")).SetFinishTimeTo(c.GetString("finish_time_to"))

	err := jobList.GetList()
	if err != nil {
		c.Data["error"] = err.Error()
	}

	c.Data["jobs"] = jobList.GetJobs()
	c.Data["count"] = jobList.GetCount()

	c.Data["paginator"] = utils.NewPaginator(c.Ctx.Request, int(10), jobList.GetCount())

	c.Data["code"] = c.GetString("code")
	c.Data["project__name"] = c.GetString("project")
	c.Data["department"] = c.GetString("department")
	c.Data["type"] = c.GetString("type")
	c.Data["employee_id"] = c.GetString("employee_id")
	c.Data["finish_time__gte"] = c.GetString("finish_time_from")
	c.Data["finish_time__lte"] = c.GetString("finish_time_to")
}

func (c *ProduceController) ClaimList() {
	c.GetInitData()

	c.TplNames = "produce/claim.tpl"
	c.Data["ClearUrl"] = "/produce/job/claim"

	filter := make(map[string]interface{})

	filter["code"] = c.GetString("code")
	filter["project__name"] = c.GetString("project")
	filter["department"] = c.GetString("department")
	filter["type"] = c.GetString("type")
	employee_id := c.GetString("employee_id")
	if employee_id == "" {
		filter["employee_id"] = strconv.Itoa(int(c.currentUser.GetUser().Id))

	} else {
		filter["employee_id"] = employee_id
	}
	filter["finish_time__gte"] = c.GetString("finish_time_from")
	filter["finish_time__lte"] = c.GetString("finish_time_to")
	filter["valid_status"] = 1
	filter["claim_status"] = 0
	filter["del_status"] = models.Job_DelStatus_effective
	for k, v := range filter {
		c.Data[k] = v
	}

	p, _ := c.GetInt("p")
	jobList := job.NewJobList(p, 10).SetConditions(filter)

	err := jobList.GetList()
	if err != nil {
		c.Data["error"] = err.Error()
	}

	c.Data["jobs"] = jobList.GetJobs()
	c.Data["count"] = jobList.GetCount()

	c.Data["paginator"] = utils.NewPaginator(c.Ctx.Request, int(10), jobList.GetCount())
}

func (c *ProduceController) SubmitList() {
	c.GetInitData()

	c.TplNames = "produce/submit.tpl"
	c.Data["ClearUrl"] = "/produce/job/submit"

	filter := make(map[string]interface{})

	filter["code"] = c.GetString("code")
	filter["project__name"] = c.GetString("project")
	filter["department"] = c.GetString("department")
	filter["type"] = c.GetString("type")
	employee_id := c.GetString("employee_id")
	if employee_id == "" {
		filter["employee_id"] = strconv.Itoa(int(c.currentUser.GetUser().Id))

	} else {
		filter["employee_id"] = employee_id
	}
	filter["finish_time__gte"] = c.GetString("finish_time_from")
	filter["finish_time__lte"] = c.GetString("finish_time_to")
	filter["valid_status"] = 1
	filter["claim_status"] = 1
	filter["submit_status__in"] = []int{0, 2}
	filter["del_status"] = models.Job_DelStatus_effective
	for k, v := range filter {
		c.Data[k] = v
	}

	p, _ := c.GetInt("p")
	jobList := job.NewJobList(p, 10).SetConditions(filter)

	err := jobList.GetList()
	if err != nil {
		c.Data["error"] = err.Error()
	}

	c.Data["jobs"] = jobList.GetJobs()
	c.Data["count"] = jobList.GetCount()

	c.Data["paginator"] = utils.NewPaginator(c.Ctx.Request, int(10), jobList.GetCount())
}

func (c *ProduceController) GetInitData() {
	c.Data["Types"] = c.getRequires()
	c.Data["Departments"] = departments

	filter := make(map[string]interface{})
	projectList := project.NewProjectList()
	projects, err := projectList.GetList(filter)
	if err == nil {
		c.Data["Projects"] = projects
	}

	users, err := user.GetList()
	if err == nil {
		c.Data["Employees"] = users
	}
}
