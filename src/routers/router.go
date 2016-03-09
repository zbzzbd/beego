package routers

import (
	"controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.UserController{}, "get:Logout;post:Logout")

	beego.Router("/modify_password", &controllers.UserController{}, "get:Get;post:ModifyPassword")

	beego.Router("/job", &controllers.JobController{})
	beego.Router("/job/create", &controllers.JobController{}, "get:ViewCreateJob;post:PostCreateJob")
	beego.Router("/job/edit/:id", &controllers.JobController{}, "get:ViewEditJob;post:PostEditJob")
	beego.Router("/job/view/:id", &controllers.JobController{}, "get:ViewJob")
	beego.Router("/job/delete/:id", &controllers.JobController{}, "get:DelJob")
	beego.Router("/job/recover/:id", &controllers.JobController{}, "get:RecoverJob")
	//投诉
	beego.Router("/job/complaint/create", &controllers.ComplaintController{}, "post:CreateComplain")
	beego.Router("/job/complaint/new", &controllers.ComplaintController{})
	beego.Router("/job/complaint/new/:id", &controllers.ComplaintController{}, "get:Show")
	beego.Router("/job/complaint/export", &controllers.ComplaintController{}, "get:Export")

	beego.Router("/job/complaint/del/:id", &controllers.ComplaintController{}, "get:Del")
	beego.Router("/job/complaint/view", &controllers.ComplaintController{}, "get:ViewCompliant")
	beego.Router("/produce/complaint/view", &controllers.ComplaintController{}, "get:ViewMyCompliant")
	beego.Router("/produce/complaint/reply/:id", &controllers.ReplyController{}, "get:Show;post:CreateReply")
	beego.Router("/search/job", &controllers.ComplaintController{}, "get:GetJob")

	beego.Router("/job/file/del", &controllers.JobController{}, "get:DelJobFile")
	beego.Router("/job/progress", &controllers.JobController{}, "get:Get")
	beego.Router("/project/job/list", &controllers.JobController{}, "get:GetJobList")
	beego.Router("/project/job/dellist", &controllers.JobController{}, "get:GetDelJobList")
	beego.Router("/project/job/export", &controllers.JobController{}, "get:Export")

	//作业审核列表与审核
	beego.Router("/project/job/valid", &controllers.JobValidationController{})
	beego.Router("/project/job/valid/:id", &controllers.JobValidationController{}, "get:Show")

	beego.Router("/project/list", &controllers.ProjectController{}, "get:ProjectList")
	beego.Router("/project/del/:id", &controllers.ProjectController{}, "get:DelProject")
	beego.Router("/project/create", &controllers.ProjectController{}, "get:ViewCreateProject;post:CreateProject")
	beego.Router("/project/detail/:id", &controllers.ProjectController{}, "get:ViewProjectDetail")
	beego.Router("/project/edit/:id", &controllers.ProjectController{}, "get:ViewEditProject;post:EditProject")
	beego.Router("/project", &controllers.ProjectController{})
	beego.Router("/project/export", &controllers.ProjectController{}, "get:Export")

	beego.Router("/produce", &controllers.ProduceController{})
	beego.Router("/produce/job/claim", &controllers.ProduceController{}, "get:ClaimList")
	beego.Router("/produce/job/assign", &controllers.JobAssignController{})
	beego.Router("/produce/job/submit", &controllers.ProduceController{}, "get:SubmitList")
	beego.Router("/produce/job/claim/:id", &controllers.JobClaimController{})
	beego.Router("/produce/job/submit/:id", &controllers.JobSubmitController{})

	beego.Router("/user/list", &controllers.UserController{}, "get:List")
	beego.Router("/user/create", &controllers.UserController{}, "get:Create;post:PostCreate")
	beego.Router("/user/edit/:id", &controllers.UserController{}, "get:Edit;post:PostEdit")
	beego.Router("/user/delete/:id", &controllers.UserController{}, "get:Delete")
	beego.Router("/user/restore/:id", &controllers.UserController{}, "get:Restore")

	beego.Router("/require/list", &controllers.RequireController{}, "get:List")
	beego.Router("/require/create", &controllers.RequireController{}, "get:Create")
	beego.Router("/require/edit", &controllers.RequireController{}, "get:Edit")
	beego.Router("/require/delete", &controllers.RequireController{}, "get:Delete")

	beego.Router("/role/error", &controllers.RoleController{})

}
