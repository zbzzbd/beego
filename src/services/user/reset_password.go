package user

import (
	"errors"
	"models"
	"utils"
)

var (
	ErrOldPassword       = errors.New("原密码不正确")
	ErrPasswordNotSame   = errors.New("新密码与确认密码输入不一致")
	ErrUpdateNewPassword = errors.New("更新新密码失败")
)

// PasswordModify 用于重新设置用户密码
type PasswordModify struct {
	user          *models.User
	oldPwd        *Password
	newPwd        *Password
	newPwdConfirm string
}

func NewPasswordModify(currentUser *User, oldPassword, newPassword, newPasswordConfirm string) *PasswordModify {
	user := currentUser.GetUser()
	return &PasswordModify{
		user:          user,
		oldPwd:        NewPassword(oldPassword),
		newPwd:        NewPassword(newPassword),
		newPwdConfirm: newPasswordConfirm,
	}
}

func (pr *PasswordModify) oldPasswordValid() bool {
	pr.oldPwd.SetSalt(pr.user.Salt)
	return pr.oldPwd.IsEncryptedSame(pr.user.Password)
}

func (pr *PasswordModify) newPasswordConfirm() bool {
	return pr.newPwd.IsConfirmSame(pr.newPwdConfirm)
}

func (pr *PasswordModify) newPasswordValid() error {
	return pr.newPwd.Valid()
}

func (pr *PasswordModify) updatePassword() error {
	pr.user.Salt = pr.newPwd.Salt()
	pr.user.Password = pr.newPwd.Encryped()
	affectedRowsNum, err := models.GetDB().Update(pr.user, "password", "salt")
	if affectedRowsNum < 1 || err != nil {
		return ErrUpdateNewPassword
	}
	return nil
}

func (pr *PasswordModify) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("accountSrv.PasswordReset.Do : %s", err.Error())
		} else {
			utils.GetLog().Debug("accountSrv.PasswordReset.Do : pr %v", pr)
		}
	}()

	if !pr.newPasswordConfirm() {
		return ErrPasswordNotSame
	}

	if err = pr.newPasswordValid(); err != nil {
		return err
	}

	if !pr.oldPasswordValid() {
		return ErrOldPassword
	}

	if err = pr.updatePassword(); err != nil {
		return err
	}
	return nil
}
