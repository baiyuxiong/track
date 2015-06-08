package tests

import (
	"github.com/baiyuxiong/track/app"
	"github.com/baiyuxiong/track/app/models"
	"net/url"
	"strconv"
)

func (t *AppTest) StartTestCompanyUsers() {
	token = ""
	t.GetToken()
	t.RegAnotherUser()
	t.AddCompanyUsers()
	t.CheckCompanyUsers()
	t.DeleteCompanyUsers()
}

func (t *AppTest) RegAnotherUser()  {
	smsData := make(url.Values)
	smsData["username"] = []string{username1}
	t.PostForm(t.GenUrl("/comm/send_sms",""), smsData)
	t.AssertContains("200")

	var s = &models.SmsCode{Username: username1}
	has, _ := app.Engine.Get(s)
	t.AssertEqual(true, has)

	smsCode1 = s.Code

	//reg
	data := make(url.Values)
	data["username"] = []string{username1}
	data["password"] = []string{password1}
	data["sms_code"] = []string{smsCode1}
	t.PostForm(t.GenUrl("/auth/reg",token), data)
	t.AssertContains("200")

	//login
	data1 := make(url.Values)
	data1["username"] = []string{username1}
	data1["password"] = []string{password1}
	t.PostForm(t.GenUrl("/auth/login",""), data1)
	t.AssertContains("token")

	var u = &models.Users{Username: username1}
	has,_ = app.Engine.Get(u)
	t.AssertEqual(true, has)
	userId1 = u.Id
	token1 = u.Token
}

func (t *AppTest) AddCompanyUsers() {
	companys := make([]models.Company, 0)
	app.Engine.Find(&companys)

	t.AssertNotEqual(len(companys),0)
	company := companys[0]

	companyId = company.Id
	t.AssertNotEqual(companyId,0)

	data := make(url.Values)
	data["company_id"] = []string{strconv.Itoa(companyId)}
	data["user_id"] = []string{strconv.Itoa(userId)}
	t.PostForm(t.GenUrl("/company_users/add",token1), data)
	t.AssertContains("没有权限")

	data["user_id"] = []string{strconv.Itoa(userId1)}
	t.PostForm(t.GenUrl("/company_users/add",token), data)
	t.AssertContains("200")

	t.PostForm(t.GenUrl("/company_users/add",token), data)
	t.AssertContains("用户已存在")
}

func (t *AppTest) CheckCompanyUsers() {
	data := make(url.Values)
	data["company_id"] = []string{strconv.Itoa(companyId)}
	data["user_id"] = []string{strconv.Itoa(userId)}
	t.PostForm(t.GenUrl("/company_users/check",token1), data)
	t.AssertContains("没有权限")

	data["user_id"] = []string{strconv.Itoa(userId1+userId)}
	t.PostForm(t.GenUrl("/company_users/check",token), data)
	t.AssertContains("该用户未申请")

	data["user_id"] = []string{strconv.Itoa(userId1)}
	t.PostForm(t.GenUrl("/company_users/check",token), data)
	t.AssertContains("200")

	companyUser := new(models.CompanyUsers)
	has,err := app.Engine.Where("company_id = ? and user_id = ?",companyId,userId1).Get(companyUser)

	t.AssertEqual(has,true)
	t.AssertEqual(err,nil)
	t.AssertEqual(companyUser.Status,1)
}


func (t *AppTest) DeleteCompanyUsers() {
	data := make(url.Values)
	data["company_id"] = []string{strconv.Itoa(companyId)}
	data["user_id"] = []string{strconv.Itoa(userId)}
	t.PostForm(t.GenUrl("/company_users/delete",token1), data)
	t.AssertContains("没有权限")

	data["user_id"] = []string{strconv.Itoa(userId1+userId)}
	t.PostForm(t.GenUrl("/company_users/delete",token), data)
	t.AssertContains("该用户未申请")

	data["user_id"] = []string{strconv.Itoa(userId1)}
	t.PostForm(t.GenUrl("/company_users/delete",token), data)
	t.AssertContains("200")

	companyUser := new(models.CompanyUsers)
	has,err := app.Engine.Where("company_id = ? and user_id = ?",companyId,userId1).Get(companyUser)

	t.AssertEqual(has,true)
	t.AssertEqual(err,nil)
	t.AssertEqual(companyUser.Status,2)
}
