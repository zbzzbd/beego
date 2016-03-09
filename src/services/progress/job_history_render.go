package progress

import (
	"models"
	"path"
	"services/upload"
	"utils"
)

type JobHistoryRender struct {
	data     *models.JobHistory
	files    []*models.File
	template string
}

func NewJobHistoryRender() *JobHistoryRender {
	return &JobHistoryRender{
		data:     &models.JobHistory{},
		template: path.Join(cardPath, "job_history.tpl"),
	}
}

func (r *JobHistoryRender) SetData(data *models.JobHistory) {
	r.data = data
	r.files, _ = upload.NewQueryFile().SetCondition("type__lte", upload.FT_Modify).SetCondition("rel_id", r.data.Id).GetFiles()
}

func (r *JobHistoryRender) GetFiles() []models.File {
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

func (r *JobHistoryRender) Render() (string, error) {
	data := struct {
		Id           uint
		IsCreate     bool
		Code         string
		Project      models.Project
		Employee     models.User
		CreateUser   models.User
		Type         string
		Department   string
		Target       string
		TargetUrl    string
		ValidTime    string
		ClaimTime    string
		FinishTime   string
		Desc         string
		Message      string
		ValidStatus  uint8
		ClaimStatus  uint8
		SubmitStatus uint8
		Updated      string

		Label string

		Files []models.File
	}{
		Id:           r.data.Id,
		IsCreate:     r.data.IsCreate,
		Code:         r.data.Code,
		Project:      *r.data.Project,
		Employee:     *r.data.Employee,
		CreateUser:   *r.data.CreateUser,
		Type:         r.data.Type,
		Department:   r.data.Department,
		Target:       r.data.Target,
		TargetUrl:    r.data.TargetUrl,
		ValidTime:    r.data.ValidTime.Format(utils.YMDHIS),
		ClaimTime:    r.data.ClaimTime.Format(utils.YMDHIS),
		FinishTime:   r.data.FinishTime.Format(utils.YMDHIS),
		Desc:         r.data.Desc,
		Message:      r.data.Message,
		ValidStatus:  r.data.ValidStatus,
		ClaimStatus:  r.data.ClaimStatus,
		SubmitStatus: r.data.SubmitStatus,

		Label: "创建作业",

		Updated: r.data.Updated.Format(utils.YMDHIS),

		Files: r.GetFiles(),
	}

	if data.IsCreate == false {
		data.Label = "修改作业"
	}

	return render(r.template, data)
}
