package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app"
)

type User struct {
	BaseController
}

func (c User) Index() revel.Result {
	return c.OK("")
}

type UserInfo struct {
	User        *models.Users `json:"user"`
	UserProfile *models.UserProfiles `json:"userProfile"`
}
func (c User) Me() revel.Result {
	var u = new(models.UserProfiles)
	has, _ := app.Engine.Get(u)
	if !has {
		return c.Err("用户信息不存在")
	}
	userInfo := &UserInfo{
		User:c.User,
		UserProfile:u,
	}
	return c.OK(userInfo)
}
