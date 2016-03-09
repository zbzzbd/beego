package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"models"
	"os"
	_ "routers"
	"services/job"
	"services/progress"
	"services/user"
	"time"
	"utils"
)

var (
	env    = flag.String("env", "dev", "设置运行环境, 有dev, test, prod三种配置环境")
	syncdb = flag.Bool("syncdb", true, "是否同步数据库表结构")
)

func init() {
	flag.Parse()
	log.Println("当前运行环境为: ", *env)
	utils.SetEnv(*env)

	if err := fixConfigFile(); err != nil {
		log.Fatal(err)
	}
	if err := utils.InitConfig("./../../conf/"); err != nil {
		log.Fatal(err)
	}
	utils.InitLogger()
	utils.InitRander()
	models.InitModels(*syncdb)
}

func fixConfigFile() (err error) {
	env := fmt.Sprintf("%s", utils.GetEnv())
	envConfFile, err := os.Open("./../../conf/" + env + ".conf")
	defer envConfFile.Close()
	appConfFile, err := os.Create("./conf/app.conf")
	defer appConfFile.Close()
	io.Copy(appConfFile, envConfFile)
	return
}

func main() {
	CreateCompanies()
	CreateUsers()
	CreateProgress()
	CreateRequires()
	//	CreateProjects()
	//	CreateJobs()
	//	CreateJobProgresses()
}

func CreateCompanies() {
	c := []models.Company{
		models.Company{
			Code:    "SH",
			Name:    "上海自由博爱",
			Address: "上海松江",
			City:    "上海",
		},
		models.Company{
			Code:    "YZ",
			Name:    "扬州自由博爱",
			Address: "江苏扬州",
			City:    "扬州",
		},
		models.Company{
			Code:    "CQ",
			Name:    "重庆技术部",
			Address: "重庆",
			City:    "重庆",
		},
	}
	models.GetDB().InsertMulti(len(c), &c)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(name, email, password string, companyId int, roles []user.RoleType) (err error) {
	c := user.NewCreation(name, email, password, companyId, roles)
	log.Println(utils.Sdump(c))
	if err = c.Do(); err != nil {
		return
	}
	return
}

func CreateProgress() {
	d := []models.Progress{
		models.Progress{
			Id:   1,
			Name: "立项",
		},
		models.Progress{
			Id:   2,
			Name: "美术设计",
		},
		models.Progress{
			Id:   3,
			Name: "网站搭建",
		},

		models.Progress{
			Id:   4,
			Name: "制作完毕",
		},
		models.Progress{
			Id:   5,
			Name: "开通报名",
		},
		models.Progress{
			Id:   6,
			Name: "关闭报名",
		},
		models.Progress{
			Id:   7,
			Name: "照片下载(证书下载)",
		},
	}
	num, err := models.GetDB().InsertMulti(1, d)
	log.Println("CreateProgress: num=%v, err=%v", num, err)
}

func CreateProjects() {
	d := models.Project{
		Name:          "宝马项目",
		Scale:         100,
		Priority:      1,
		ClientName:    "xh",
		Started:       time.Now(),
		BussinessUser: &models.User{Id: 2},
		ArtUser:       &models.User{Id: 5},
		TechUser:      &models.User{Id: 4},
		Registrant:    &models.User{Id: 1},
		Progress:      &models.Progress{Id: 1},
		GameDate:      time.Now(),
		RegStartDate:  time.Now(),
		RegCloseDate:  time.Now(),
	}
	num, err := models.GetDB().Insert(&d)

	d1 := models.Project{
		Name:          "官网项目",
		Scale:         100,
		Priority:      1,
		ClientName:    "gw",
		Started:       time.Now(),
		BussinessUser: &models.User{Id: 3},
		ArtUser:       &models.User{Id: 5},
		TechUser:      &models.User{Id: 4},
		Registrant:    &models.User{Id: 1},
		Progress:      &models.Progress{Id: 1},
		GameDate:      time.Now(),
		RegStartDate:  time.Now(),
		RegCloseDate:  time.Now(),
	}
	num, err = models.GetDB().Insert(&d1)

	log.Println("CreateProjects: num=%v, err=%v", num, err)
}

func CreateJobs() {
	d := models.Job{
		Code: "20151111SH1",
		Project: &models.Project{
			Id: 1,
		},
		Employee:   &models.User{Id: 5},
		CreateUser: &models.User{Id: 2},
		Type:       "模板制作",
		Department: "美术",
		Target:     "模板制作",
		TargetUrl:  "http://www.chinarun.com",
		Desc:       "模板制作模板制作模板制作",
		Message:    "模板制作",
		Created:    time.Now().AddDate(0, 0, 1),
		FinishTime: time.Now().AddDate(0, 0, 1),
	}
	num, err := models.GetDB().Insert(&d)
	job.AddHistory(&d, &models.User{Id: 1}, true)

	d = models.Job{
		Code: "20151111SH2",
		Project: &models.Project{
			Id: 2,
		},
		Employee:   &models.User{Id: 4},
		CreateUser: &models.User{Id: 3},
		Type:       "Logo更新",
		Department: "技术",
		Target:     "Logo更新",
		TargetUrl:  "http://www.chinarun.com",
		Desc:       "Logo更新Logo更新",
		Message:    "Logo更新",
		Created:    time.Now().AddDate(0, 0, 2),
		FinishTime: time.Now().AddDate(0, 0, 2),
	}
	num, err = models.GetDB().Insert(&d)
	job.AddHistory(&d, &models.User{Id: 2}, true)

	d = models.Job{
		Code: "20151111SH3",
		Project: &models.Project{
			Id: 2,
		},
		Employee:   &models.User{Id: 4},
		CreateUser: &models.User{Id: 3},
		Type:       "报名通道关闭",
		Department: "技术",
		Target:     "报名通道关闭",
		TargetUrl:  "http://www.chinarun.com",
		Desc:       "Logo更新Logo更新",
		Message:    "报名通道关闭",
		Created:    time.Now().AddDate(0, 0, 3),
		FinishTime: time.Now().AddDate(0, 0, 3),
	}
	num, err = models.GetDB().Insert(&d)
	job.AddHistory(&d, &models.User{Id: 3}, true)

	d = models.Job{
		Code: "20151111SH4",
		Project: &models.Project{
			Id: 1,
		},
		Employee:   &models.User{Id: 5},
		CreateUser: &models.User{Id: 2},
		Type:       "后台资料查询",
		Department: "美术",
		Target:     "后台资料查询",
		TargetUrl:  "http://www.chinarun.com",
		Desc:       "后台资料查询",
		Message:    "后台资料查询",
		Created:    time.Now().AddDate(0, 0, 4),
		FinishTime: time.Now().AddDate(0, 0, 4),
	}
	num, err = models.GetDB().Insert(&d)
	job.AddHistory(&d, &models.User{Id: 4}, true)

	d = models.Job{
		Code: "20151111SH5",
		Project: &models.Project{
			Id: 2,
		},
		Employee:   &models.User{Id: 4},
		CreateUser: &models.User{Id: 2},
		Type:       "链接更新",
		Department: "技术",
		Target:     "链接更新",
		TargetUrl:  "http://www.chinarun.com",
		Desc:       "链接更新链接更新链接更新链接更新链接更新",
		Message:    "链接更新",
		Created:    time.Now().AddDate(0, 0, 5),
		FinishTime: time.Now().AddDate(0, 0, 5),
	}
	num, err = models.GetDB().Insert(&d)
	job.AddHistory(&d, &models.User{Id: 5}, true)

	log.Println("CreateJobs: num=%v, err=%v", num, err)
}

func CreateJobProgresses() {
	createJobProgress(1, progress.PT_Create, "创建作业", "job_histories", 1)
	createJobProgress(2, progress.PT_Create, "创建作业", "job_histories", 2)
	createJobProgress(3, progress.PT_Create, "创建作业", "job_histories", 3)
	createJobProgress(4, progress.PT_Create, "创建作业", "job_histories", 4)
	createJobProgress(5, progress.PT_Create, "创建作业", "job_histories", 5)
}

func createJobProgress(jobid int, pt progress.ProgressType, desc string, tableName string, pk int) {
	p := models.JobProgress{
		JobId:          uint(jobid),
		ProgressType:   uint(pt),
		Desc:           desc,
		EventTableName: tableName,
		PrimaryKey:     uint(pk),
	}
	models.GetDB().Insert(&p)
}
