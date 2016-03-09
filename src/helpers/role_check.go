package helpers

import "services/user"

const (
	Role_Refuse   = 0
	Role_Ok       = 1
	Role_ReadOnly = 2
)

type RoleCheck struct {
	Url  string
	Role *user.Role
}

func NewRoleCheck(Url string, Role *user.Role) *RoleCheck {
	roleCheck := RoleCheck{
		Url:  Url,
		Role: Role,
	}

	return &roleCheck
}

func (this *RoleCheck) Check() (retRole int8) {
	if this.Role.HasRole(string(user.ProjectManager)) {
		retRole = this.CheckProjectManager()
	} else if this.Role.HasRole(string(user.BussinessMen)) {
		retRole = this.CheckBussinessMen()
	} else if this.Role.HasRole(string(user.ArtGuy)) {
		retRole = this.CheckArtGuy()
	} else if this.Role.HasRole(string(user.TechGuy)) {
		retRole = this.CheckTechGuy()
	} else if this.Role.HasRole(string(user.Customer)) {
		retRole = this.CheckCustomer()
	} else if this.Role.HasRole(string(user.Interactive)) {
		retRole = this.CheckInteractive()
	}

	return
}

func (this *RoleCheck) CheckProjectManager() int8 {
	refuseUrls := []string{
		"/produce/job/submit",
		"/produce/complaint/view",
	}
	for _, v := range refuseUrls {
		if v == this.Url {
			return Role_Refuse
		}
	}

	readOnlyUrls := []string{
		"/produce/job/claim",
		"/job/progress",
		"/job/complaint/new",
		"/job/complaint/view",
		"/project/job/list",
	}
	for _, v := range readOnlyUrls {
		if v == this.Url {
			return Role_ReadOnly
		}
	}

	return Role_Ok
}

func (this *RoleCheck) CheckInteractive() int8 {
	refuseUrls := []string{
		"/project/create",
	}
	for _, v := range refuseUrls {
		if v == this.Url {
			return Role_Refuse
		}
	}

	readOnlyUrls := []string{
		//"/produce/job/claim",
		"/project/list",
		"/project/job/valid",
		"/project/job/list",
	}
	for _, v := range readOnlyUrls {
		if v == this.Url {
			return Role_ReadOnly
		}
	}

	return Role_Ok
}

func (this *RoleCheck) CheckCustomer() int8 {
	refuseUrls := []string{
		//"/produce/job/submit",
		"/project/create",
		//"/produce/complaint/view",
	}
	for _, v := range refuseUrls {
		if v == this.Url {
			return Role_Refuse
		}
	}

	return Role_Ok
}

func (this *RoleCheck) CheckBussinessMen() int8 {
	refuseUrls := []string{
		"/produce/job/submit",
		"/project/create",
		"/produce/complaint/view",
	}
	for _, v := range refuseUrls {
		if v == this.Url {
			return Role_Refuse
		}
	}

	readOnlyUrls := []string{
		"/produce/job/claim",
		"/project/list",
		"/project/job/valid",
		"/project/job/list",
	}
	for _, v := range readOnlyUrls {
		if v == this.Url {
			return Role_ReadOnly
		}
	}

	return Role_Ok
}

func (this *RoleCheck) CheckArtGuy() int8 {
	refuseUrls := []string{
		"/job/create",
		"/job/progress",
		"/job/complaint/new",
		"/job/complaint/view",
		"/project/create",
	}
	for _, v := range refuseUrls {
		if v == this.Url {
			return Role_Refuse
		}
	}

	readOnlyUrls := []string{
		"/project/job/valid",
		"/project/job/list",
	}
	for _, v := range readOnlyUrls {
		if v == this.Url {
			return Role_ReadOnly
		}
	}

	return Role_Ok
}

func (this *RoleCheck) CheckTechGuy() int8 {
	refuseUrls := []string{
		"/job/create",
		"/job/progress",
		"/job/complaint/new",
		"/job/complaint/view",
		"/project/create",
	}
	for _, v := range refuseUrls {
		if v == this.Url {
			return Role_Refuse
		}
	}

	readOnlyUrls := []string{
		"/project/job/list",
		"/project/job/valid",
	}
	for _, v := range readOnlyUrls {
		if v == this.Url {
			return Role_ReadOnly
		}
	}

	return Role_Ok
}
