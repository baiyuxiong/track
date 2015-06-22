package controllers

import (
	"github.com/baiyuxiong/track/app"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/lib"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/revel/revel"
	"regexp"
	"time"
	"strings"
)

type Auth struct {
	BaseController
}

func (c Auth) Reg(username, password, smsCode string) revel.Result {
	c.Validation.Required(username).Message("手机号不能为空")
	c.Validation.Required(password).Message("密码不能为空")
	c.Validation.Required(smsCode).Message("验证码不能为空")

	c.Validation.Match(username, regexp.MustCompile("^(1)\\d{10}$")).Message("手机号格式不正确")

	if c.Validation.HasErrors() {
		return c.Err(utils.ValidationErrorToString(c.Validation.Errors))
	}

	//验证是否已被注册
	var u = &models.Users{Username: username}
	has, _ := app.Engine.Get(u)

	if has {
		return c.Err("用户名名已被注册，不可用")
	}

	//短信验证码
	isOk, message := lib.CheckSmsCode(u.Username, smsCode)
	if !isOk {
		return c.Err(message)
	}

	//注册
	now := time.Now()
	u.IpAddress = strings.Split(c.Request.RemoteAddr, ":")[0]
	u.Salt = utils.Salt()
	u.Password = utils.EncryptPassword(u.Salt, password)
	u.Token = ""
	u.IsActivited = 1
	u.ActivatedAt = now
	u.CreatedAt = now
	u.UpdatedAt = now

	_, err := app.Engine.Insert(u)
	if err != nil{
		return c.Err("注册用户失败")
	}

	profile := new(models.UserProfiles)
	profile.UserId = u.Id
	profile.Name = u.Username
	profile.Phone = u.Username
	profile.CreatedAt = now
	profile.UpdatedAt = now
	_,err = app.Engine.Insert(profile)
	if err != nil{
		app.Engine.Id(u.Id).Delete(u)
		return c.Err("注册用户失败"+err.Error())
	}

	return c.OK(u)
}

func (c Auth) Login(username, password string) revel.Result {
	c.Validation.Required(username).Message("手机号不能为空")
	c.Validation.Required(password).Message("密码不能为空")
	c.Validation.Match(username, regexp.MustCompile("^(1)\\d{10}$")).Message("手机号格式不正确")

	if c.Validation.HasErrors() {
		return c.Err(utils.ValidationErrorToString(c.Validation.Errors))
	}

	var u = &models.Users{Username: username}
	has, _ := app.Engine.Get(u)

	if !has {
		return c.Err("用户名或密码出错")	}

	if utils.EncryptPassword(u.Salt, password) != u.Password {
		return c.Err("用户名或密码出错")
	}

	token := utils.Token(u.Id,u.IpAddress)
	u.Token = token
	u.UpdatedAt = time.Now()
	_, err := app.Engine.Id(u.Id).Cols("token").Cols("updated_at").Update(u)
	if err != nil {
		app.Engine.Id(u.Id).Delete(new(models.Users))
		return c.Err("登录失败")
	}

	return c.OK(u)
}


func (c Auth) ChangePassword(oldPassword,newPassword string) revel.Result {
	temp := utils.EncryptPassword(c.User.Salt,oldPassword)
	if temp == c.User.Password{
		password_encrypt := utils.EncryptPassword(c.User.Salt,newPassword)
		_, err := app.Engine.Id(c.User.Id).Cols("password").Update(&models.Users{Password: password_encrypt,UpdatedAt:time.Now()})
		if err == nil {
			return c.OK("")
		}else{
			return c.Err("修改密码失败，请联系管理员")
		}
	}
	return c.Err("修改密码失败，请确认旧密码正确")
}

func (c Auth) GetPassword(username,newPassword,smsCode string) revel.Result {
	var u = &models.Users{Username: username}
	has, _ := app.Engine.Get(u)

	if !has {
		return c.Err("用户名不存在")
	}

	//短信验证码
	isOk, message := lib.CheckSmsCode(username, smsCode)
	if !isOk {
		return c.Err(message)
	}
	
	password_encrypt := utils.EncryptPassword(u.Salt,newPassword)
	_, err := app.Engine.Id(u.Id).Cols("password").Update(&models.Users{Password: password_encrypt,UpdatedAt:time.Now()})
	
	if err == nil {
		return c.OK(u)
	}
	return c.Err("修改密码失败，请联系管理员")
}


func (c Auth) Logout() revel.Result {
	_, err := app.Engine.Id(c.User.Id).Cols("token").Update(&models.Users{Token: "",UpdatedAt:time.Now()})
	if err != nil{
		return c.Err("退出失败")
	}
	return c.OK("")
}