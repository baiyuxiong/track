package controllers

import (
	"github.com/revel/revel"
)

type Task struct {
	BaseController
}

func (c Task) Index() revel.Result {
	return c.OK("")
}
