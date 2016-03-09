package project

import (
	"models"
)

type Progress struct {
	progress *models.Progress
}

func NewProgress() *Progress {
	return &Progress{
		progress: &models.Progress{},
	}
}
func GetProgressList() (err error, progressList []*models.Progress) {

	err, progressList = models.GetProgressList()
	return
}

func (p *Progress) GetProgressById(id string) (*Progress, error) {
	err := p.progress.GetProgressById(id)
	return p, err
}

func (p *Progress) GetModelProgress() *models.Progress {
	return p.progress
}
