package controllers

import (
	"errors"
	"strconv"

	"services/job"
	"services/project"
	"services/user"
	"utils"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Prepare() {
	c.BaseController.Prepare()
}

func (c *IndexController) Finish() {
	defer c.BaseController.Finish()
}

func (c *IndexController) UserType() int {
	if c.currentUser.GetRole().HasRole(string(user.ProjectManager)) {
		return 1
	} else if c.currentUser.GetRole().HasRole(string(user.BussinessMen)) {
		return 2
	} else if c.currentUser.GetRole().HasRole(string(user.ArtGuy)) || c.currentUser.GetRole().HasRole(string(user.TechGuy)) {
		return 3
	} else if c.currentUser.GetRole().HasRole(string(user.Customer)) {
		return 4
	} else if c.currentUser.GetRole().HasRole(string(user.Interactive)) {
		return 5
	} else {
		return 0
	}
}

func (c *IndexController) SetTpl() error {
	user_type := c.UserType()
	if user_type == 1 {
		//订单管理员
		c.TplNames = "index/projectManager.tpl"
	} else if user_type == 2 {
		//业务单元
		c.TplNames = "index/bussiness.tpl"
	} else if user_type == 3 {
		//制作单元
		c.TplNames = "index/produce.tpl"
	} else if user_type == 4 {
		c.TplNames = "index/customer.tpl"
	} else if user_type == 5 {
		c.TplNames = "index/interactive.tpl"
	} else {
		return errors.New("err user_type")
	}
	return nil
}
func (c *IndexController) SetJobsDate() {
	user_type := c.UserType()
	if user_type == 1 {
		var jobs job.JobsCountForManager
		c.Data["Job"] = jobs.Get()
	} else if user_type == 2 || user_type == 4 || user_type == 5 {
		var jobs job.JobsCountForBussinessman
		c.Data["Job"] = jobs.Get(strconv.Itoa(int(c.currentUser.GetUser().Id)))
	} else {
		var jobs job.JobsCountForTech
		c.Data["Job"] = jobs.Get(strconv.Itoa(int(c.currentUser.GetUser().Id)))
	}
}
func (c *IndexController) SetProjectsDate() {
	filter := make(map[string]interface{})
	filter["limit"] = 10
	filter["offset"] = 0
	p, _ := c.GetInt("p")
	if p > 1 {
		filter["offset"] = (p - 1) * filter["limit"].(int)
	}
	manager_user_id := c.currentUser.GetUser().Id
	filter["user"] = manager_user_id
	pl := project.NewProjectList()
	projects, err := pl.GetList(filter)
	if err != nil {
		c.Data["error"] = err.Error()
		return
	}
	c.Data["Projects"] = projects
	count := pl.Count()
	c.Data["paginator"] = utils.NewPaginator(c.Ctx.Request, int(10), count)
}

func (c *IndexController) Get() {
	if !c.isLogin {
		c.Redirect("/login", 302)
	}
	if c.SetTpl() != nil {
		c.Redirect("/login", 302)
	}
	c.SetProjectsDate()

	c.SetJobsDate()
}
