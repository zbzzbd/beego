package progress

import (
	"models"
	"path"
	"utils"
	"services/upload"
)

type JobSubmitRender struct {
	data     *models.JobSubmit
	files    []*models.File
	template string
}

func NewJobSubmitRender() *JobSubmitRender {
	return &JobSubmitRender{
		data:     &models.JobSubmit{},
		template: path.Join(cardPath, "job_submit.tpl"),
	}
}

func (r *JobSubmitRender) SetData(data *models.JobSubmit) {
	r.data = data
	r.files, _ = upload.NewQueryFile().SetCondition("type", upload.FT_Submit).SetCondition("rel_id", r.data.Id).GetFiles()
}

func (r *JobSubmitRender) GetFiles() ([]models.File) {
	files := make([]models.File, len(r.files))
	for k, v := range r.files {
		files[k] = models.File{
			Type: v.Type,
			RelId: v.RelId,
			Name: v.Name,
			Url: v.Url,
		}
	}

	return files
}

func (r *JobSubmitRender) Render() (string, error) {
	strStatus := make(map[uint8]string)
	strStatus[1] = "完成任务，提交作业"
	strStatus[2] = "作业验收不通过，请重新制作"
	strStatus[3] = "Good，作业验收通过"

	data := struct {
		Id        uint
		Job       models.Job
		ProduceUser models.User
		Status    string
		Remark    string

		Created  string

		Label    string

		Files []models.File
	}{
		Id          : r.data.Id,
		Job         : *r.data.Job,
		ProduceUser   : *r.data.ProduceUser,
		Status      : strStatus[r.data.Status],
		Remark      : r.data.Remark,
		Created     : r.data.Created.Format(utils.YMDHIS),

		Label       : "作业提交",

		Files:   r.GetFiles(),
	}

	if r.data.Status > 1 {
		data.Label = "作业验收"
	}

	return render(r.template, data)
}
