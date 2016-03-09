package progress

import (
	"models"
	"path"
	"utils"
)

type JobClaimRender struct {
	data     *models.JobClaim
	template string
}

func NewJobClaimRender() *JobClaimRender {
	return &JobClaimRender{
		data:     &models.JobClaim{},
		template: path.Join(cardPath, "job_claim.tpl"),
	}
}

func (r *JobClaimRender) SetData(data *models.JobClaim) {
	r.data = data
}

func (r *JobClaimRender) Render() (string, error) {
	strStatus := make(map[uint8]string)
	strStatus[1] = "认领任务"
	strStatus[2] = "任务无法完成，拒绝认领"

	data := struct {
		Id        uint
		Job       models.Job
		ClaimUser models.User
		Status    string
		Remark    string

		Created  string
	}{
		Id          : r.data.Id,
		Job         : *r.data.Job,
		ClaimUser   : *r.data.ClaimUser,
		Status      : strStatus[r.data.Status],
		Remark      : r.data.Remark,
		Created     : r.data.Created.Format(utils.YMDHIS),
	}
	return render(r.template, data)
}
