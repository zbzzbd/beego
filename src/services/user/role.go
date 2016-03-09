package user

import "strings"

type RoleType string

const (
	Admin          RoleType = "admin"
	ProjectManager RoleType = "projectManager"
	BussinessMen   RoleType = "bussinessman"
	ArtGuy         RoleType = "artGuy"
	TechGuy        RoleType = "techGuy"
	Customer       RoleType = "customer"
	Interactive    RoleType = "interactive"
)

var roleTypeMap map[RoleType]string = make(map[RoleType]string)

func init() {
	roleTypeMap[Admin] = "管理员"
	roleTypeMap[ProjectManager] = "项目管理员"
	roleTypeMap[BussinessMen] = "业务"
	roleTypeMap[ArtGuy] = "美术"
	roleTypeMap[TechGuy] = "技术"
	roleTypeMap[Customer] = "客服"
	roleTypeMap[Interactive] = "互动"
}

func (roleType RoleType) Desc() string {
	return roleTypeMap[roleType]
}

type Role struct {
	roles []RoleType
}

func NewRole(roles []string) *Role {
	r := new(Role)
	for _, role := range roles {
		r.roles = append(r.roles, RoleType(role))
	}
	return r
}

func NewRoleWithType(roles []RoleType) *Role {
	r := new(Role)
	r.roles = roles
	return r
}

func (r *Role) String() string {
	roleStr := []string{}
	if len(r.roles) > 0 {
		for _, role := range r.roles {
			roleStr = append(roleStr, string(role))
		}
		return strings.Join(roleStr, ",")
	}
	return string(r.roles[0])
}

func (r *Role) HasRole(role string) bool {
	for _, r := range r.roles {
		if RoleType(role) == r {
			return true
		}
	}
	return false
}

func (r *Role) IsValid() bool {
	return true
}
