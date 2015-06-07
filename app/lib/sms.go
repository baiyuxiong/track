package lib

import (
	"github.com/baiyuxiong/track/app"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
	"time"
)

func CheckSmsCode(Username string, code string) (isOk bool, message string) {
	isOk = false
	var s = &models.SmsCode{Username: Username}

	has, _ := app.Engine.Get(s)

	if !has {
		message = "请先请求验证码"
	} else if time.Now().Sub(s.UpdatedAt) > time.Duration(utils.SMS_EXPIRY*time.Minute) {
		message = "验证码已过期"
	} else if code != s.Code {
		message = "验证码不正确"
	} else {
		isOk = true
	}

	return
}
