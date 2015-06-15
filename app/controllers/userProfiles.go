package controllers

import (
	"github.com/revel/revel"
)

type UserProfiles struct {
	BaseController
}

func (c UserProfiles) Index() revel.Result {
	return c.OK("")
}
