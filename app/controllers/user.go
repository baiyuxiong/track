package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app"
	"time"
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

func (c User) EditProfile(name, phone, avatar, avatar_thumb1, avatar_thumb2 string) revel.Result {
	u := new(models.UserProfiles)

	u.Name = name
	u.Phone = phone
	u.Avatar = avatar
	u.AvatarThumb1 = avatar_thumb1
	u.AvatarThumb2 = avatar_thumb2
	u.UpdatedAt = time.Now()

	app.Engine.Id(c.User.Id).Cols("name").Cols("phone").Cols("avatar").Cols("avatar_thumb1").Cols("avatar_thumb2").Cols("updated_at").Update(u)

	return c.OK(nil)
}