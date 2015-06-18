package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app"
	"github.com/baiyuxiong/track/app/lib"
	"time"
	"fmt"
)

type Project struct {
	BaseController
}

func (c Project) ListByOwner() revel.Result {
	projects := make([]models.Project, 0)
	err := app.Engine.Where("owner_id = ?", c.User.Id).OrderBy("updated_at desc").Find(&projects)
	if err != nil {
		return c.Err(err.Error())
	}
	return c.OK(projects)
}

func (c Project) ListByCompany(companyId int) revel.Result {
	if !lib.IsCompanyCheckedUser(companyId, c.User.Id) {
		return c.Err("没有权限")
	}

	projects := make([]models.Project, 0)
	err := app.Engine.Where("company_id = ?", companyId).OrderBy("updated_at desc").Find(&projects)

	if err != nil {
		return c.Err(err.Error())
	}
	return c.OK(projects)
}

type CompaniesAdnPorjects struct {
	Companys *lib.MyCompanies    `json:"companys"`
	Projects map[string][]models.Project    `json:"projects"`
}
func (c Project) ListCompanyAndProject() revel.Result {
	companys := lib.GetMyCompanies(c.User.Id)

	allCompanyProjects := make(map[string][]models.Project)

	for _, company := range companys.CompaniesOwned {
		projects := make([]models.Project, 0)
		app.Engine.Where("company_id = ?", company.Id).OrderBy("updated_at desc").Find(&projects)
		allCompanyProjects[fmt.Sprintf("%d", company.Id)] = projects
	}
	return c.OK(
		CompaniesAdnPorjects{
			Companys:companys,
			Projects:allCompanyProjects,
		})
}

func (c Project) Id(id int) revel.Result {
	Project := &models.Project{}
	_, err := app.Engine.Id(id).Get(Project)
	if err != nil {
		return c.Err(err.Error())
	}

	if !lib.IsCompanyCheckedUser(Project.CompanyId, c.User.Id) {
		return c.Err("没有权限")
	}

	return c.OK(Project)
}

func (c Project) Add(companyId int, name, info string) revel.Result {
	result := c.validateName(name)

	if nil != result {
		return result
	}

	if !lib.IsCompanyCheckedUser(companyId, c.User.Id) {
		return c.Err("没有权限")
	}

	prj := &models.Project{
		CompanyId:companyId,
		OwnerId:c.User.Id,
		Name:name,
		Info:info,
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
	}

	_, err := app.Engine.Insert(prj)
	if err != nil {
		return c.Err("添加失败，请联系管理员")
	}

	return c.OK(prj)
}

func (c Project) Update(id int, name, info string) revel.Result {
	c.Validation.Required(id).Message("参数错误")

	result := c.validateName(name)
	if result != nil {
		return result;
	}

	if !c.isOwner(id) {
		return c.Err("没有权限")
	}

	cmp := &models.Project{
		Name:name,
		Info:info,
		UpdatedAt:time.Now(),
	}

	_, err := app.Engine.Id(id).Cols("name").Cols("info").Update(cmp)
	if err != nil {
		return c.Err("更新失败")
	}
	return c.OK("")
}

func (c Project) validateName(name string) revel.Result {
	c.Validation.Required(name).Message("名称不能为空")
	c.Validation.MinSize(name, utils.PROJECT_NAME_MINSIZE).Message("名称不能太短")

	if c.Validation.HasErrors() {
		return c.Err(utils.ValidationErrorToString(c.Validation.Errors))
	}
	return nil
}

func (c Project) isOwner(id int) bool {
	Project := &models.Project{}
	_, err := app.Engine.Id(id).Get(Project)
	if err != nil {
		return false
	}

	if Project.OwnerId != c.User.Id {
		return false
	}
	return true
}