package helpers

import "strings"

var mapLabels map[string]string = map[string]string{
	"/":                       "首页",
	"/job/create":             "作业登记",
	"/job/progress":           "作业进程",
	"/job/complaint/view":     "投诉进程",
	"/project/create":         "项目登记",
	"/project/list":           "项目进程",
	"/project/job/list":       "作业总单",
	"/produce/complaint/view": "客户投诉",
}

func GetPageLabel(strUrl string) string {
	if strUrl == "" {
		return "首页"
	} else if strings.Contains(strUrl, "/job/view") {
		return "作业详情"
	} else if strings.Contains(strUrl, "/job/edit") {
		return "作业编辑"
	} else if strings.Contains(strUrl, "/job/complaint/new") {
		return "投诉登记"
	} else if strings.Contains(strUrl, "/project/job/valid") {
		return "作业审核"
	} else if strings.Contains(strUrl, "/produce/job/claim") {
		return "作业认领"
	} else if strings.Contains(strUrl, "/produce/job/submit") {
		return "提交作业"
	} else if strings.Contains(strUrl, "project/edit") {
		return "项目进程"
	}

	return mapLabels[strUrl]
}
