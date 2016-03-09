package models

import "time"

const TABLE_NAME_FILE = "files"

type File struct {
	Id    uint   `orm:"pk;auto"`
	Type  uint   `orm:"default(1)"` // 文件类型： 1 － job， 2 － job_modify， 3 － job_valid， 4 － job_claim， 5 － job_submit， 6 － job_complain， 7 － job_complain_replay
	RelId uint   // 由文件类型决定rel_id对应是哪个表的id
	Url   string `orm:"size(64)"`
	Name  string `orm:"size(32)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (c *File) TableName() string {
	return TABLE_NAME_FILE
}
