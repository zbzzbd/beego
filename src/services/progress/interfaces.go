package progress

type JobProgresser interface {
	GetJobId() uint
	GetType() ProgressType
	GetTableName() string
	GetDesc() string
	GetPrimaryKey() uint
}
