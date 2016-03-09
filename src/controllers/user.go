package controllers

import (
	"fmt"
	"models"
	"services/company"
	"services/user"
	"strconv"
	"time"
	"utils"
)

type UserController struct {
	BaseController
}

func (c *UserController) Prepare() {
	c.BaseController.Prepare()
}

func (c *UserController) Finish() {
	defer c.BaseController.Finish()
}

func (c *UserController) Get() {
	c.TplNames = "login/modify_password.tpl"
}

func (c *UserController) ModifyPassword() {
	oldPassword := c.GetString("oldPassword")
	newPassword := c.GetString("newPassword")
	newPasswordConfirm := c.GetString("newPassword_confirm")
	fmt.Println(oldPassword, newPassword, newPasswordConfirm)
	err := user.NewPasswordModify(c.currentUser, oldPassword, newPassword, newPasswordConfirm).Do()
	if err != nil {
		c.SetJsonResponse("error", err.Error())
	}
	c.SetJsonResponse("id", int(c.currentUser.GetUser().Id))

	c.GetJsonResponse().ServeJson()
}

func (c *UserController) Logout() {
	c.DelSession("user_id")

	c.Redirect("/login", 302)
}

func (c *UserController) GetViewCompanies() {
	c.Data["companies"] = company.NewCompanyList().GetList()
}

func (c *UserController) GetViewUsers() {
	allUser, _ := user.GetList()

	companyCode := c.GetString("company")
	role := c.GetString("role")
	users := make([]*models.User, 0)
	for _, v := range allUser {
		if v.Company.Code == companyCode {
			if role == "" || v.Roles == role {
				users = append(users, v)
			}
		}
	}

	c.Data["allUsers"] = users
}

func (c *UserController) GetViewRoles() {
	var roles = []struct {
		Department string
		Role       user.RoleType
	}{
		{Department: user.ProjectManager.Desc(), Role: user.ProjectManager},
		{Department: user.BussinessMen.Desc(), Role: user.BussinessMen},
		{Department: user.TechGuy.Desc(), Role: user.TechGuy},
		{Department: user.ArtGuy.Desc(), Role: user.ArtGuy},
		{Department: user.Customer.Desc(), Role: user.Customer},
		{Department: user.Interactive.Desc(), Role: user.Interactive},
	}
	c.Data["roles"] = roles
}

func (c *UserController) GetViewQuery() {
	c.Data["company"] = c.GetString("company")
	c.Data["user_id"] = c.GetString("user_id")
	c.Data["role"] = c.GetString("role")
	c.Data["email"] = c.GetString("email")
	c.Data["mobile"] = c.GetString("mobile")
}

func (c *UserController) GetViewData() {
	c.GetViewRoles()
	c.GetViewCompanies()
	c.GetViewUsers()
	c.GetViewQuery()
}

func (c *UserController) List() {
	c.GetViewData()

	p, _ := c.GetInt("p")

	userList := user.NewUserList().IncludeDeleted(true)
	c.Data["users"] = userList.SetCondition("company__code", c.GetString("company")).SetCondition("limit", 10).SetCondition("offset", p).
		SetCondition("id", c.GetString("user_id")).SetCondition("roles", c.GetString("role")).
		SetCondition("email", c.GetString("email")).SetCondition("mobile", c.GetString("mobile")).GetList()

	c.Data["count"] = userList.GetCount()
	c.Data["paginator"] = utils.NewPaginator(c.Ctx.Request, int(10), userList.GetCount())

	c.TplNames = "users/list.tpl"
}

func (c *UserController) Create() {
	c.GetViewData()

	c.TplNames = "users/create.tpl"
}

func (c *UserController) PostCreate() {
	roles := []user.RoleType{user.RoleType(c.GetString("role"))}
	companyId, _ := c.GetInt("company")
	err := user.NewCreation(c.GetString("name"), c.GetString("email"), "123456", companyId, roles).Do()

	if err != nil {
		c.SetJsonResponse("error", err.Error())
	}

	c.GetJsonResponse().ServeJson()
}

func (c *UserController) Edit() {
	c.GetViewData()

	uid := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(uid)
	user, err := user.NewUserWithId(uint(id))

	if err != nil {
		c.Data["error"] = "用户不存在"
	}

	c.Data["user"] = user.GetUser()

	c.TplNames = "users/edit.tpl"
}

func (c *UserController) PostEdit() {
	uid, _ := c.GetInt("id")
	companyId, _ := c.GetInt("company")
	user, err := user.NewUserWithId(uint(uid))
	saveUser := user.GetUser()
	if err != nil {
		c.SetJsonResponse("error", err.Error())
	} else {
		saveUser.Company = &models.Company{Id: uint(companyId)}
		saveUser.Roles = c.GetString("role")
		saveUser.Name = c.GetString("name")
		saveUser.Email = c.GetString("email")
		saveUser.Mobile = c.GetString("mobile")
		err = saveUser.Update("company_id", "Roles", "name", "email", "mobile")
	}

	if err != nil {
		c.SetJsonResponse("error", err.Error())
	}

	c.GetJsonResponse().ServeJson()
}

func (c *UserController) Delete() {
	uid := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(uid)
	user, err := user.NewUserWithId(uint(id))
	if err != nil {
		c.SetJsonResponse("error", err.Error())
	} else {
		saveUser := user.GetUser()
		saveUser.Deleted = time.Now()
		err = saveUser.Update("deleted")
	}

	if err != nil {
		c.SetJsonResponse("error", err.Error())
	}

	c.GetJsonResponse().ServeJson()
}

func (c *UserController) Restore() {
	uid := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(uid)
	user, err := user.NewUserWithId(uint(id))
	if err != nil {
		c.SetJsonResponse("error", err.Error())
	} else {
		saveUser := user.GetUser()
		var t time.Time
		saveUser.Deleted = t
		err = saveUser.Update("deleted")
	}

	if err != nil {
		c.SetJsonResponse("error", err.Error())
	}

	c.GetJsonResponse().ServeJson()
}
