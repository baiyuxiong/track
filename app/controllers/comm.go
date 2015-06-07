package controllers

import (
	"github.com/baiyuxiong/track/app"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app/models"
	"github.com/revel/revel"
	"time"
	"regexp"
)

type Comm struct {
	BaseController
}

func (c Comm) SendSms(username string) revel.Result {
	c.Validation.Required(username).Message("手机号不能为空")
	c.Validation.Match(username, regexp.MustCompile("^(1)\\d{10}$")).Message("手机号格式不正确")

	if c.Validation.HasErrors() {
		return c.Err(utils.ValidationErrorToString(c.Validation.Errors))
	}

	//验证是否已被注册
	var s = &models.SmsCode{Username: username}
	has, _ := app.Engine.Get(s)

	code := utils.Sms_code()
	s.Code = code
	s.UpdatedAt = time.Now()

	var err error
	if has {
		_, err = app.Engine.Update(s)
	} else {
		_, err = app.Engine.Insert(s)
	}

	if err != nil {
		return c.Err(err.Error())
	} else {
		return c.OK(nil)
	}
}
