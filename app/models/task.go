package models

import (
	"time"
)

type Task struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	CompanyId        int       `json:"company_id" xorm:"not null INT(11)"`
	ProjectId        int       `json:"project_id" xorm:"not null INT(11)"`
	OwnerId          int       `json:"owner_id" xorm:"not null INT(11)"`
	InChargeId       int       `json:"in_charge_id" xorm:"not null INT(11)"`
	InChargeProgress int       `json:"in_charge_progress" xorm:"not null TINYINT(4)"`
	Priority         int       `json:"priority" xorm:"not null TINYINT(4)"`
	Status           int       `json:"status" xorm:"not null TINYINT(4)"`
	Name             string    `json:"name" xorm:"not null VARCHAR(256)"`
	Info             string    `json:"info" xorm:"not null TEXT"`
	Deadline         time.Time `json:"deadline" xorm:"not null DATETIME"`
	CreatedAt        time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt        time.Time `json:"updated_at" xorm:"not null DATETIME"`
}
