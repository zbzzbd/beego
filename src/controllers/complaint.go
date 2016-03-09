package controllers

import (
	"services/export"
	"services/job_complaint"
	"services/project"
	"services/upload"
	"services/user"
	"strconv"
	"utils"
)

type ComplaintController struct {
	JobController
}

func (c *ComplaintController) Prepare() {
	c.BaseController.Prepare()
}

func (c *ComplaintController) Finish() {
	defer c.BaseController.Finish()
}

func (c *ComplaintController) Show() {
	c.getInitData()
	c.TplNames = "complaint/show.tpl"

	id := c.Ctx.Input.Param(":id")
	c.getJobCards(id)
}

func (c *ComplaintController) Del() {
	data := make(map[string]interface{})
	data["Complain_id"] = c.Ctx.Input.Param(":id")

	_, err := complaint.NewComplaintList(1, -1).UpdateComplaintDelStatus(data)
	if err != nil {
		c.SetJsonResponse("error", err.Error())

	}
	c.GetJsonResponse().ServeJson()
}

func (c *ComplaintController) Get() {
	c.JobController.GetJobList()
	c.TplNames = "complaint/complaint_index.tpl"
}

func (c *ComplaintController) Export() {
	c.GetInitData()
	p, _ := c.GetInt("p")
	complaintList := complaint.NewComplaintList(p, -1).SetCode(c.GetString("code")).SetProjectName(c.GetString("project")).SetDepartment(c.GetString("department")).
		SetEmployeeId(c.GetString("employee_id")).SetFinishTimeFrom(c.GetString("end_time_from")).SetFinishTimeTo(c.GetString("end_time_to")).SetResponse(c.GetString("response")).
		SetConditions("del_status", 0)
	err := complaintList.GetComplainList()

	if err != nil {
		c.Data["error"] = err.Error()
	}
	records := complaintList.GetExportData()
	c.Ctx.Output.Download(export.NewExporteFile(complaintList).MakeCsvFile(records).GetFilePath(), complaintList.GetOutFileName())

}

func (c *ComplaintController) GetInitData() {
	response := []string{"1", "0"}
	c.Data["Types"] = c.getRequires()
	c.Data["Departments"] = departments
	c.Data["Response"] = response

	//获取project对象
	filter := make(map[string]interface{})
	projectList := project.NewProjectList()
	projects, err := projectList.GetList(filter)
	if err == nil {
		c.Data["Projects"] = projects
	}

	//获取作业单元
	users, err := user.GetList()
	if err == nil {
		c.Data["Employees"] = users

	}

}

//查看我的投诉列表
func (c *ComplaintController) ViewMyCompliant() {
	c.GetInitData()
	c.TplNames = "complaint/my_complaint_list.tpl"

	p, _ := c.GetInt("p")
	complaintList := complaint.NewComplaintList(p, 10).SetCode(c.GetString("code")).SetProjectName(c.GetString("project")).SetDepartment(c.GetString("department")).
		SetEmployeeId(c.GetString("employee_id")).SetFinishTimeFrom(c.GetString("end_time_from")).SetFinishTimeTo(c.GetString("end_time_to")).SetResponse(c.GetString("response")).
		SetConditions("del_status", 0)
	err := complaintList.GetComplainList()
	if err != nil {
		c.Data["error"] = err.Error()
	}
	c.Data["complains"] = complaintList.GetComplaints()
	c.Data["count"] = complaintList.GetCount()
	c.Data["paginator"] = utils.NewPaginator(c.Ctx.Request, int(10), complaintList.GetCount())

	c.Data["job__code"] = c.GetString("code")
	c.Data["project__name"] = c.GetString("project")
	c.Data["job__department"] = c.GetString("department")
	c.Data["response"] = c.GetString("response")
	employee_id := c.GetString("employee_id")

	if employee_id == "" {
		c.Data["employee_id"] = strconv.Itoa(int(c.currentUser.GetUser().Id))
	} else {
		c.Data["employee_id"] = employee_id
	}
	c.Data["end_time__gte"] = c.GetString("end_time_from")
	c.Data["end_time__lte"] = c.GetString("end_time_to")

}

//查询投诉列表
func (c *ComplaintController) ViewCompliant() {
	c.GetInitData()
	c.TplNames = "complaint/complain_list.tpl"
	p, _ := c.GetInt("p")
	complaintList := complaint.NewComplaintList(p, 10).SetCode(c.GetString("code")).SetProjectName(c.GetString("project")).SetDepartment(c.GetString("department")).
		SetEmployeeId(c.GetString("employee_id")).SetFinishTimeFrom(c.GetString("end_time_from")).SetFinishTimeTo(c.GetString("end_time_to")).SetResponse(c.GetString("response")).
		SetConditions("del_status", 0)
	err := complaintList.GetComplainList()
	if err != nil {
		c.Data["error"] = err.Error()
	}
	c.Data["complains"] = complaintList.GetComplaints()
	c.Data["count"] = complaintList.GetCount()
	c.Data["paginator"] = utils.NewPaginator(c.Ctx.Request, int(10), complaintList.GetCount())

	c.Data["job__code"] = c.GetString("code")
	c.Data["project__name"] = c.GetString("project")
	c.Data["job__department"] = c.GetString("department")
	c.Data["response"] = c.GetString("response")
	employee_id := c.GetString("employee_id")

	if employee_id == "" {
		c.Data["employee_id"] = strconv.Itoa(int(c.currentUser.GetUser().Id))

	} else {
		c.Data["employee_id"] = employee_id
	}

	c.Data["end_time__gte"] = c.GetString("end_time_from")
	c.Data["end_time__lte"] = c.GetString("end_time_to")
}

//创建一个投诉
func (c *ComplaintController) CreateComplain() {
	c.getInitData()
	postData := make(map[string]interface{})
	postData["JobId"] = c.GetString("job_id")
	postData["Type"] = c.GetString("Type")
	postData["Complain"] = c.GetString("Desc")
	postData["response"], _ = c.GetInt("reply")
	postData["editstatus"] = c.GetString("editstatus")
	postData["employee_id"] = c.GetString("employee_id")
	postData["source"] = c.GetString("source")
	complianId, jobId, err := complaint.NewcomplainCreator(postData, c.currentUser.GetUser()).Do()

	if err != nil {
		c.SetJsonResponse("error", err.Error())
	} else {
		c.SetJsonResponse("id", jobId)
		c.SetJsonResponse("complainid", complianId)
		//c.uploadFiles(job.FT_Complain,jobId)
	}

	c.GetJsonResponse().ServeJson()
}

//根据编号进行查询一个job
func (c *ComplaintController) GetJob() {
	defer func() {
		c.GetJsonResponse().ServeJson()
	}()

	id := c.Input().Get("id")
	if id == "" {
		return
	}
	filter := make(map[string]interface{})
	filter["code"] = id

	jobModel, err := complaint.GetOneJob(filter)
	c.SetJsonResponse("Job", jobModel)
	if err != nil {
		c.SetJsonResponse("error", err.Error())
		return
	}

	jobFiles, err := upload.NewQueryFile().SetCondition("rel_id", id).GetJobAllFiles()
	c.SetJsonResponse("JobFiles", jobFiles)
}
