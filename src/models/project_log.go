package models

import (
	"time"
)

type ProjectLog struct {
	Id            uint      `orm:"pk;auto"`
	Project       *Project  `orm:"column(project_id);rel(fk)"`
	User          *User     `orm:"column(user_id);rel(fk)"`
	Operation     uint      `orm:"column(project_operation)"`
	ProjectColumn string    `orm:"column(project_column);null"`
	From          string    `orm:"null"`
	To            string    `orm:"null"`
	Time          time.Time `orm:"auto_now_add;type(datetime)"`
}

func (p *ProjectLog) TableName() string {
	return "project_logs"
}

func GetProjectLogForEdit(projectId, userId uint) *ProjectLog {
	var projectLog ProjectLog
	projectLog.Project = &Project{Id: projectId}
	projectLog.User = &User{Id: userId}
	projectLog.Operation = 2

	return &projectLog
}

func (p *ProjectLog) SetColumn(column, from, to string) *ProjectLog {
	p.ProjectColumn = column
	p.From = from
	p.To = to
	return p
}
