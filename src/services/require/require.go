package require

import (
	"models"

	"github.com/astaxie/beego/orm"
)

type RequireList struct {
	requires []*models.Require
}

func NewRequireList() *RequireList {
	requireList := RequireList{}
	return &requireList
}

func (requireList *RequireList) GetRequireNames() []string {
	requireList.getList()

	requireNames := make([]string, 0)
	for _, v := range requireList.requires {
		requireNames = append(requireNames, v.Name)
	}

	return requireNames
}

func (requireList *RequireList) getList() (err error) {
	_, err = models.GetDB().QueryTable(models.TABLE_NAME_REQUIRE).All(&requireList.requires)

	return
}

func (requireList *RequireList) Add(name string) (int64, error) {
	var require models.Require = models.Require{Name: name}
	return models.GetDB().Insert(&require)
}

func (requireList *RequireList) Edit(name, newName string) (err error) {
	o := models.GetDB()
	err = o.Begin()
	//事务处理过程
	_, SomeError := models.GetDB().QueryTable(models.TABLE_NAME_REQUIRE).Filter("name", name).Update(orm.Params{"name": newName})
	_, SomeError = models.GetDB().QueryTable(models.TABLE_NAME_JOB).Filter("type", name).Update(orm.Params{"type": newName})
	_, SomeError = models.GetDB().QueryTable(models.TABLE_NAME_JOBHISTORY).Filter("type", name).Update(orm.Params{"type": newName})
	if SomeError != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return

}

func (requireList *RequireList) Delete(name string) (int64, error) {
	return models.GetDB().QueryTable(models.TABLE_NAME_REQUIRE).Filter("name", name).Delete()
}
