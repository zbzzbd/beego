package progress

import (
	"fmt"
	"models"
	"utils"
)

var registerProgressTypes map[ProgressType]bool

func init() {
	registerProgressTypes = make(map[ProgressType]bool)
}

func Register(ptype ProgressType) {
	if _, ok := registerProgressTypes[ptype]; ok {
		panic("RegisterProgressType已经被注册过了")
	} else {
		registerProgressTypes[ptype] = true
	}
}

// 描述一个项目的进程列表
type Progress struct {
	progress JobProgresser
}

func NewProgress(progress JobProgresser) *Progress {
	return &Progress{
		progress: progress,
	}
}

func (p *Progress) AddOne() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("progress.Progress.AddOne : error =  %s", err.Error())
		}
	}()

	if err = p.validProgressType(); err != nil {
		return
	}

	if err = p.save(); err != nil {
		return
	}

	return
}

func (p *Progress) save() (err error) {
	jp := &models.JobProgress{
		JobId:          p.progress.GetJobId(),
		EventTableName: p.progress.GetTableName(),
		ProgressType:   uint(p.progress.GetType()),
		Desc:           p.progress.GetDesc(),
		PrimaryKey:     p.progress.GetPrimaryKey(),
	}
	err = jp.Insert()
	return
}

func (p *Progress) validProgressType() (err error) {
	ptype := p.progress.GetType()
	if _, ok := registerProgressTypes[ptype]; !ok {
		err = fmt.Errorf("RegisterProgressType : %d 还未注册", ptype)
		return
	}
	return
}
