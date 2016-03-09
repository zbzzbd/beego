package progress

import (
	"models"
	"path"
	"utils"
)

type JobAssignRender struct {
	data     *models.JobAssign
	template string
}

func NewJobAssignRender() *JobAssignRender {
	return &JobAssignRender{
		data:     &models.JobAssign{},
		template: path.Join(cardPath, "job_assign.tpl"),
	}
}

func (r *JobAssignRender) SetData(data *models.JobAssign) {
	r.data = data
}

func (r *JobAssignRender) Render() (string, error) {

	data := struct {
		Id       uint
		Job      models.Job
		FromUser models.User
		ToUser   models.User
		Remark   string
		Created  string
	}{
		Id:       r.data.Id,
		Job:      *r.data.Job,
		FromUser: *r.data.FromUser,
		ToUser:   *r.data.ToUser,
		Remark:   r.data.Remark,
		Created:  r.data.Created.Format(utils.YMDHIS),
	}
	return render(r.template, data)
}
