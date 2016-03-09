package progress

import (
	"models"
	"strconv"
	"strings"
	"utils"
)

type ProgressList struct {
	jobId                           uint
	progresses                      []*models.JobProgress
	progressIds                     []uint
	groupedProgresses               map[string][]*models.JobProgress
	progressIdWithData              map[uint]interface{}
	progressIdWithRenderedTemplates map[uint]string
}

func NewProgressList(jobid string) *ProgressList {
	jobIdint, _ := strconv.Atoi(jobid)
	return &ProgressList{
		jobId:                           uint(jobIdint),
		progresses:                      []*models.JobProgress{},
		progressIds:                     []uint{},
		groupedProgresses:               make(map[string][]*models.JobProgress),
		progressIdWithData:              make(map[uint]interface{}),
		progressIdWithRenderedTemplates: make(map[uint]string),
	}
}

func (pl *ProgressList) fetchProgresses() (err error) {
	_, err = models.GetDB().QueryTable(models.TABLE_NAME_JOBPROGRESS).Filter("job_id", pl.jobId).OrderBy("-created").All(&pl.progresses)
	return
}

//显示的顺序
func (pl *ProgressList) getProgressIds() {
	for _, progress := range pl.progresses {
		pl.progressIds = append(pl.progressIds, progress.Id)
	}
}

func (pl *ProgressList) groupProgresses() {
	for _, progress := range pl.progresses {
		tableName := progress.EventTableName
		if _, ok := pl.groupedProgresses[tableName]; !ok {
			pl.groupedProgresses[tableName] = []*models.JobProgress{}
		}
		pl.groupedProgresses[tableName] = append(pl.groupedProgresses[tableName], progress)
	}
}

func (pl *ProgressList) fetchGroupedProgresses() (err error) {
	for tableName, progresses := range pl.groupedProgresses {

		primaryKeys := []uint{} //某一种tabneName里的所有id
		primaryKeys_progressIds := make(map[uint]uint)

		for _, progress := range progresses {
			primaryKeys = append(primaryKeys, progress.PrimaryKey)
			primaryKeys_progressIds[progress.PrimaryKey] = progress.Id
		}

		switch tableName {
		case models.TABLE_NAME_JOBHISTORY: //作业创建进程
			if err = pl.handleJobHistory(primaryKeys, primaryKeys_progressIds); err != nil {
				return
			}

		case models.TABLE_NAME_JOBVALID:
			if err = pl.handleJobValidation(primaryKeys, primaryKeys_progressIds); err != nil {
				return
			}

		case models.TABLE_NAME_JOBCLAIM:
			if err = pl.handleJobClaim(primaryKeys, primaryKeys_progressIds); err != nil {
				return
			}
		case models.TABLE_NAME_JOBSUBMIT:
			if err = pl.handleJobSubmit(primaryKeys, primaryKeys_progressIds); err != nil {
				return
			}
		case models.TABLE_NAME_JOBCOMPLAINT:
			if err = pl.handleJobComplaint(primaryKeys, primaryKeys_progressIds); err != nil {
				return
			}
		case models.TABLE_NAME_JOBCOMPLAINTREPLY:
			if err = pl.handleJobComplaintReply(primaryKeys, primaryKeys_progressIds); err != nil {
				return
			}
		case models.TABLE_NAME_JOBASSIGN:
			if err = pl.handleJobAssigment(primaryKeys, primaryKeys_progressIds); err != nil {
				return
			}
		}
	}
	return
}

func (pl *ProgressList) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("progress.ProgressList.Do : err = %s", err.Error())
		}
		//utils.GetLog().Debug("progress.progresslist.Do : obj = %s", utils.Sdump(pl))
	}()

	if err = pl.fetchProgresses(); err != nil {
		return
	}
	pl.getProgressIds()
	pl.groupProgresses()
	if err = pl.fetchGroupedProgresses(); err != nil {
		return
	}
	return
}

func (pl *ProgressList) RenderedTemplates() string {
	templates := []string{}
	for _, progressId := range pl.progressIds {
		templates = append(templates, pl.progressIdWithRenderedTemplates[progressId])
	}

	return strings.Join(templates, "")
}

func (pl *ProgressList) handleJobValidation(primaryKeys []uint, primaryKeys_progressIds map[uint]uint) (err error) {
	jvs := []*models.JobValid{}
	_, err = models.GetDB().QueryTable(models.TABLE_NAME_JOBVALID).Filter("id__in", primaryKeys).RelatedSel("OperationUser").All(&jvs)

	render := NewJobValidRender()

	for _, jv := range jvs {
		progressId := primaryKeys_progressIds[jv.Id]
		pl.progressIdWithData[progressId] = jv //纪录到全局的数据结构中

		render.SetData(jv)
		view, err := render.Render()
		if err != nil {
			return err
		}
		pl.progressIdWithRenderedTemplates[progressId] = view
	}
	return
}

func (pl *ProgressList) handleJobHistory(primaryKeys []uint, primaryKeys_progressIds map[uint]uint) (err error) {
	jhs := []*models.JobHistory{}
	_, err = models.GetDB().QueryTable(models.TABLE_NAME_JOBHISTORY).Filter("id__in", primaryKeys).
		RelatedSel("Project", "Employee", "CreateUser").All(&jhs)

	render := NewJobHistoryRender()

	for _, jh := range jhs {
		progressId := primaryKeys_progressIds[jh.Id]
		pl.progressIdWithData[progressId] = jh //纪录到全局的数据结构中

		render.SetData(jh)
		view, err := render.Render()
		if err != nil {
			return err
		}
		pl.progressIdWithRenderedTemplates[progressId] = view
	}
	return
}

func (pl *ProgressList) handleJobClaim(primaryKeys []uint, primaryKeys_progressIds map[uint]uint) (err error) {
	jhs := []*models.JobClaim{}
	_, err = models.GetDB().QueryTable(models.TABLE_NAME_JOBCLAIM).Filter("id__in", primaryKeys).RelatedSel("ClaimUser").All(&jhs)

	render := NewJobClaimRender()

	for _, jh := range jhs {
		progressId := primaryKeys_progressIds[jh.Id]
		pl.progressIdWithData[progressId] = jh //纪录到全局的数据结构中

		render.SetData(jh)
		view, err := render.Render()
		if err != nil {
			return err
		}
		pl.progressIdWithRenderedTemplates[progressId] = view
	}
	return
}

func (pl *ProgressList) handleJobSubmit(primaryKeys []uint, primaryKeys_progressIds map[uint]uint) (err error) {
	jhs := []*models.JobSubmit{}
	_, err = models.GetDB().QueryTable(models.TABLE_NAME_JOBSUBMIT).Filter("id__in", primaryKeys).RelatedSel("ProduceUser").All(&jhs)

	render := NewJobSubmitRender()
	for _, jh := range jhs {
		progressId := primaryKeys_progressIds[jh.Id]
		pl.progressIdWithData[progressId] = jh //纪录到全局的数据结构中

		render.SetData(jh)
		view, err := render.Render()
		if err != nil {
			return err
		}
		pl.progressIdWithRenderedTemplates[progressId] = view
	}
	return
}

func (pl *ProgressList) handleJobComplaint(primaryKeys []uint, primaryKeys_progressIds map[uint]uint) (err error) {
	jhs := []*models.JobComplaint{}
	_, err = models.GetDB().QueryTable(models.TABLE_NAME_JOBCOMPLAINT).Filter("id__in", primaryKeys).RelatedSel("User").All(&jhs)
	render := NewJobComplaintRender()

	for _, jh := range jhs {
		progressId := primaryKeys_progressIds[jh.Id]
		pl.progressIdWithData[progressId] = jh //纪录到全局的数据结构中

		render.SetData(jh)
		view, err := render.Render()
		if err != nil {
			return err
		}
		pl.progressIdWithRenderedTemplates[progressId] = view
	}
	return
}

func (pl *ProgressList) handleJobComplaintReply(primaryKeys []uint, primaryKeys_progressIds map[uint]uint) (err error) {
	jhs := []*models.JobComplaintReply{}
	_, err = models.GetDB().QueryTable(models.TABLE_NAME_JOBCOMPLAINTREPLY).Filter("id__in", primaryKeys).RelatedSel("User").All(&jhs)

	render := NewJobComplaintReplyRender()

	for _, jh := range jhs {
		progressId := primaryKeys_progressIds[jh.Id]
		pl.progressIdWithData[progressId] = jh //纪录到全局的数据结构中

		render.SetData(jh)
		view, err := render.Render()
		if err != nil {
			return err
		}
		pl.progressIdWithRenderedTemplates[progressId] = view
	}
	return
}

func (pl *ProgressList) handleJobAssigment(primaryKeys []uint, primaryKeys_progressIds map[uint]uint) (err error) {
	jhs := []*models.JobAssign{}
	_, err = models.GetDB().QueryTable(models.TABLE_NAME_JOBASSIGN).Filter("id__in", primaryKeys).RelatedSel("Job", "FromUser", "ToUser").All(&jhs)

	render := NewJobAssignRender()

	for _, jh := range jhs {
		progressId := primaryKeys_progressIds[jh.Id]
		pl.progressIdWithData[progressId] = jh //纪录到全局的数据结构中

		render.SetData(jh)
		view, err := render.Render()
		if err != nil {
			return err
		}
		pl.progressIdWithRenderedTemplates[progressId] = view
	}
	return
}
