package tests

import (
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app"
	"net/url"
	"strconv"
)

func (t *AppTest) StartTestProject() {
//	token = ""
//	t.GetToken()
//	t.ClearCompanyUsersTable()
	t.ClearProjectTable()
	t.User1UncheckFromCompanyMember()
	t.AddProjectTest()
	t.UpdateProjectTest()
	t.ListByOwnerTest()
	t.ListByCompanyTest()
	t.IdTest()
}

func (t *AppTest) User1UncheckFromCompanyMember() {
	affected, err := app.Engine.Where("user_id = ?",userId1).Cols("status").Update(&models.CompanyUsers{Status: utils.COMPANY_USER_STATUS_CHECK_NO})
	t.AssertNotEqual(affected,0)
	t.AssertEqual(err,nil)
}

func (t *AppTest) AddProjectTest() {
	data := make(url.Values)
	data["companyId"] = []string{strconv.Itoa(companyId)}
	data["info"] = []string{projectName}
	t.PostForm(t.GenUrl("/project/add",token1), data)
	t.AssertContains("名称不能为空")

	data["name"] = []string{"123"}
	t.PostForm(t.GenUrl("/project/add",token1), data)
	t.AssertContains("名称不能太短")

	data["name"] = []string{projectName}
	t.PostForm(t.GenUrl("/project/add",token1), data)
	t.AssertContains("没有权限")

	t.PostForm(t.GenUrl("/project/add",token), data)
	t.AssertContains("200")

	var p = &models.Project{Name: projectName}
	has, _ := app.Engine.Get(p)

	t.AssertEqual(has,true)

	projectId = p.Id
}

func (t *AppTest) UpdateProjectTest() {
	data := make(url.Values)
	data["id"] = []string{strconv.Itoa(projectId)}
	data["info"] = []string{projectName}
	t.PostForm(t.GenUrl("/project/update",token), data)
	t.AssertContains("名称不能为空")

	data["name"] = []string{"123"}
	t.PostForm(t.GenUrl("/project/update",token), data)
	t.AssertContains("名称不能太短")

	data["name"] = []string{newProjectName}
	t.PostForm(t.GenUrl("/project/update",token1), data)
	t.AssertContains("没有权限")

	t.PostForm(t.GenUrl("/project/update",token), data)
	t.AssertContains("200")

	var p = &models.Project{Name: newProjectName}
	has, _ := app.Engine.Get(p)

	t.AssertEqual(has,true)
}

func (t *AppTest) ListByOwnerTest() {
	t.Get(t.GenUrl("/project/listByOwner",token1))
	t.AssertNotContains(newProjectName)

	t.Get(t.GenUrl("/project/listByOwner",token))
	t.AssertContains(newProjectName)
}
func (t *AppTest) ListByCompanyTest() {
	t.Get(t.GenUrl("/project/listByCompany",token1)+"&companyId="+strconv.Itoa(companyId))
	t.AssertNotContains(newProjectName)
	t.AssertContains("没有权限")

	t.Get(t.GenUrl("/project/listByCompany",token)+"&companyId="+strconv.Itoa(companyId))
	t.AssertContains(newProjectName)
}
func (t *AppTest)IdTest() {
	t.Get(t.GenUrl("/project/id",token1)+"&id="+strconv.Itoa(projectId))
	t.AssertNotContains(newProjectName)
	t.AssertContains("没有权限")

	t.Get(t.GenUrl("/project/id",token)+"&id="+strconv.Itoa(projectId))
	t.AssertContains(newProjectName)
}