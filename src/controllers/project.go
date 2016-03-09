package controllers

import (
	"fmt"

	"models"
	"services/project"
	"services/user"
	"strconv"
	"time"
	"utils"
)

type ProjectController struct {
	BaseController
}

func (c *ProjectController) Prepare() {
	c.BaseController.Prepare()
}

func (c *ProjectController) Finish() {
	defer c.BaseController.Finish()
}

func (c *ProjectController) Get() {
	c.Redirect("/project/list", 302)
}

func (c *ProjectController) DelProject() {
	data := make(map[string]interface{})

	id := c.Ctx.Input.Param(":id")
	data["project_id"] = id
	_, err := project.NewProject().UpdateProjectStatus(data)
	if err != nil {
		c.SetJsonResponse("error", err.Error())
	}
	c.GetJsonResponse().ServeJson()

}
func (c *ProjectController) Export() {
	//获取导出条件
	filter := c.GetExportFilter()

	//查询结果
	Projects, err := project.NewProjectList().GetList(filter)
	if err == nil {
		records := project.GetExportData(Projects)
		//导出结果
		c.Ctx.Output.Download(project.NewProjectExport().MakeFile(records).GetFilePath(), project.NewProject().GetOutName(filter))
	}

}
func (c *ProjectController) ViewCreateProject() {
	c.TplNames = "project/create.tpl"

	if c.Data["RoleError"] != nil && c.Data["RoleError"].(bool) {
		return
	}

	c.Data["CreateJobStatus"] = "active"

	users := user.NewUserList()
	c.Data["Priority"] = c.GetPriority()

	err, progress := models.GetProgressList()
	if err == nil {
		c.Data["Progress"] = progress
	}
	c.Data["BussinessUser"] = users.GetRoleList(user.BussinessMen) //[]map[string]string{{"Id": "1", "Name": "许航"}}

	c.Data["ArtUser"] = users.GetRoleList(user.ArtGuy) // []map[string]string{{"Id": "1", "Name": "许航"}}

	c.Data["TechUser"] = users.GetRoleList(user.TechGuy) //[]map[string]string{{"Id": "1", "Name": "许航"}}
	c.Data["Now"] = time.Now().Local()
	c.GetProjectNames()
}

func (c *ProjectController) ViewProjectDetail() {
	id := c.Ctx.Input.Param(":id")

	c.Data["CreateJobStatus"] = "active"
	c.TplNames = "project/detail.tpl"

	project := project.NewProject()
	project, _ = project.GetOne(id)

	c.Data["Project"] = project.GetModelProject()
}

func (c *ProjectController) ViewEditProject() {
	id := c.Ctx.Input.Param(":id")

	c.TplNames = "project/edit.tpl"

	project := project.NewProject()
	project, _ = project.GetOne(id)

	c.Data["Project"] = project.GetModelProject()

	users := user.NewUserList()
	c.Data["Priority"] = c.GetPriority()
	c.Data["return"] = 1
	err, progress := models.GetProgressList()
	if err == nil {
		c.Data["Progress"] = progress
	}

	c.Data["BussinessUser"] = users.GetRoleList(user.BussinessMen) //[]map[string]string{{"Id": "1", "Name": "许航"}}
	c.Data["ArtUser"] = users.GetRoleList(user.ArtGuy)             // []map[string]string{{"Id": "1", "Name": "许航"}}
	c.Data["TechUser"] = users.GetRoleList(user.TechGuy)           //[]map[string]string{{"Id": "1", "Name": "许航"}}
	c.Data["Id"] = id
}

func (c *ProjectController) EditProject() {
	id := c.Ctx.Input.Param(":id")
	var retJson = make(map[string]interface{})
	retJson["result"] = 0
	retJson["msg"] = ""

	var params = make(map[string]interface{})

	params["id"] = id
	params["name"] = c.GetString("name")
	params["scale"] = c.GetString("scale")
	params["started"] = c.GetString("started")
	params["priority"] = c.GetString("priority")
	params["clientName"] = c.GetString("client_name")
	params["contractNo"] = c.GetString("contract_no")
	params["serviceItem"] = c.GetString("service_item")
	params["bussinessUser"] = c.GetString("bussiness_user")
	params["progress"] = c.GetString("progress")

	params["regStartDate"] = c.GetString("reg_start_date")
	params["gameDate"] = c.GetString("game_date")

	params["regCloseDate"] = c.GetString("reg_close_date")
	params["artUser"] = c.GetString("art_user")
	params["techUser"] = c.GetString("tech_user")
	//登记人
	params["registrant"] = c.currentUser.GetUser().Id

	project, err := project.NewProject().GetOne(id)
	if err != nil {
		retJson["result"] = -102
		retJson["msg"] = "can't find the project"
		goto END
	}
	err = project.Update(params)
	if err != nil {
		retJson["result"] = -103
		retJson["msg"] = "update the project failed"
	}

END:
	c.Data["json"] = retJson
	c.ServeJson()
}

func (c *ProjectController) HandleFilter() map[string]interface{} {
	filter := make(map[string]interface{})
	filter["limit"] = 10
	filter["offset"] = 0
	id := c.GetString("project_name")
	if id != "" {
		filter["id"] = id
		p, _ := project.NewProject().GetOne(id)
		c.Data["project_default"] = p.GetModelProject()
	}
	bussiness_user := c.GetString("bussiness_user")
	if bussiness_user != "" {
		filter["bussiness_user_id"] = bussiness_user
		user_id, _ := strconv.Atoi(bussiness_user)
		user, _ := user.NewUserWithId(uint(user_id))
		c.Data["bussiness_user_default"] = user.GetUser()
	}

	progress_filter := c.GetString("progress")
	if progress_filter != "" {
		filter["progress"] = progress_filter
		progress, _ := project.NewProgress().GetProgressById(progress_filter)
		c.Data["progress_default"] = progress.GetModelProgress()
	}

	start_date := c.GetString("start_date")
	if start_date != "" {
		filter["start_date"] = start_date
		c.Data["start_date_default"] = start_date
	}

	end_date := c.GetString("end_date")
	if end_date != "" {
		filter["end_date"] = end_date
		c.Data["end_date_default"] = end_date
	}

	p, _ := c.GetInt("p")
	if p > 1 {
		filter["offset"] = (p - 1) * filter["limit"].(int)
	}
	return filter
}

func (c *ProjectController) ProjectList() {

	c.TplNames = "project/list.tpl"

	filter := c.HandleFilter()

	getAll := make(map[string]interface{})
	allProjects, err := project.NewProjectList().GetList(getAll)
	if err == nil {
		c.Data["AllProjects"] = allProjects
	}

	projects, err := project.NewProjectList().GetList(filter)
	if err == nil {
		c.Data["Projects"] = projects
	}

	users := user.NewUserList()
	c.Data["BussinessUser"] = users.GetRoleList(user.BussinessMen)

	err, progress := project.GetProgressList()
	if err == nil {
		c.Data["Progress"] = progress
	}

	c.Data["paginator"] = utils.NewPaginator(c.Ctx.Request, int(10), len(allProjects))
}

func (c *ProjectController) CreateProject() {
	var retJson = make(map[string]interface{})
	retJson["result"] = 0
	retJson["msg"] = ""

	var params = make(map[string]interface{})

	params["name"] = c.GetString("name")
	params["started"] = c.GetString("started")
	params["scale"] = c.GetString("scale")

	params["priority"] = c.GetString("priority")
	params["clientName"] = c.GetString("client_name")
	params["bussinessUser"] = c.GetString("bussiness_user")
	params["progress"] = c.GetString("progress")
	params["contractNo"] = c.GetString("contract_no")
	params["serviceItem"] = c.GetString("service_item")
	params["regStartDate"] = c.GetString("reg_start_date")
	params["gameDate"] = c.GetString("game_date")

	params["regCloseDate"] = c.GetString("reg_close_date")
	params["artUser"] = c.GetString("art_user")
	if params["artUser"] == "" {
		params["artUser"] = "0"
	}
	params["techUser"] = c.GetString("tech_user")

	if params["techUser"] == "" {
		params["techUser"] = "0"
	}
	//登记人
	params["registrant"] = c.currentUser.GetUser().Id
	project := project.NewProject()
	err, projectId := project.Create(params)
	if err != nil {
		retJson["result"] = -101
		retJson["msg"] = err.Error()
		fmt.Printf("err---->%v\n", err)
	}
	retJson["id"] = projectId
	c.Data["json"] = retJson
	c.ServeJson()
}

func (c *ProjectController) GetProjectNames() {
	projects, _ := project.NewProjectList().GetAllProjectNames()
	c.Data["ProjectNames"] = projects
	return
}
func (c *ProjectController) GetPriority() []map[string]string {
	return []map[string]string{
		{"Priority": "1"},
		{"Priority": "2"},
		{"Priority": "3"},
		{"Priority": "4"},
	}
}

func (c *ProjectController) GetExportFilter() map[string]interface{} {
	filter := make(map[string]interface{})
	id := c.GetString("project_name")
	if id != "" {
		filter["id"] = id
	}
	bussiness_user := c.GetString("bussiness_user")
	if bussiness_user != "" {
		filter["bussiness_user_id"] = bussiness_user
	}
	progress_filter := c.GetString("progress")
	if progress_filter != "" {
		filter["progress"] = progress_filter
	}
	start_date := c.GetString("start_date")
	if start_date != "" {
		filter["start_date"] = start_date
	}
	end_date := c.GetString("end_date")
	if end_date != "" {
		filter["end_date"] = end_date
	}
	return filter
}
