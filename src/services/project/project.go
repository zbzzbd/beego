package project

import (
	"fmt"
	"models"
	"services/notify"
	"strconv"
	"utils"

	"time"

	"github.com/astaxie/beego/orm"
)

type Project struct {
	project          *models.Project
	isFetchedProject bool
}

func NewProject() *Project {
	return &Project{
		project: &models.Project{},
	}
}
func (p *Project) GetModelProject() *models.Project {
	return p.project
}

func (p *Project) UpdateProjectStatus(filter map[string]interface{}) (num int64, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.project.UpdateProjectStatus: error : %s", err.Error())
		} else {
			utils.GetLog().Error("services.project updateProjectStatus: debug: filter=%s,id=%v", utils.Sdump(filter))
		}
	}()
	q := models.GetDB().QueryTable(models.TABLE_NAME_PROJECT)
	for k, v := range filter {
		if k == "project_id" {
			num, err = q.Filter("Id", v).Update(orm.Params{"del_status": 1})
		}
	}
	return
}

func (p *Project) GetOne(id string) (*Project, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("id出错:id= %s", id)
	}

	if p.isFetchedProject {
		return p, nil
	}

	err = models.GetDB().QueryTable("projects").Filter("Id", uint(idInt)).RelatedSel("BussinessUser", "Progress", "ArtUser", "TechUser", "Registrant").One(p.project)
	if err != nil {
		return nil, fmt.Errorf("找不到项目:id = %s", id)
	}
	p.isFetchedProject = true

	return p, nil
}

//根据 ID 获取user
func (p *Project) GetUserByID(id uint) (*models.User, error) {
	user := &models.User{}
	err := models.GetDB().QueryTable(models.TABLE_NAME_USER).Filter("id", id).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//根据 项目进程id ，获取process
func (p *Project) GetProgressByID(id uint) (*models.Progress, error) {
	progress := &models.Progress{}
	err := models.GetDB().QueryTable(models.TABLE_NAME_PROGRESS).Filter("id", id).One(progress)
	if err != nil {
		return nil, err
	}
	return progress, nil
}

//获取导出名称

func (p *Project) GetOutName(filter map[string]interface{}) string {
	outName := "项目进程"
	for k, v := range filter {
		fmt.Println(k)
		if k == "id" {
			p, _ := p.GetOne(fmt.Sprint(v))
			outName = outName + "-" + p.GetModelProject().Name
		} else if k == "bussiness_user_id" {
			buserid, _ := strconv.Atoi(fmt.Sprint(v))
			user, _ := p.GetUserByID(uint(buserid))
			outName = outName + "-" + user.Name
		} else if k == "progress" {
			progid, _ := strconv.Atoi(fmt.Sprint(v))
			progress, _ := p.GetProgressByID(uint(progid))
			outName = outName + "-" + progress.Name
		} else {
			outName = outName + "-" + fmt.Sprint(v)
		}
	}
	return outName + ".csv"
}

//验证输入数据是否规范
func (p *Project) CheckData() error {
	return nil
}

func (p *Project) Create(data map[string]interface{}) (err error, projectId int64) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("project.Project.Create : error : %s", err.Error())
		} else {
			utils.GetLog().Debug("project.Project.Create : debug : %s", utils.Sdump(p))
		}

	}()

	Orm := orm.NewOrm()
	err = Orm.Begin()
	if err != nil {
		return
	}

	p.project.Name = data["name"].(string)

	priority, _ := strconv.Atoi(data["priority"].(string))
	p.project.Priority = uint(priority)

	progress, _ := strconv.Atoi(data["progress"].(string))
	p.project.Progress = &models.Progress{Id: uint(progress)}

	scale, _ := strconv.Atoi(data["scale"].(string))
	p.project.Scale = uint(scale)

	p.project.ClientName = data["clientName"].(string)
	p.project.ContractNo = data["contractNo"].(string)
	p.project.ServiceItem = data["serviceItem"].(string)

	p.project.GameDate, err = utils.FormatTime(data["gameDate"].(string))

	p.project.RegStartDate, err = utils.FormatTime(data["regStartDate"].(string))
	p.project.RegCloseDate, err = utils.FormatTime(data["regCloseDate"].(string))
	bussinessUserId, _ := strconv.Atoi(data["bussinessUser"].(string))
	p.project.BussinessUser = &models.User{Id: uint(bussinessUserId)}

	//registrant, _ := strconv.Atoi(data["registrant"].(string))
	p.project.Registrant = &models.User{Id: data["registrant"].(uint)}

	artUserId, _ := strconv.Atoi(data["artUser"].(string))
	p.project.ArtUser = &models.User{Id: uint(artUserId)}
	techUserId, _ := strconv.Atoi(data["techUser"].(string))
	p.project.TechUser = &models.User{Id: uint(techUserId)}

	p.project.Started, err = utils.FormatTime(data["started"].(string))

	p.project.Created = time.Now().Local()

	err = p.CheckData()
	if err != nil {
		return
	}

	projectId, err = Orm.Insert(p.project)
	if err != nil {
		err1 := Orm.Rollback()
		if err1 != nil {
			return err1, projectId
		}
		return err, projectId
	}

	var projectLog models.ProjectLog
	projectLog.Project = &models.Project{Id: uint(projectId)}
	projectLog.User = &models.User{Id: data["registrant"].(uint)}
	projectLog.Operation = models.CREATE

	_, err = Orm.Insert(&projectLog)
	if err != nil {
		fmt.Println(err)
		err1 := Orm.Rollback()
		if err1 != nil {
			return err1, projectId
		}
		return err, projectId
	}

	err = Orm.Commit()
	if err != nil {
		return
	}

	//emial notifier
	if err = p.sendEmailNotifier(); err != nil {
		return
	}

	return
}

func (p *Project) BussinessUser() (*models.User, error) {
	return p.project.BussinessUser, nil
}

func (p Project) sendEmailNotifier() (err error) {
	n := notify.NewEmail(p)
	err = n.Send()
	return
}

//############################################################
// implements EmailNotifier
//############################################################
func (p Project) GetTo() []string {
	if !p.isFetchedProject {
		p.GetOne(fmt.Sprintf("%d", p.project.Id))
	}
	return []string{
		(&p).project.BussinessUser.Email,
		(&p).project.TechUser.Email,
		(&p).project.ArtUser.Email,
	}
}

func (p Project) GetTemplateName() notify.EmailTemplate {
	return notify.PM_Notify
}

func (p Project) GetVariables() []map[string]interface{} {
	if !p.isFetchedProject {
		p.GetOne(fmt.Sprintf("%d", p.project.Id))
	}

	data := []map[string]interface{}{}

	bussinessUserContent := map[string]interface{}{
		"name":            (&p).project.BussinessUser.Name,
		"job_name":        "项目创建",
		"job_link":        fmt.Sprintf("%s", utils.BaseUrl()),
		"job_status_desc": "项目创建",
		"job_created_at":  p.project.Created.Format(utils.YMDHIS),
		"datetime":        time.Now().Format(utils.YMDHIS),
	}
	data = append(data, bussinessUserContent)

	techUserContent := map[string]interface{}{
		"name":            (&p).project.TechUser.Name,
		"job_name":        "项目创建",
		"job_link":        fmt.Sprintf("%s", utils.BaseUrl()),
		"job_status_desc": "项目创建",
		"job_created_at":  p.project.Created.Format(utils.YMDHIS),
		"datetime":        time.Now().Format(utils.YMDHIS),
	}
	data = append(data, techUserContent)

	artUserContent := map[string]interface{}{
		"name":            (&p).project.ArtUser.Name,
		"job_name":        "项目创建",
		"job_link":        fmt.Sprintf("%s", utils.BaseUrl()),
		"job_status_desc": "项目创建",
		"job_created_at":  p.project.Created.Format(utils.YMDHIS),
		"datetime":        time.Now().Format(utils.YMDHIS),
	}
	data = append(data, artUserContent)

	return data
}

//############################################################
// implements EmailNotifier
//############################################################
