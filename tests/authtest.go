package tests

import (
	"github.com/baiyuxiong/track/app"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
	"net/url"
)


func (t *AppTest) StartTestAuth() {
	t.InvalidRequest()
	t.ClearUserProfileTable();
	t.ClearUserTable()
	t.ClearSmsCodeTable()
	t.LoginValidation()
	t.LoginUserNotExists()
	t.RegValidation()
	t.CommSmsValidation()
	t.RegUser()
	t.RegUserAgain()
	t.LoginCorrect()
	t.ChangePasswordTest()
	t.LogoutCorrect()
	t.GetPasswordTest()
	//t.ClearUserTable()
}

func (t *AppTest) InvalidRequest() {
	data := make(url.Values)
	t.PostForm("/auth/login", data)
	t.AssertContains("INVALID_REQUEST")
	
	t.PostForm("/auth/logout", data)
	t.AssertContains("INVALID_REQUEST")
	
	t.PostForm(t.GenUrl("/auth/login",token), data)
	t.AssertNotContains("INVALID_REQUEST")
	
	t.PostForm("/auth/logout?"+utils.URL_CLIENT_ID_KEY+"=123", data)
	t.AssertContains("INVALID_REQUEST")
	
	t.PostForm(t.GenUrl("/auth/logout",token), data)
	t.AssertContains("INVALID_REQUEST")
}


func (t *AppTest) LoginValidation() {
	data := make(url.Values)
	t.PostForm(t.GenUrl("/auth/login",token), data)
	t.AssertContains("手机号不能为空")

	data["username"] =[]string{"13456789012"}
	t.PostForm(t.GenUrl("/auth/login",token), data)
	t.AssertContains("密码不能为空")

	data["username"] = []string{"23456789012"}
	data["password"] = []string{"123456"}
	t.PostForm(t.GenUrl("/auth/login",token), data)
	t.AssertContains("手机号格式不正确")
}

//用户不存在
func (t *AppTest) LoginUserNotExists() {
	data := make(url.Values)
	data["username"] = []string{"13456789012"}
	data["password"] = []string{"123456"}
	t.PostForm(t.GenUrl("/auth/login",token), data)
	t.AssertContains("用户名或密码出错")
}

func (t *AppTest) RegValidation() {
	data := make(url.Values)
	t.PostForm(t.GenUrl("/auth/reg",token), data)
	t.AssertContains("手机号不能为空")

	data["username"] = []string{"13456789012"}
	t.PostForm(t.GenUrl("/auth/reg",token), data)
	t.AssertContains("密码不能为空")

	data["username"] = []string{"23456789012"}
	data["password"] = []string{"123456"}
	t.PostForm(t.GenUrl("/auth/reg",token), data)
	t.AssertContains("验证码不能为空")

	data["username"] = []string{"23456789012"}
	data["password"] = []string{"123456"}
	data["smsCode"] = []string{"123456"}
	t.PostForm(t.GenUrl("/auth/reg",token), data)
	t.AssertContains("手机号格式不正确")

	data["username"] = []string{"13456789012"}
	data["password"] = []string{"123456"}
	data["smsCode"] = []string{"123456"}
	t.PostForm(t.GenUrl("/auth/reg",token), data)
	t.AssertContains("请先请求验证码")
}

func (t *AppTest) CommSmsValidation() {
	data := make(url.Values)
	t.PostForm(t.GenUrl("/comm/sendSms",token), data)
	t.AssertContains("手机号不能为空")

	data["username"] = []string{"23456789012"}
	t.PostForm(t.GenUrl("/comm/sendSms",token), data)
	t.AssertContains("手机号格式不正确")

	data["username"] = []string{username}
	t.PostForm(t.GenUrl("/comm/sendSms",token), data)
	t.AssertContains("200")
	
	t.PostForm(t.GenUrl("/comm/sendSms",token), data)
	t.AssertContains("200")
	
	var s = &models.SmsCode{Username: username}
	has, _ := app.Engine.Get(s)
	t.AssertEqual(true, has)

	smsCode = s.Code
}

func (t *AppTest) RegUser() {
	data := make(url.Values)
	data["username"] = []string{username}
	data["password"] = []string{password}

	tempCode := utils.Sms_code()
	data["smsCode"] = []string{tempCode}
	t.PostForm(t.GenUrl("/auth/reg",token), data)

	if tempCode == smsCode {
		t.AssertContains("200")
	} else {
		t.AssertContains("验证码不正确")
		data["smsCode"] = []string{smsCode}
		t.PostForm(t.GenUrl("/auth/reg",token), data)
		t.AssertContains("200")
	}
	
	var u = &models.Users{Username: username}
	has, err := app.Engine.Get(u)
	t.AssertEqual(nil,err)
	t.AssertEqual(true,has)
}
func (t *AppTest) RegUserAgain() {
	data := make(url.Values)
	data["username"] = []string{username}
	data["password"] = []string{password}
	data["smsCode"] = []string{smsCode}
	
	t.PostForm(t.GenUrl("/auth/reg",token), data)
	t.AssertContains("用户名名已被注册，不可用")
}

func (t *AppTest) LoginCorrect() {
	data := make(url.Values)
	data["username"] = []string{username}
	data["password"] = []string{password}
	t.PostForm(t.GenUrl("/auth/login",""), data)
	t.AssertContains("token")

	var u = &models.Users{Username: username}
	has, err := app.Engine.Get(u)
	t.AssertEqual(nil,err)
	t.AssertEqual(true,has)
	token = u.Token
}
func (t *AppTest) LogoutCorrect() {
	var u = &models.Users{Username: username}
	has, err := app.Engine.Get(u)
	t.AssertEqual(nil,err)
	t.AssertEqual(true,has)
	
	token = u.Token
	data := make(url.Values)
	t.PostForm(t.GenUrl("/auth/logout",token), data)
	t.AssertContains("200")
	
	var u1 = &models.Users{Username: username}
	has, err = app.Engine.Get(u1)
	t.AssertEqual(nil,err)
	t.AssertEqual(true,has)
	t.AssertEqual(u1.Token,"")
}

func (t *AppTest) ChangePasswordTest() {
	newPassword := utils.RandString(utils.Alnum,8)

	//println("token - " , token)
	data := make(url.Values)
	data["oldPassword"] = []string{password+"a"}
	data["newPassword"] = []string{newPassword}
	t.PostForm(t.GenUrl("/auth/changePassword",token), data)
	t.AssertContains("修改密码失败，请确认旧密码正确")

	data["oldPassword"] = []string{password}
	t.PostForm(t.GenUrl("/auth/changePassword",token), data)
	t.AssertContains("200")

	data["username"] = []string{username}
	data["password"] = []string{newPassword}
	t.PostForm(t.GenUrl("/auth/login",token), data)
	t.AssertContains("token")

	password = newPassword
}
func (t *AppTest) GetPasswordTest() {
	//username,new_password,sms_code
	data := make(url.Values)
	data["username"] = []string{"10000000000"}
	data["newPassword"] = []string{newPassword}
	data["smsCode"] = []string{smsCode}
	t.PostForm(t.GenUrl("/auth/getPassword",token), data)
	t.AssertContains("用户名不存在")

	data["username"] = []string{username}
	data["smsCode"] = []string{smsCode+"a"}
	t.PostForm(t.GenUrl("/auth/getPassword",token), data)
	t.AssertContains("验证码不正确")

	data["smsCode"] = []string{smsCode}
	t.PostForm(t.GenUrl("/auth/getPassword",token), data)
	t.AssertContains("200")

	data["username"] = []string{username}
	data["password"] = []string{newPassword}
	t.PostForm(t.GenUrl("/auth/login",token), data)
	t.AssertContains("token")

	password = newPassword
}