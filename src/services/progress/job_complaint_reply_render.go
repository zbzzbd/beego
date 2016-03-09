package progress

import (
	"models"
	"path"
	"services/upload"
	"utils"
)

type JobComplaintReplyRender struct {
	data     *models.JobComplaintReply
	files    []*models.File
	template string
}

func NewJobComplaintReplyRender() *JobComplaintReplyRender {
	return &JobComplaintReplyRender{
		data:     &models.JobComplaintReply{},
		template: path.Join(cardPath, "job_complaint_reply.tpl"),
	}
}

func (r *JobComplaintReplyRender) SetData(data *models.JobComplaintReply) {
	r.data = data
	r.files, _ = upload.NewQueryFile().SetCondition("type", upload.FT_ComplainReplay).SetCondition("rel_id", r.data.Id).GetFiles()
}

func (r *JobComplaintReplyRender) GetFiles() []models.File {
	files := make([]models.File, len(r.files))
	for k, v := range r.files {
		files[k] = models.File{
			Type:  v.Type,
			RelId: v.RelId,
			Name:  v.Name,
			Url:   v.Url,
		}
	}
	return files
}

func (r *JobComplaintReplyRender) Render() (string, error) {
	data := struct {
		Id      uint
		Message string
		User    models.User
		Created string
		Files   []models.File
	}{
		Id:      r.data.Id,
		Message: r.data.Message,
		User:    *r.data.User,
		Created: r.data.Created.Format(utils.YMDHIS),
		Files:   r.GetFiles(),
	}

	return render(r.template, data)
}
