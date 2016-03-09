package models

import "time"

const TABLE_NAME_PROJECT = "projects"

const (
	CREATE = 1
	EDIT   = 2
)

const (
	NAME           = "name"
	SCALE          = "scale"
	PRIORITY       = "priority"
	CLIENT_NAME    = "client_name"
	CONRACT_NO     = "contract_no"
	SERVICE_ITEM   = "service_item"
	GAME_DATE      = "game_date"
	REG_START_DATE = "reg_start_date"
	REG_CLOSE_DATA = "reg_close_date"
	STARTED        = "started"
	PROGRESS       = "progress"
	BUSSINESS_USER = "bussiness_user"
	ART_USER_ID    = "art_user_id"
	TECH_USER_ID   = "tech_user_id"
	REGISTRANT     = "registrant"
)

type Project struct {
	Id           uint      `orm:"pk;auto"`
	Name         string    `orm:"size(128);unique"`
	Scale        uint      `orm:"default(0);null"`
	Priority     uint      `orm:"default(0);null"`
	ClientName   string    `orm:"size(128);null"`
	ContractNo   string    `orm:"size(128);null"`
	ServiceItem  string    `orm:"size(128);null"`
	GameDate     time.Time `orm:"auto_now_add;type(datetime)"`
	RegStartDate time.Time `orm:"auto_now_add;type(datetime)"`
	RegCloseDate time.Time `orm:"auto_now_add;type(datetime)"`
	Started      time.Time `orm:"auto_now_add;type(datetime)"`
	Progress     *Progress `orm:"column(progress);rel(fk);null"`

	JobsNum uint `orm:"-"`

	BussinessUser *User `orm:"column(bussiness_user_id);rel(fk)"`
	ArtUser       *User `orm:"column(art_user_id);rel(fk);null"`
	TechUser      *User `orm:"column(tech_user_id);rel(fk);null"`

	Registrant *User `orm:"column(registrant);rel(fk)"`

	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
	DelStatus uint      `orm:"default(0)"` //状态0 表示没有被删除的Project  1 表示已经被删除的project
}

func (p *Project) TableName() string {
	return TABLE_NAME_PROJECT
}
