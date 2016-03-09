package upload

import (
	"models"
)

type QueryFile struct {
	conditions map[string]interface{}
}

func NewQueryFile() *QueryFile {
	conditions := make(map[string]interface{})
	return &QueryFile{conditions: conditions}
}

func (self *QueryFile) SetCondition(column string, value interface{}) *QueryFile {
	self.conditions[column] = value
	return self
}

func (self *QueryFile) GetFiles() (jobFiles []*models.File, err error) {
	var limit int = 20
	if self.conditions["limit"] != nil {
		limit = self.conditions["limit"].(int)
	}

	var offset int = 0
	if self.conditions["offset"] != nil {
		offset = self.conditions["offset"].(int)
	}

	q := models.GetDB().QueryTable(models.TABLE_NAME_FILE)
	for k, v := range self.conditions {
		q = q.Filter(k, v)
	}
	_, err = q.Limit(limit).Offset(offset).All(&jobFiles)

	return
}

func (self *QueryFile) GetJobAllFiles() (jobFiles []*models.File, err error) {
	_, err = models.GetDB().Raw("SELECT f.id, f.url, f.name FROM "+models.TABLE_NAME_FILE+
		" f join "+models.TABLE_NAME_JOBHISTORY+" h on f.rel_id=h.id join "+models.TABLE_NAME_JOB+
		" j on j.id=h.job_id WHERE j.id = ? and f.type<3", self.conditions["rel_id"]).QueryRows(&jobFiles)

	return
}

func (self *QueryFile) DelFile() (err error) {
	q := models.GetDB().QueryTable(models.TABLE_NAME_FILE)
	for k, v := range self.conditions {
		q = q.Filter(k, v)
	}
	_, err = q.Delete()

	return
}
