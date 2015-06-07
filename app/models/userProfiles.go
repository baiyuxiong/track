package models

import (
	"time"
)

type UserProfiles struct {
	UserId       int       `json:"user_id" xorm:"not null pk unique index INT(11)"`
	Gender       int       `json:"gender" xorm:"not null TINYINT(4)"`
	Name         string    `json:"name" xorm:"not null VARCHAR(32)"`
	Avatar       string    `json:"avatar" xorm:"not null VARCHAR(256)"`
	AvatarThumb1 string    `json:"avatar_thumb1" xorm:"not null VARCHAR(256)"`
	AvatarThumb2 string    `json:"avatar_thumb2" xorm:"not null VARCHAR(256)"`
	AvatarThumb3 string    `json:"avatar_thumb3" xorm:"not null VARCHAR(256)"`
	CreatedAt    time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt    time.Time `json:"updated_at" xorm:"not null DATETIME"`
}
