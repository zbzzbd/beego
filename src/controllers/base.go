package controllers

import (
	"helpers"
	//	"log"
	"services/user"
	"strconv"
	"time"
	"utils"

	"github.com/astaxie/beego"
)

const (
	SiteTitle         = "项目管理系统"
	DefaultTimeLayout = "2006-01-02 15:04"
	StaticPathPhoto   = "photo"
	StatusActive      = "active"
)

type BaseController struct {
	beego.Controller
	currentUser      *user.User
	pageRequestStart time.Time
	pageReqestEnd    time.Duration
	isLogin          bool
	noNeedLogin      bool
	userId           int
	jsonResponse     map[string]interface{}
}

// Prepare 公用的函数, 作一请求初始化
func (c *BaseController) Prepare() {
	c.jsonResponse = make(map[string]interface{})
	c.Data["NeedTemplate"] = true
	c.pageRequestStart = time.Now()
	//log.Printf("request url : %s\n", c.Ctx.Request.URL.RawPath)

	c.setView()
	c.isLogin = c.IsLogin()
	c.setCurrentUser()

	c.checkRole()

	pageLabel := helpers.GetPageLabel(c.Ctx.Request.URL.RequestURI())
	c.Data["PageLabel"] = pageLabel
	c.SetTitle(pageLabel)
}

// Finish 在请求结束后需要手动调用
func (c *BaseController) Finish() {
	c.pageReqestEnd = time.Since(c.pageRequestStart)
	utils.GetLog().Trace("request finished in : %s", c.pageReqestEnd)
}

func (c *BaseController) SetTitle(title string) string {
	if title == "" {
		return SiteTitle
	}
	return SiteTitle + " | " + title
}

func (c *BaseController) setView() {
	beego.ViewsPath = "views"

	c.setViewLayout()
}

func (c *BaseController) setViewLayout() {
	c.LayoutSections = make(map[string]string)

	c.LayoutSections["Header"] = "layouts/header.tpl"
	c.LayoutSections["Footer"] = "layouts/footer.tpl"
	c.LayoutSections["LayoutMenu"] = "menu/main.tpl"

	c.Layout = "layouts/default.tpl"

	c.setCommonViewData()
}

func (c *BaseController) setCommonViewData() {
	c.Data["is_login"] = c.isLogin
	c.Data["IsDev"] = utils.IsDev()
	c.Data["title"] = c.SetTitle("")
}

func (c *BaseController) IsLogin() bool {
	return c.GetSession("user_id") != nil
}

func (c *BaseController) setCurrentUser() {
	if !c.isLogin {
		return
	}
	userId := c.GetSession("user_id").(uint)
	var err error

	defer func() {
		if err != nil {
			utils.GetLog().Error("controllers.BaseController.setCurrentUser : NewUserWithId error : userId = %v, noNeedLogin=%v", userId, c.noNeedLogin)
		}
	}()

	if c.currentUser, err = user.NewUserWithId(userId); err != nil {
		if !c.noNeedLogin {
			c.Redirect("/login", 302)
		}
		return
	}
	c.userId = int(userId)

	c.Data["user_info"] = c.currentUser.GetUser()

	c.setUserCookie()
}

func (c *BaseController) setUserCookie() {
	c.Ctx.SetCookie("uid", strconv.Itoa(int(c.currentUser.GetUser().Id)))
	c.Ctx.SetCookie("uname", c.currentUser.GetUser().Name)
}

func (c *BaseController) checkRole() {
	if c.Ctx.Request.Method != "GET" || c.currentUser == nil {
		return
	}

	strUrl := c.Ctx.Request.URL.RequestURI()
	role := helpers.NewRoleCheck(strUrl, c.currentUser.GetRole()).Check()
	switch role {
	case helpers.Role_Refuse:
		c.Data["RoleError"] = true
	case helpers.Role_ReadOnly:
		c.Data["IsReadonly"] = true
	}

}

func (c *BaseController) SetJsonResponse(key string, value interface{}) {
	c.jsonResponse[key] = value
}

func (c *BaseController) GetJsonResponse() *BaseController {
	c.Data["json"] = c.jsonResponse
	return c
}
