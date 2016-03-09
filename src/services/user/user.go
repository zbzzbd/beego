package user

import (
	"errors"
	"models"
	"strings"
	"time"
	"utils"

	"github.com/astaxie/beego/orm"
)

type User struct {
	user     *models.User
	email    *Email
	password *Password
}

func NewUser() *User {
	return &User{
		user: &models.User{},
	}
}

func NewUserWithId(id uint) (u *User, err error) {
	u = &User{
		user: &models.User{Id: id},
	}
	err = models.GetDB().QueryTable("users").Filter("id", id).RelatedSel().One(u.user)
	return
}

func (u *User) Login(email string, password string) (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("user.User.Login : error :%s", err.Error())
		}
		utils.GetLog().Debug("user.User.Login : user = %s", utils.Sdump(u))
	}()

	u.email = NewEmail(email)
	u.password = NewPassword(password)

	if err = u.validLogin(); err != nil {
		return err
	}

	if err = u.afterLogin(); err != nil {
		return err
	}

	return nil
}

func (u *User) validLogin() (err error) {
	err = models.GetDB().QueryTable("users").Filter("email", u.email.EmailAddress()).RelatedSel().One(u.user)
	if err != nil || err == orm.ErrNoRows {
		return errors.New("用户不存在或密码错误")
	}

	if !u.validPassword() {
		return errors.New("用户不存在或者密码错误")
	}

	if u.user.IsDeleted() {
		return errors.New("用户不存在或者密码错误")
	}
	return nil
}

func (u *User) afterLogin() (err error) {
	if err = u.updateLoginTime(); err != nil {
		return
	}

	if err = u.addLoginRecord(); err != nil {
		return
	}
	return
}

func (u *User) updateLoginTime() (err error) {
	u.user.LastLoginTime = time.Now()
	return u.user.Update("last_login_time")
}

func (u *User) addLoginRecord() (err error) {
	ul := &models.UserLogin{
		User: u.user,
	}
	return ul.Insert()
}

func (u *User) validPassword() bool {
	u.password.SetSalt(u.user.Salt)
	return u.password.IsEncryptedSame(u.user.Password)
}

func (u *User) GetUser() *models.User {
	return u.user
}

func (u *User) GetRole() *Role {
	roles := []string{}
	if strings.Contains(u.user.Roles, ",") {
		roles = strings.Split(u.user.Roles, ",")
	} else {
		roles = []string{u.user.Roles}
	}
	return NewRole(roles)
}
