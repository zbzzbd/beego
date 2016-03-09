package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"models"
	userSrv "services/user"
	"strings"
)

func TimeFormat(start_time interface{}) string {
	if start_time == nil {
		return ""
	}

	var t time.Time
	var err error
	switch val := start_time.(type) {
	case string:
		t, err = time.Parse(time.RFC3339, val)
		if err != nil {
			return ""
		}
	case time.Time:
		t = val
	default:
		return ""
	}

	strTime := t.Format("2006-01-02 15:04")
	if strTime == "0001-01-01 00:00" {
		strTime = ""
	}

	return strTime
}

func ToFloat64(item interface{}) (num float64) {
	switch val := item.(type) {
	case int8:
		num = float64(val)
	case uint8:
		num = float64(val)
	case int32:
		num = float64(val)
	case uint32:
		num = float64(val)
	case uint:
		num = float64(val)
	case int64:
		num = float64(val)
	case uint64:
		num = float64(val)
	case float32:
		num = float64(val)
	case int:
		num = float64(val)
	case float64:
		num = val
	case string:
		num, _ = strconv.ParseFloat(val, 64)
	case json.Number:
		num, _ = val.Float64()
	}

	return
}

func Gt(num1, num2 interface{}) bool {
	return ToFloat64(num1) > ToFloat64(num2)
}

func Strcat(s ...string) string {
	return strings.Join(s, "")
}

func ShowMenu(strUrl string, u ...interface{}) string {
	var show string = "hidden"

	if u == nil || u[0] == nil {
		return show
	}

	role := u[0].(*models.User).Roles

	if strings.Contains(role, string(userSrv.ProjectManager)) {
		if strings.Index(strUrl, "/project") == 0 {
			show = "show"
		}
	}

	if strings.Contains(role, string(userSrv.BussinessMen)) || strings.Contains(role, string(userSrv.Interactive)) {
		if strings.Index(strUrl, "/job") == 0 {
			show = "show"
		}
	}

	if strings.Contains(role, string(userSrv.ArtGuy)) || strings.Contains(role, string(userSrv.TechGuy)) || strings.Contains(role, string(userSrv.Interactive)) {
		if strings.Index(strUrl, "/produce") == 0 {
			show = "show"
		}
	}
	if strings.Contains(role, string(userSrv.Customer)) {
		if strings.Index(strUrl, "/customer") == 0 {
			show = "show"
		}

	}

	return show
}

func AddInt(v ...interface{}) (sum int) {
	for _, item := range v {
		switch val := item.(type) {
		case int8:
			sum += int(val)
		case uint8:
			sum += int(val)
		case int32:
			sum += int(val)
		case uint32:
			sum += int(val)
		case uint:
			sum += int(val)
		case int64:
			sum += int(val)
		case uint64:
			sum += int(val)
		case float32:
			sum += int(val)
		case float64:
			sum += int(val)
		case int:
			sum += val
		case string:
			num, err := strconv.Atoi(val)
			if err == nil {
				sum += num
			}
		case json.Number:
			num, err := val.Int64()
			if err == nil {
				sum += int(num)
			}
		}
	}

	return
}

func GetStatusDesc(job models.Job) string {
	if job.SubmitStatus == models.Job_Submit_Acceptance_OK {
		return "验收通过"
	} else if job.SubmitStatus == models.Job_Submit_Acceptance_Refuse {
		return "验收未通过"
	} else if job.SubmitStatus == models.Job_Submit_Wait_Acceptance {
		return "待验收"
	}

	if job.ClaimStatus == models.Job_Claim_Ok {
		return "制作中"
	} else if job.ClaimStatus == models.Job_Claim_Refuse {
		return "认领被拒绝"
	}

	if job.ValidStatus == models.Job_Valid_Cancel {
		return "任务取消"
	} else if job.ValidStatus == models.Job_Valid_Refuse {
		return "审核未通过，待修改"
	} else if job.ValidStatus == models.Job_Valid_OK {
		return "待认领制作"
	}

	return "待审核"
}

func GetStatusColor(job models.Job) string {
	if job.SubmitStatus == models.Job_Submit_Acceptance_OK {
		return "green"
	} else if job.SubmitStatus == models.Job_Submit_Acceptance_Refuse {
		return "yellow"
	} else if job.SubmitStatus == models.Job_Submit_Wait_Acceptance {
		return "brown"
	}

	if job.ClaimStatus == models.Job_Claim_Ok {
		return "blue"
	} else if job.ClaimStatus == models.Job_Claim_Refuse {
		return "orange"
	}

	if job.ValidStatus == models.Job_Valid_Cancel {
		return "red"
	} else if job.ValidStatus == models.Job_Valid_Refuse {
		return "pink"
	} else if job.ValidStatus == models.Job_Valid_OK {
		return "violet"
	}

	return "teal"
}

func GetTimeDiff(startTime time.Time, endTime time.Time) string {
	if startTime.IsZero() {
		return ""
	}
	duration := startTime.Sub(endTime)
	if duration < 0 {
		return ""
	}
	return fmt.Sprintf("%.0f", duration.Minutes())
}

func TimeIsZero(startTime time.Time) bool {
	return startTime.IsZero()
}

func GetRoleDesc(role string) string {
	return userSrv.RoleType(role).Desc()
}
