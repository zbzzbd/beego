package controllers

import (
	"services/job_submit"
	"services/upload"
)

type JobSubmitController struct {
	JobController
}

func (c *JobSubmitController) Prepare() {
	c.BaseController.Prepare()
}

func (c *JobSubmitController) Finish() {
	defer c.BaseController.Finish()
}

func (c *JobSubmitController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["error"] = ""
	c.Data["status"], _ = c.GetInt("status")
	c.TplNames = "produce/job_submit.tpl"

	c.getJobCards(id)

	c.Data["JobFiles"], _ = upload.NewQueryFile().SetCondition("rel_id", id).SetCondition("type", upload.FT_Submit).GetFiles()
}

func (c *JobSubmitController) Post() {
	id := c.Ctx.Input.Param(":id")
	status, _ := c.GetInt("status")

	submitId, err := jobsubmit.NewJobSubmit(c.currentUser.GetUser()).
		SetCondition("Id", id).SetCondition("Status", status).SetCondition("Remark", c.GetString("remark")).Do()

	if err != nil {
		c.SetJsonResponse("error", err.Error())

	}
	c.uploadFiles(submitId, upload.FT_Submit)
	c.GetJsonResponse().ServeJson()
}
