package controllers

import (
	"services/require"
)

type RequireController struct {
	BaseController
}

func (c *RequireController) Prepare() {
	c.BaseController.Prepare()
}

func (c *RequireController) Finish() {
	defer c.BaseController.Finish()
}

func (c *RequireController) List() {
	c.Data["requires"] = require.NewRequireList().GetRequireNames()
	c.Data["require"] = c.GetString("require")
	c.TplNames = "require/list.tpl"
}

func (c *RequireController) Create() {
	ret := make(map[string]interface{})

	name := c.GetString("name")
	_, err := require.NewRequireList().Add(name)
	if err != nil {
		ret["error"] = err.Error()
	}

	c.Data["json"] = ret
	c.ServeJson()
}

func (c *RequireController) Edit() {
	name := c.GetString("name")
	newName := c.GetString("newName")
	err := require.NewRequireList().Edit(name, newName)
	if err != nil {
		c.SetJsonResponse("error", err.Error())
	}

	c.GetJsonResponse().ServeJson()
}

func (c *RequireController) Delete() {
	name := c.GetString("name")
	_, err := require.NewRequireList().Delete(name)
	if err != nil {
		c.SetJsonResponse("error", err.Error())
	}

	c.GetJsonResponse().ServeJson()
}
