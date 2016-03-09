package controllers

import (
	"forms"
	"models"
	"services/job"
	"services/job_valid"
)

type JobValidationController struct {
	JobController
}

func (c *JobValidationController) Prepare() {
	c.BaseController.Prepare()
}

func (c *JobValidationController) Finish() {
	defer c.BaseController.Finish()
}

func (c *JobValidationController) Get() {
	c.TplNames = "job_valid/index.tpl"

	conditions := map[string]interface{}{
		"valid_status": 0,
		"del_status":   models.Job_DelStatus_effective,
	}
	jl, _ := c.findJobList(conditions)
	c.makePaginator(jl)
	c.setFindConditions()
}

func (c *JobValidationController) Show() {
	c.TplNames = "job_valid/show.tpl"

	id := c.Ctx.Input.Param(":id")
	c.getJobCards(id)

	c.Data["job"], _ = job.NewJobList(1, 1).SetCondition("id", id).GetOne()
}

func (c *JobValidationController) Post() {
	jvf, err := forms.NewJobValidForm(c.Ctx.Request)
	if err != nil {
		c.SetJsonResponse("error", err.Error())
		c.GetJsonResponse().ServeJson()
		return
	}

	jvf.CurrentUserId = uint(c.userId)
	jv := jobvalid.NewJobValidation(jvf)
	if err := jv.Do(); err != nil {
		c.SetJsonResponse("error", err.Error())
		c.GetJsonResponse().ServeJson()
		return
	}

	c.GetJsonResponse().ServeJson()
}
