package controllers

type RoleController struct {
	BaseController
}

func (c *RoleController) Prepare() {
	c.BaseController.Prepare()
}

func (c *RoleController) Finish() {
	defer c.BaseController.Finish()
}

func (c *RoleController) Get() {
	c.TplNames = "role/error.tpl"
}
