package models

import "utils"

func IsValueExists(table string, column string, value interface{}) bool {
	utils.GetLog().Debug("models.IsValueExists : table: %s, columns: %s, value: %v", table, column, value)
	return GetDB().QueryTable(table).Filter(column, value).Exist()
}
