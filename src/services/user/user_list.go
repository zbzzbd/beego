package user

import (
	"models"
	"strings"
	"utils"
)

func GetList() (users []*models.User, err error) {
	return NewUserList().GetList(), nil
}

type UserList struct {
	users  []*models.User
	filter map[string]interface{}
	count  int64
	isDo   bool
	includeDel bool
}

func NewUserList() *UserList {
	userList := UserList{}
	userList.filter = make(map[string]interface{})

	return &userList
}

func (userList *UserList) do() *UserList {
	if userList.isDo {
		return  userList
	}
	userList.isDo = true

	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.user.GetList : error : %s", err.Error())
		} else {
			utils.GetLog().Debug("services.user.GetList : debug : users=%v", utils.Sdump(userList.users))
		}
	}()

	q := models.GetDB().QueryTable("users").RelatedSel("Company")
	if !userList.includeDel {
		q = q.Filter("deleted__isnull", true)
	}

	for k, v := range userList.filter {
		if k != "limit" && k != "offset" && v != "" {
			q = q.Filter(k, v)
		} else if k == "limit" {
			limit := v.(int)
			if limit < 0 {
				limit = 10
			}
			q = q.Limit(limit)
		} else if k == "offset" {
			limit := userList.filter["limit"].(int)
			if limit < 0 {
				limit = 10
			}

			offset := v.(int)
			if offset > 0 {
				offset = (offset - 1) * limit
			}
			q = q.Offset(offset)
		}
	}

	userList.count, _ = q.Count()
	_, err = q.All(&userList.users)
	if err != nil {
		return userList
	}

	return userList
}

func (userList *UserList) IncludeDeleted(includeDel bool) *UserList {
	userList.includeDel = true

	return userList
}

func (userList *UserList) GetList() []*models.User {
	userList.do()

	return userList.users
}

func (userList *UserList) GetCount() int64 {
	userList.do()

	return userList.count
}

func (userList *UserList) SetCondition(k string, v interface{}) *UserList {
	userList.filter[k] = v
	userList.isDo = false

	return userList
}

func (userList *UserList) GetRoleList(role RoleType) []*models.User {
	userList.do()

	var newUsers []*models.User
	for _, v := range userList.users {
		if strings.Contains(v.Roles, string(role)) {
			newUsers = append(newUsers, v)
		}
	}

	return newUsers
}

func (ul *UserList) GetRoleListExcept(role RoleType, userIds []uint) []*models.User {
	users := ul.GetRoleList(role)
	var newUsers []*models.User
	for _, user := range users {
		if !ul.contains(userIds, user.Id) {
			newUsers = append(newUsers, user)
		}
	}
	return newUsers
}

func (ul *UserList) contains(container []uint, target uint) bool {
	for _, item := range container {
		if target == item {
			return true
		}
	}
	return false
}
