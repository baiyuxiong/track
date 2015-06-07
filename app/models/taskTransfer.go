package models

import (
	"time"
)

type TaskTransfer struct {
	Id         int64     `json:"id" xorm:"BIGINT(20)"`
	CompanyId  int       `json:"company_id" xorm:"not null INT(11)"`
	ProjectId  int       `json:"project_id" xorm:"not null INT(11)"`
	TaskId     int       `json:"task_id" xorm:"not null INT(11)"`
	InChargeId int       `json:"in_charge_id" xorm:"not null INT(11)"`
	AssignTo   int       `json:"assign_to" xorm:"INT(11)"`
	Progress   int       `json:"progress" xorm:"not null TINYINT(4)"`
	CreatedAt  time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"not null DATETIME"`
}
