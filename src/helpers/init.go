package helpers

import "github.com/astaxie/beego"

func InitHelper() {
	beego.AddFuncMap("TimeFormat", TimeFormat)
	beego.AddFuncMap("Strcat", Strcat)
	beego.AddFuncMap("ShowMenu", ShowMenu)
	beego.AddFuncMap("AddInt", AddInt)
	beego.AddFuncMap("GetStatusDesc", GetStatusDesc)
	beego.AddFuncMap("GetStatusColor", GetStatusColor)
	beego.AddFuncMap("GetTimeDiff", GetTimeDiff)
	beego.AddFuncMap("GetRoleDesc", GetRoleDesc)
	beego.AddFuncMap("RequiredStar", RequiredStar)
	beego.AddFuncMap("TimeIsZero", TimeIsZero)
}
