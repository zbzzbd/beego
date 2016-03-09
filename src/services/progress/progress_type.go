package progress

type ProgressType int

const (
	_                 ProgressType = iota
	PT_Create                      // 创建作业
	PT_Modify                      // 修改作业
	PT_Valid                       // 审核作业
	PT_Claim                       // 认领作业
	PT_Submit                      // 提交作业
	PT_Complain                    // 作业投诉
	PT_ComplainReplay              // 投诉回复
	PT_Assign                      // 作业转发
)
