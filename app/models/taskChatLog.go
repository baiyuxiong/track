package models

type TaskChatLog struct {
	TaskId int64  `json:"task_id" xorm:"not null BIGINT(11)"`
	Log    string `json:"log" xorm:"not null TEXT"`
}
