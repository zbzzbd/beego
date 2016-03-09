package progress

import (
	"models"
	"path"
	"utils"
)

type JobComplaintRender struct {
	data     *models.JobComplaint
	template string
}

func NewJobComplaintRender() *JobComplaintRender {
	return &JobComplaintRender{
		data:     &models.JobComplaint{},
		template: path.Join(cardPath, "job_complaint.tpl"),
	}
}

func (r *JobComplaintRender) SetData(data *models.JobComplaint) {
	r.data = data
}

func (r *JobComplaintRender) Render() (string, error) {
	data := struct {
		Id         uint
		CreateUser models.User
		Employee   models.User
		Complain   string
		Type       string
		Response   uint
		Created    string
		EditStatus uint
	}{
		Id:         r.data.Id,
		CreateUser: *r.data.User,
		Employee:   *r.data.Employee,
		Complain:   r.data.Complain,
		Type:       r.data.Type,
		Response:   r.data.Response,
		Created:    r.data.Created.Format(utils.YMDHIS),
		EditStatus: r.data.EditStatus,
	}

	return render(r.template, data)
}
