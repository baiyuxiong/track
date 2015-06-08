package tests

import (
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app"
)

var username string = "13456789012"
var password string = "123456"
var new_password string = "12345678"
var sms_code string = ""
var token string=""
var userId int = 0

var username1 string = "13400001111"
var password1 string = "654321"
var userId1 int = 0
var token1 string=""
var smsCode1 string = ""

var companyName string = "track cmp"
var companyId int = 0

var projectName string = "track project"
var newProjectName string = "track project 123"
var projectId int = 0


func (t *AppTest) ClearUserTable() {
	sql := "truncate table users"
	_, err := app.Engine.Exec(sql)
	t.AssertEqual(nil, err)
}
func (t *AppTest) ClearSmsCodeTable() {
	sql := "truncate table sms_code"
	_, err := app.Engine.Exec(sql)
	t.AssertEqual(nil, err)
}

func (t *AppTest) ClearCompanyTable() {
	sql := "truncate table company"
	_, err := app.Engine.Exec(sql)
	t.AssertEqual(nil, err)
}

func (t *AppTest) ClearProjectTable() {
	sql := "truncate table project"
	_, err := app.Engine.Exec(sql)
	t.AssertEqual(nil, err)
}

func (t *AppTest) ClearCompanyUsersTable() {
	sql := "truncate table company_users"
	_, err := app.Engine.Exec(sql)
	t.AssertEqual(nil, err)
}

func (t *AppTest) GenUrl(base,token string,) string{
	base = base + "?"+ utils.URL_CLIENT_ID_KEY+"="+utils.URL_CLIENT_ID
	base = base + "&"+ utils.URL_TOKEN_KEY+"="+token
	return base
}