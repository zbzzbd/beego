package controllers

import (
	"models"
	"services/job"
	"services/job_assign"
	"services/user"
)

type JobAssignController struct {
	JobController
}

func (c *JobAssignController) Prepare() {
	c.BaseController.Prepare()
}

func (c *JobAssignController) Finish() {
	defer c.BaseController.Finish()
}

func (c *JobAssignController) GetUserList(department string) ([]*models.User, error) {
	//fork shitty codes!
	var filterRoleType user.RoleType
	if department == user.TechGuy.Desc() {
		filterRoleType = user.TechGuy
	} else if department == user.ArtGuy.Desc() {
		filterRoleType = user.ArtGuy
	}

	users := user.NewUserList().GetRoleListExcept(filterRoleType, []uint{uint(c.userId)})
	return users, nil
}

func (c *JobAssignController) Get() {
	c.TplNames = "produce/job_assign.tpl"
	id := c.GetString("id")
	jobItem, _ := job.NewJobList(1, 1).SetCondition("id", id).GetOne()
	c.Data["error"] = ""
	c.getJobCards(id)
	users, _ := c.GetUserList(jobItem.Department)
	c.Data["toUsers"] = users
	c.Data["jobId"] = id
}

func (c *JobAssignController) Post() {
	ret := make(map[string]interface{})
	ret["error"] = ""

	jobId, _ := c.GetInt("job_id")
	fromUser := c.userId
	toUser, _ := c.GetInt("to_user")
	remark := c.GetString("remark")

	ja := jobassign.NewJobAssignment(jobId, fromUser, toUser, remark)
	if err := ja.Do(); err != nil {
		c.SetJsonResponse("error", err.Error())
	}

	c.GetJsonResponse().ServeJson()
}
