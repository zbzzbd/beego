package user

import (
	"errors"
	"models"
	"time"
	"utils"
)

var (
	ErrEmail            = errors.New("邮箱格式错误")
	ErrEmailExists      = errors.New("邮箱已经被注册")
	ErrCompanyNotExists = errors.New("公司不存在")
	ErrRole             = errors.New("角色出错")
)

type Creation struct {
	name      string
	email     *Email
	password  *Password
	companyId int
	roles     *Role
}

func NewCreation(name, email, password string, companyId int, roles []RoleType) *Creation {
	c := &Creation{
		name:      name,
		email:     NewEmail(email),
		password:  NewPassword(password),
		companyId: companyId,
		roles:     NewRoleWithType(roles),
	}
	return c
}

func (c *Creation) valid() (err error) {
	if !c.email.IsValid() {
		err = ErrEmail
		return
	}
	if c.email.IsExists() {
		err = ErrEmailExists
		return
	}

	if !c.isCompanyExists() {
		err = ErrCompanyNotExists
		return
	}

	if !c.roles.IsValid() {
		err = ErrRole
		return
	}
	return
}

func (c *Creation) isCompanyExists() bool {
	return models.IsValueExists("companies", "id", c.companyId)
}

func (c *Creation) String() string {
	return utils.Sdump(c)
}

func (c *Creation) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("user.Creatin.Do : err = %v , obj = %s ", err, c)
		}
	}()

	if err = c.valid(); err != nil {
		return
	}

	c.password.GenSalt()
	c.password.GenPwd()

	u := models.User{
		Name:          c.name,
		Email:         c.email.EmailAddress(),
		Password:      c.password.Encryped(),
		Salt:          c.password.Salt(),
		Company:       &models.Company{Id: uint(c.companyId)},
		Roles:         c.roles.String(),
		LastLoginTime: time.Now(),
	}

	models.GetDB().Insert(&u)

	return
}
