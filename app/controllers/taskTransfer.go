package controllers

import (
	"github.com/revel/revel"
)

type TaskTransfer struct {
	BaseController
}

func (c TaskTransfer) Index() revel.Result {
	return c.OK("")
}
