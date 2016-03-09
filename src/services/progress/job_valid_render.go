package progress

import (
	"models"
	"path"
	"utils"
)

type JobValidRender struct {
	data     *models.JobValid
	template string
}

func NewJobValidRender() *JobValidRender {
	return &JobValidRender{
		data:     &models.JobValid{},
		template: path.Join(cardPath, "job_valid.tpl"),
	}
}

func (r *JobValidRender) SetData(data *models.JobValid) {
	r.data = data
}

func (r *JobValidRender) Render() (string, error) {
	var finishTimeString string
	if !r.data.FinishTime.IsZero() {
		finishTimeString = r.data.FinishTime.Format(utils.YMDHIS)
	}

	jobValidationResult := map[int]string{
		1: "内容明确, 作业排产",
		2: "作业有误, 请重新修改",
		3: "作业取消",
	}

	data := struct {
		UserName   string
		Result     string
		Message    string
		FinishTime string
		Created    string
	}{
		UserName:   r.data.OperationUser.Name,
		Result:     jobValidationResult[int(r.data.Status)],
		Message:    r.data.Message,
		FinishTime: finishTimeString,
		Created:    r.data.Created.Format(utils.YMDHIS),
	}
	return render(r.template, data)
}
