package main

import (
	"controllers"
	"flag"
	"log"
	"models"
	_ "routers"
	"utils"

	"helpers"

	"github.com/astaxie/beego"
)

var (
	env    = flag.String("env", "dev", "设置运行环境, 有dev, test, prod三种配置环境")
	syncdb = flag.Bool("syncdb", false, "set syncdb")
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	flag.Parse()
	log.Println("当前运行环境为: ", *env)
	utils.SetEnv(*env)

	if err := utils.InitConfig("./conf/"); err != nil {
		log.Fatal(err)
	}
	utils.InitLogger()
	utils.InitRander()

	models.InitModels(*syncdb)
	helpers.InitHelper()
}

func initHTTP() {
	beego.HttpPort, _ = utils.GetConf().Int("app::httpport")
	beego.CopyRequestBody, _ = utils.GetConf().Bool("app:copyrequestbody")
}

func initSession() {
	beego.SessionOn = true
	beego.SessionProvider = "file"
	beego.SessionGCMaxLifetime = 86400 * 30
	beego.SessionName = "sessionid"
	beego.SessionCookieLifeTime = 86400 * 30
	beego.SessionAutoSetCookie = true
	beego.SessionSavePath = "/tmp/achilles"
}

func initStatic() error {
	uploadPath := utils.GetConf().String("upload::root")
	err := utils.EnsurePath(uploadPath)
	if err != nil {
		return err
	}

	jobPath := uploadPath + utils.GetConf().String("upload::job")
	err = utils.EnsurePath(jobPath)
	if err != nil {
		return err
	}

	exportPath := uploadPath + utils.GetConf().String("upload::export")
	err = utils.EnsurePath(exportPath)
	if err != nil {
		return err
	}

	beego.DirectoryIndex = true
	beego.StaticDir["/public"] = "public/"
	beego.SetStaticPath("/img", "public/img")
	beego.SetStaticPath("/export", exportPath)
	beego.SetStaticPath("/"+controllers.StaticPathPhoto, uploadPath)

	return nil
}

func initApp() error {
	initHTTP()
	initSession()
	return initStatic()
}

func main() {
	if initApp() != nil {
		return
	}

	beego.Run()
}
