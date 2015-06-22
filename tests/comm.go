package tests

import (
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app"
)

var username string = "13456789012"
var password string = "123456"
var newPassword string = "12345678"
var smsCode string = ""
var token string=""
var userId int = 0

var username1 string = "13400001111"
var password1 string = "654321"
var userId1 int = 0
var token1 string=""
var smsCode1 string = ""

var companyName string = "西安一元网络科技有限公司"
var companyId int = 0

var projectName string = "完成团队创建和加入"
var newProjectName string = "完成团队创建、加入和详情"
var projectId int = 0

var taskName string = "备份代码到github"
var taskInfo string = "定时更新代码"
var taskTransferInfo string = "我的部分已完成，该你了"
var taskId int = 0


func (t *AppTest) ClearUserTable() {
	sql := "truncate table users"
	_, err := app.Engine.Exec(sql)
	t.AssertEqual(nil, err)
}
func (t *AppTest) ClearUserProfileTable() {
	sql := "truncate table user_profiles"
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

func (t *AppTest) ClearTaskTable() {
	sql := "truncate table task"
	_, err := app.Engine.Exec(sql)
	t.AssertEqual(nil, err)
}


func (t *AppTest) ClearTaskTransferTable() {
	sql := "truncate table task_transfer"
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