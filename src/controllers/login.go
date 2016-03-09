package controllers

import userSrv "services/user"

type LoginController struct {
	BaseController
}

func (c *LoginController) Prepare() {
	c.noNeedLogin = true
	c.BaseController.Prepare()

	c.LayoutSections["LayoutMenu"] = ""
	c.LayoutSections["Header"] = ""

	c.Data["NeedTemplate"] = false
}

func (c *LoginController) Finish() {
	defer c.BaseController.Finish()
}

func (c *LoginController) Get() {
	c.Data["title"] = c.SetTitle("登录")

	c.TplNames = "login/index.tpl"
}

func (c *LoginController) Post() {
	email := c.GetString("email")
	password := c.GetString("password")
	u := userSrv.NewUser()
	err := u.Login(email, password)
	if err != nil {
		c.SetJsonResponse("error", err.Error())
	} else {
		c.SetSession("user_id", u.GetUser().Id)
	}
	c.GetJsonResponse().ServeJson()
}
