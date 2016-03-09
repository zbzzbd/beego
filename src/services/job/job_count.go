package job

import (
	"models"
	"utils"
)

func GetCount(filter map[string]interface{}) (cnt int64, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.job.GetCount : error : %s", err.Error())
		} else {
			utils.GetLog().Debug("services.job.GetCount : debug : filter=%s, count=%v", utils.Sdump(filter), cnt)
		}
	}()

	q := models.GetDB().QueryTable(models.TABLE_NAME_JOB)
	for k, v := range filter {
		q = q.Filter(k, v)
	}
	cnt, err = q.Count()

	return
}

func GetCreateCount(createUserId, employeeId string) (cnt int64, err error) {
	filter := make(map[string]interface{})
	if createUserId != "" {
		filter["create_user_id"] = createUserId
	}

	if employeeId != "" {
		filter["employee_id"] = employeeId
	}

	cnt, err = GetCount(filter)

	return
}

func GetFinishCount(createUserId, employeeId string) (cnt int64, err error) {
	filter := make(map[string]interface{})
	filter["submit_status"] = models.Job_Submit_Acceptance_OK
	if createUserId != "" {
		filter["create_user_id"] = createUserId
	}

	if employeeId != "" {
		filter["employee_id"] = employeeId
	}

	cnt, err = GetCount(filter)

	return
}

func GetModifyCount(createUserId, employeeId string) (cnt int64, err error) {
	filter := make(map[string]interface{})
	filter["valid_status"] = models.Job_Valid_None
	if createUserId != "" {
		filter["create_user_id"] = createUserId
	}

	if employeeId != "" {
		filter["employee_id"] = employeeId
	}

	cnt, err = GetCount(filter)

	return
}

func GetDoingCount(createUserId, employeeId string) (cnt int64, err error) {
	filter := make(map[string]interface{})
	filter["valid_status"] = models.Job_Valid_OK
	filter["claim_status"] = models.Job_Claim_Ok
	filter["submit_status__lt"] = models.Job_Submit_Acceptance_OK
	if createUserId != "" {
		filter["create_user_id"] = createUserId
	}

	if employeeId != "" {
		filter["employee_id"] = employeeId
	}

	cnt, err = GetCount(filter)

	return
}

func GetCancelCount(createUserId, employeeId string) (cnt int64, err error) {
	filter := make(map[string]interface{})
	filter["valid_status"] = models.Job_Valid_Cancel
	if createUserId != "" {
		filter["create_user_id"] = createUserId
	}

	if employeeId != "" {
		filter["employee_id"] = employeeId
	}

	cnt, err = GetCount(filter)

	return
}

// 待审核
// uid - 业务单元，创建作业的人
func GetWaitValidCount(createUserId, employeeId string) (cnt int64, err error) {
	filter := make(map[string]interface{})
	filter["valid_status"] = models.Job_Valid_None
	if createUserId != "" {
		filter["create_user_id"] = createUserId
	}

	if employeeId != "" {
		filter["employee_id"] = employeeId
	}

	cnt, err = GetCount(filter)

	return
}

// 待修改
// uid - 业务单元，创建作业的人
func GetValidRefuseCount(createUserId, employeeId string) (cnt int64, err error) {
	filter := make(map[string]interface{})
	filter["valid_status"] = models.Job_Valid_Refuse
	if createUserId != "" {
		filter["create_user_id"] = createUserId
	}

	if employeeId != "" {
		filter["employee_id"] = employeeId
	}

	cnt, err = GetCount(filter)

	return
}

// 待认领
func GetWaitClaimCount(createUserId, employeeId string) (cnt int64, err error) {
	filter := make(map[string]interface{})
	filter["claim_status"] = models.Job_Claim_None
	filter["valid_status"] = models.Job_Valid_OK
	if createUserId != "" {
		filter["create_user_id"] = createUserId
	}

	if employeeId != "" {
		filter["employee_id"] = employeeId
	}

	cnt, err = GetCount(filter)

	return
}

// 待验收
func GetWaitFinishCount(createUserId, employeeId string) (cnt int64, err error) {
	filter := make(map[string]interface{})
	filter["submit_status"] = models.Job_Submit_Wait_Acceptance
	if createUserId != "" {
		filter["create_user_id"] = createUserId
	}

	if employeeId != "" {
		filter["employee_id"] = employeeId
	}

	cnt, err = GetCount(filter)

	return
}
