package controllers

import (
	"services/job_claim"
	"services/upload"
	"strconv"
)

type JobClaimController struct {
	JobController
}

func (c *JobClaimController) Prepare() {
	c.BaseController.Prepare()
}

func (c *JobClaimController) Finish() {
	defer c.BaseController.Finish()
}

func (c *JobClaimController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["error"] = ""
	c.TplNames = "produce/job_claim.tpl"

	c.getJobCards(id)
	c.Data["JobFiles"], _ = upload.NewQueryFile().SetCondition("rel_id", id).SetCondition("type", upload.FT_Claim).GetFiles()
}

func (c *JobClaimController) Post() {
	id := c.Ctx.Input.Param(":id")
	status, _ := c.GetInt("status")

	jobId, _ := strconv.Atoi(id)
	_, err := jobclaim.NewJobClaim(uint(jobId), c.currentUser.GetUser()).SetCondition("Status", status).SetCondition("Remark", c.GetString("remark")).Do()

	if err != nil {
		c.SetJsonResponse("error", err.Error())
	}

	c.GetJsonResponse().ServeJson()
}
