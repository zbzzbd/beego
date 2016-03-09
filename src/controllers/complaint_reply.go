package controllers

import (
	"services/job_complaint"
	"services/upload"
)

type ReplyController struct {
	ComplaintController
}

func (c *ReplyController) Prepare() {
	c.BaseController.Prepare()
}

func (c *ReplyController) Finish() {
	defer c.BaseController.Finish()
}

//回复页面
func (c *ReplyController) Show() {
	c.TplNames = "complaint/reply_show.tpl"

	jobId := c.GetString("jobid")
	c.getJobCards(jobId)

	id := c.Ctx.Input.Param(":id")
	c.Data["error"] = ""
	c.Data["Id"] = id

	c.Data["JobFiles"], _ = upload.NewQueryFile().SetCondition("type", upload.FT_ComplainReplay).SetCondition("rel_id", id).GetFiles()
}

//创建一个回复
func (c *ReplyController) CreateReply() {

	postData := make(map[string]interface{})
	id := c.Ctx.Input.Param(":id")
	postData["Desc"] = c.GetString("desc")
	postData["Complain_id"] = id

	repId, jobId, err := complaint.NewComplaintReplytor(postData, c.currentUser.GetUser()).Do()

	if err != nil {
		c.SetJsonResponse("error", err.Error())
	} else {
		//更新投诉的状态
		_, errUp := complaint.UpdateComplaint(postData)
		if errUp != nil {
			c.SetJsonResponse("errUp", errUp.Error())
		}

		c.SetJsonResponse("replyId", repId)
		c.SetJsonResponse("id", id)
		c.SetJsonResponse("jobid", jobId)
		c.uploadFiles(repId, upload.FT_ComplainReplay)
		//upload.NewUploadFile(c.Ctx.Request, "files", "job").SetFileType(upload.FT_ComplainReplay).SetRelId(jobId).Do()
	}

	c.GetJsonResponse().ServeJson()
}
