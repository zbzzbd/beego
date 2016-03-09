package job

import ()

type JobsCountForManager struct {
	Count            int64
	FinishCount      int64
	WaitValidCount   int64
	ValidRefuseCount int64
	DoingCount       int64
}

type JobsCountForBussinessman struct {
	CreateCount int64
	FinishCount int64
	ModifyCount int64
	CancelCount int64
	DoingCount  int64
}
type JobsCountForTech struct {
	FinishCount     int64
	WaitClaimCount  int64
	DoingCount      int64
	WaitFinishCount int64
}

func (j *JobsCountForManager) Get() *JobsCountForManager {

	j.Count, _ = GetCount(nil)
	j.FinishCount, _ = GetFinishCount("", "")
	j.WaitValidCount, _ = GetWaitValidCount("", "")
	j.ValidRefuseCount, _ = GetValidRefuseCount("", "")
	j.DoingCount, _ = GetDoingCount("", "")
	return j
}
func (j *JobsCountForBussinessman) Get(bussiness_user string) *JobsCountForBussinessman {
	if bussiness_user != "" {
		j.CreateCount, _ = GetCreateCount(bussiness_user, "")
		j.FinishCount, _ = GetFinishCount(bussiness_user, "")
		j.ModifyCount, _ = GetModifyCount(bussiness_user, "")
		j.CancelCount, _ = GetCancelCount(bussiness_user, "")
		j.DoingCount, _ = GetDoingCount(bussiness_user, "")

	}
	return j
}
func (j *JobsCountForTech) Get(tech_user string) *JobsCountForTech {

	if tech_user != "" {
		j.FinishCount, _ = GetFinishCount("", tech_user)

		j.WaitClaimCount, _ = GetWaitClaimCount("", tech_user)
		j.DoingCount, _ = GetDoingCount("", tech_user)
		j.WaitFinishCount, _ = GetWaitFinishCount("", tech_user)
	}
	return j
}
