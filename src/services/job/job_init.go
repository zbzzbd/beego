package job

import "services/progress"

const (
	DefaultTimeLayout = "2006-01-02 15:04"
)

func init() {
	//注册事件
	registerJobProgress()
}

func registerJobProgress() {
	//let it crash
	progress.Register(progress.PT_Create)
	progress.Register(progress.PT_Modify)
}
