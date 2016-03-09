package project

import (
	"helpers"
	"models"

	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

var timeFormat string = "2006-01-02 15:04"

type JobsCount struct {
	ProjectId int
	JobsNum   int64
}

type ProjectList struct {
	projectList []*models.Project
	count       int64
}

func NewProjectList() *ProjectList {
	return &ProjectList{
		projectList: []*models.Project{},
	}
}
func (pl *ProjectList) Count() int64 {
	return pl.count
}

func isKeyFitForFilter(key string) bool {
	if key == "limit" || key == "offset" || key == "user" {
		return false
	}
	return true
}

func (pl *ProjectList) GetList(filter map[string]interface{}) (projects []*models.Project, err error) {

	qs := models.GetDB().QueryTable("projects").RelatedSel("BussinessUser", "Progress", "ArtUser", "TechUser", "Registrant").OrderBy("-created").Filter("del_status", 0)

	for k, v := range filter {
		if !isKeyFitForFilter(k) {
			continue
		}
		if k == "start_date" {
			start, _ := time.Parse(timeFormat, v.(string))
			qs = qs.Filter("started__gte", start)
			continue
		}
		if k == "end_date" {
			end, _ := time.Parse(timeFormat, v.(string))
			qs = qs.Filter("started__lte", end)
			continue
		}
		qs = qs.Filter(k, v)

	}
	if filter["user"] != nil {
		cond := orm.NewCondition()
		cond1 := cond.And("tech_user_id", filter["user"]).Or("art_user_id", filter["user"]).Or("registrant", filter["user"]).Or("bussiness_user_id", filter["user"])
		qs = qs.SetCond(cond1)
	}

	count, err := qs.Count()
	if err != nil {
		return nil, err
	}

	pl.count = count

	if filter["limit"] != nil {
		qs = qs.Limit(filter["limit"].(int))
	}

	if filter["offset"] != nil {
		qs = qs.Offset(filter["offset"].(int))
	}

	_, err = qs.All(&pl.projectList)
	if err != nil {

		return nil, err
	}

	//
	err = setProjectListJobsNum(pl.projectList)
	if err != nil {

		return nil, err
	}
	return pl.projectList, nil
}

func setProjectListJobsNum(projects []*models.Project) error {
	if len(projects) == 0 {
		return nil
	}
	sql := getSqlSyntaxForJobsCountByProjects(projects)

	var counts []JobsCount
	_, err := models.GetDB().Raw(sql).QueryRows(&counts)

	if err != nil {
		return err
	}

	projectJobsCount := make(map[int]int64)
	for _, count := range counts {
		projectJobsCount[count.ProjectId] = count.JobsNum
	}
	for _, project := range projects {
		project.JobsNum = uint(projectJobsCount[int(project.Id)])
	}
	return nil
}

func getSqlSyntaxForJobsCountByProjects(projects []*models.Project) string {
	var projectIds []string
	for _, p := range projects {
		projectIds = append(projectIds, strconv.Itoa(int(p.Id)))
	}
	projectIdsString := strings.Join(projectIds, ",")
	projectIdsString = "(" + projectIdsString + ")"
	return fmt.Sprintf("select project_id, count(project_id) as jobs_num from jobs where project_id IN %s group by project_id", projectIdsString)
}

func (pl *ProjectList) getProjectIds() (projectIds []uint) {
	for _, p := range pl.projectList {
		projectIds = append(projectIds, p.Id)
	}
	return
}

func (pl *ProjectList) GetAllProjectNames() ([]*models.Project, error) {
	_, err := models.GetDB().Raw("SELECT id, name FROM projects").QueryRows(&pl.projectList)
	if err != nil {
		return nil, err
	}

	return pl.projectList, nil
}

func GetExportData(projects []*models.Project) [][]string {
	records := [][]string{{"序号", "启动时间", "项目名称", "服务项目", "合同编号", "赛事规模", "客户名称", "业务担当", "比赛日期", "开通报名时间", "关闭报名时间", "进程现况", "美术单元", "技术单元", "优先级", "登记人"}}
	light := len(projects)
	for i := 0; i < light; i++ {
		project := projects[i]
		data := []string{strconv.Itoa(i + 1), helpers.TimeFormat(project.Started), project.Name, project.ServiceItem, project.ContractNo, strconv.Itoa(int(project.Scale)), project.ClientName, project.BussinessUser.Name, helpers.TimeFormat(project.GameDate),
			helpers.TimeFormat(project.RegStartDate), helpers.TimeFormat(project.RegCloseDate), project.Progress.Name, project.ArtUser.Name, project.TechUser.Name, strconv.Itoa(int(project.Priority)), project.Registrant.Name,
		}
		records = append(records, data)
	}
	return records
}
