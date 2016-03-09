package job

import (
	"models"
	"time"
	"utils"
)

func GetTodayJobCount() (cnt int64) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.job.GetTodayJobCount : error : %s", err.Error())
		} else {
			utils.GetLog().Debug("services.job.GetTodayJobCount : count=%v", cnt)
		}
	}()

	q := models.GetDB().QueryTable(models.TABLE_NAME_JOB).Filter("created__gte", time.Now().Format("2006-01-02"))
	cnt, err = q.Count()

	return
}

func AddHistory(j *models.Job, user *models.User, isCreate bool) (jid uint, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.job.AddHistory : error : %s, ", err.Error())
		} else {
			utils.GetLog().Debug("services.job.AddHistory : debug : Job=%s", utils.Sdump(j))
		}
	}()

	var jh models.JobHistory
	jh.Job = j
	jh.IsCreate = isCreate
	jh.Code = j.Code
	jh.CreateUser = j.CreateUser
	jh.Project = j.Project
	jh.Employee = j.Employee
	jh.Type = j.Type
	jh.Department = j.Department
	jh.Target = j.Target
	jh.TargetUrl = j.TargetUrl
	jh.Desc = j.Desc
	jh.Message = j.Message
	jh.FinishTime = j.FinishTime
	jh.ValidTime = j.ValidTime
	jh.ClaimTime = j.ClaimTime
	jh.ValidStatus = j.ValidStatus
	jh.ClaimStatus = j.ClaimStatus
	jh.SubmitStatus = j.SubmitStatus

	var insertId int64
	insertId, err = models.GetDB().Insert(&jh)
	jid = uint(insertId)

	return
}
