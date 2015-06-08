package tests

import (
	"github.com/baiyuxiong/track/app"
	"github.com/baiyuxiong/track/app/models"
	"net/url"
	"strconv"
)


func (t *AppTest) StartTestCompany() {
	token = ""
	t.GetToken()
	t.ClearCompanyTable()
	t.ClearCompanyUsersTable()
	t.AddCompany()
	t.UpdateCompany()
	t.UpdateCompanyLogo()
	t.CompanyByID()
	t.CompanyList()
	t.AuthTest()
}

func (t *AppTest) GetToken() {
	data := make(url.Values)
	data["username"] = []string{username}
	data["password"] = []string{new_password}
	t.PostForm(t.GenUrl("/auth/login",token), data)

	t.AssertContains("token")

	var u = &models.Users{Username: username}
	has, err := app.Engine.Get(u)
	t.AssertEqual(nil,err)
	t.AssertEqual(true,has)

	t.AssertNotEqual(u.Token,"")
	token = u.Token
	userId = u.Id
}

//add
func (t *AppTest) AddCompany() {
	data := make(url.Values)
	t.PostForm(t.GenUrl("/company/add",token), data)
	t.AssertContains("名称不能为空")

	data["name"] = []string{"abc"}
	t.PostForm(t.GenUrl("/company/add",token), data)
	t.AssertContains("名称不能太短")

	data["name"] = []string{companyName}
	t.PostForm(t.GenUrl("/company/add",token), data)
	t.AssertContains("200")

	c := new(models.Company)
	has, err := app.Engine.Where("name=?", companyName).Get(c)

	t.AssertEqual(has,true)
	t.AssertEqual(err,nil)

	companyId = c.Id

	cu := new(models.CompanyUsers)
	has, err = app.Engine.Where("company_id=? and user_id=?", companyId,userId).Get(cu)

	t.AssertEqual(has,true)
	t.AssertEqual(err,nil)
}

//update
func (t *AppTest) UpdateCompany() {
	data := make(url.Values)
	t.PostForm(t.GenUrl("/company/update",token), data)
	t.AssertContains("参数错误")

	data["id"] = []string{strconv.Itoa(companyId)}
	data["name"] = []string{"abc"}
	t.PostForm(t.GenUrl("/company/update",token), data)
	t.AssertContains("名称不能太短")

	companyName = companyName+"a"
	data["name"] = []string{companyName}
	t.PostForm(t.GenUrl("/company/update",token), data)
	t.AssertContains("200")

	c := new(models.Company)
	has, err := app.Engine.Where("name=?", companyName).Get(c)

	t.AssertEqual(has,true)
	t.AssertEqual(err,nil)
}

//update
func (t *AppTest) UpdateCompanyLogo() {
	data := make(url.Values)
	t.PostForm(t.GenUrl("/company/update_logo",token), data)
	t.AssertContains("参数错误")

	logo := "abcdef"
	data["id"] = []string{strconv.Itoa(companyId)}
	data["logo"] = []string{logo}
	t.PostForm(t.GenUrl("/company/update_logo",token), data)
	t.AssertContains("200")

	c := new(models.Company)
	has, err := app.Engine.Where("logo=?", logo).Get(c)

	t.AssertEqual(has,true)
	t.AssertEqual(err,nil)
}
//by id
func (t *AppTest) CompanyByID() {
	t.Get(t.GenUrl("/company/id",token)+"&id="+strconv.Itoa(companyId))
	t.AssertContains(companyName)

	t.Get(t.GenUrl("/company/id",token)+"&id="+strconv.Itoa(companyId+1))
	t.AssertNotContains(companyName)
}

//list
func (t *AppTest) CompanyList() {
	t.Get(t.GenUrl("/company/list",token)+"&id="+strconv.Itoa(companyId))
	t.AssertContains(companyName)
}

//update auth
func (t *AppTest) AuthTest() {
	cmp := &models.Company{
		OwnerId:userId+1,
	}
	_, err := app.Engine.Id(companyId).Cols("owner_id").Update(cmp)
	t.AssertEqual(nil,err)

	t.Get(t.GenUrl("/company/list",token))
	t.AssertNotContains(companyName)

	newCompanyName := companyName+"b"
	data := make(url.Values)
	data["id"] = []string{strconv.Itoa(companyId)}
	data["name"] = []string{newCompanyName}
	t.PostForm(t.GenUrl("/company/update",token), data)
	t.AssertContains("没有权限")

	delete(data,"name")
	data["logo"] = []string{"111111"}
	t.PostForm(t.GenUrl("/company/update_logo",token), data)
	t.AssertContains("没有权限")

	c := new(models.Company)
	has, err := app.Engine.Id(companyId).Get(c)

	t.AssertEqual(has,true)
	t.AssertEqual(err,nil)
	t.AssertNotEqual(data["logo"],c.Logo)
	t.AssertNotEqual(newCompanyName,c.Name)

	cmp.OwnerId = userId
	_, err = app.Engine.Id(companyId).Cols("owner_id").Update(cmp)
	t.AssertEqual(nil,err)
}