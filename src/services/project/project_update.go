package project

import (
	"models"
	"strconv"
	"time"
	"utils"

	"github.com/astaxie/beego/orm"
)

func (p *Project) Update(data map[string]interface{}) (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("project.Project.Update : error : %s", err.Error())

		} else {
			utils.GetLog().Debug("project.Project.Update : debug : %s", utils.Sdump(p))
		}
	}()

	Orm := orm.NewOrm()
	err = Orm.Begin()
	if err != nil {
		return

	}

	var projectLogs []models.ProjectLog
	//project_id, _ := strconv.Atoi(data["id"].(string))
	//p.project.Id = uint(project_id)

	p.project.Registrant = &models.User{Id: data["registrant"].(uint)}

	if p.project.Name != data["name"].(string) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.NAME, p.project.Name, data["name"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.Name = data["name"].(string)

	}

	scale, _ := strconv.Atoi(data["scale"].(string))
	if p.project.Scale != uint(scale) {

		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.SCALE, strconv.Itoa(int(p.project.Scale)), data["scale"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.Scale = uint(scale)

	}

	priority, _ := strconv.Atoi(data["priority"].(string))

	if p.project.Priority != uint(priority) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.PRIORITY, strconv.Itoa(int(p.project.Priority)), data["priority"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.Priority = uint(priority)

	}

	if p.project.ClientName != data["clientName"].(string) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.CLIENT_NAME, p.project.ClientName, data["clientName"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.ClientName = data["clientName"].(string)

	}
	if p.project.ContractNo != data["contractNo"].(string) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.CONRACT_NO, p.project.ContractNo, data["contractNo"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.ContractNo = data["contractNo"].(string)
	}
	if p.project.ServiceItem != data["serviceItem"].(string) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.SERVICE_ITEM, p.project.ServiceItem, data["serviceItem"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.ServiceItem = data["serviceItem"].(string)
	}

	gameDate, err := utils.FormatTime(data["gameDate"].(string))
	if p.project.GameDate.Format("2006-01-02 15:04") != data["gameDate"].(string) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.GAME_DATE, p.project.GameDate.Format("2006-01-02 15:04"), data["gameDate"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.GameDate = gameDate

	}

	regStartDate, err := utils.FormatTime(data["regStartDate"].(string))
	if p.project.RegStartDate.Format("2006-01-02 15:04") != data["regStartDate"].(string) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.REG_START_DATE, p.project.RegStartDate.Format("2006-01-02 15:04"), data["regStartDate"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.RegStartDate = regStartDate

	}

	regCloseDate, err := utils.FormatTime(data["regCloseDate"].(string))
	if p.project.RegCloseDate.Format("2006-01-02 15:04") != data["regCloseDate"].(string) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.REG_CLOSE_DATA, p.project.RegCloseDate.Format("2006-01-02 15:04"), data["regCloseDate"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.RegCloseDate = regCloseDate

	}

	progress, _ := strconv.Atoi(data["progress"].(string))
	if p.project.Progress.Id != uint(progress) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.PROGRESS, strconv.Itoa(int(p.project.Progress.Id)), data["progress"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.Progress = &models.Progress{Id: uint(progress)}

	}

	bussinessUserId, _ := strconv.Atoi(data["bussinessUser"].(string))
	if p.project.BussinessUser.Id != uint(bussinessUserId) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.BUSSINESS_USER, strconv.Itoa(int(p.project.BussinessUser.Id)), data["bussinessUser"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.BussinessUser = &models.User{Id: uint(bussinessUserId)}

	}

	artUserId, _ := strconv.Atoi(data["artUser"].(string))
	if p.project.ArtUser.Id != uint(artUserId) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.ART_USER_ID, strconv.Itoa(int(p.project.ArtUser.Id)), data["artUser"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.ArtUser = &models.User{Id: uint(artUserId)}

	}

	techUserId, _ := strconv.Atoi(data["techUser"].(string))
	if p.project.TechUser.Id != uint(techUserId) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.TECH_USER_ID, strconv.Itoa(int(p.project.TechUser.Id)), data["techUser"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.TechUser = &models.User{Id: uint(techUserId)}

	}

	started, err := utils.FormatTime(data["started"].(string))
	if p.project.Started.Format("2006-01-02 15:04") != data["started"].(string) {
		projectLog := models.GetProjectLogForEdit(p.project.Id, data["registrant"].(uint)).SetColumn(models.STARTED, p.project.Started.Format("2006-01-02 15:04"), data["started"].(string))
		projectLogs = append(projectLogs, *projectLog)
		p.project.Started = started

	}

	p.project.Updated = time.Now().Local()
	utils.GetLog().Debug("p.project= %s", utils.Sdump(p.project))
	err = p.CheckData()
	if err != nil {
		return

	}
	_, err = Orm.Update(p.project)
	if err != nil {
		err1 := Orm.Rollback()
		if err1 != nil {
			return err1

		}
		return err

	}

	_, err = Orm.InsertMulti(1, projectLogs)
	if err != nil {
		err1 := Orm.Rollback()
		if err1 != nil {
			return err1

		}
		return err

	}

	err = Orm.Commit()
	if err != nil {
		return

	}

	return

}
