package models

import (
	"time"
)

type Users struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	IpAddress   string    `json:"ip_address" xorm:"not null VARCHAR(15)"`
	Username    string    `json:"username" xorm:"not null CHAR(11)"`
	Password    string    `json:"password" xorm:"not null VARCHAR(80)"`
	Salt        string    `json:"salt" xorm:"CHAR(16)"`
	Email       string    `json:"email" xorm:"not null VARCHAR(100)"`
	Token       string    `json:"token" xorm:"VARCHAR(32)"`
	IsActivited int       `json:"is_activited" xorm:"not null TINYINT(4)"`
	ActivatedAt time.Time `json:"activated_at" xorm:"DATETIME"`
	CreatedAt   time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"not null DATETIME"`
}
