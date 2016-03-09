package models

import (
	"fmt"
	"utils"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var db orm.Ormer

func GetDB() orm.Ormer {
	return db
}

func showSql() {
	if utils.IsDev() {
		orm.Debug = true
	}
}

func connectDB() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	name := utils.GetConf().String("db::name")
	user := utils.GetConf().String("db::user")
	pass := utils.GetConf().String("db::pass")
	host := utils.GetConf().String("db::host")
	port := utils.GetConf().String("db::port")

	//data source name

	//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, host, port, name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", user, pass, host, port, name)
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModel(new(User), new(Project), new(ProjectLog))
	orm.RegisterModel(new(Company), new(Job), new(Progress))
	orm.RegisterModel(new(File), new(JobComplaint), new(JobProgress))
	orm.RegisterModel(new(JobValid), new(JobClaim), new(JobSubmit))
	orm.RegisterModel(new(UserLogin), new(JobAssign))
	orm.RegisterModel(new(JobHistory), new(JobComplaintReply))
	orm.RegisterModel(new(Require))
	db = orm.NewOrm()
	setMaxIdleConns()
}

func setMaxIdleConns() {
	maxIdleConns, _ := utils.GetConf().Int("db::max_idle_conns")
	orm.SetMaxIdleConns("default", maxIdleConns)
}

func syncTables() {
	if !utils.IsProd() {
		err := orm.RunSyncdb("default", true, true)
		if err != nil {
			panic("数据库同步出错:" + err.Error())
		}
	}
}

func InitModels(syncdb bool) {
	showSql()
	connectDB()

	if syncdb {
		syncTables()
	}
}
