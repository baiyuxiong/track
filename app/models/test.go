package models

type Test struct {
	Id         int `json:"id" xorm:"not null pk autoincr INT(11)"`
	AssignFrom int `json:"assign_from" xorm:"not null INT(11)"`
}
