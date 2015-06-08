package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app/lib"
	"github.com/baiyuxiong/track/app"
	"time"
)

type Company struct {
	BaseController
}

func (c Company) List() revel.Result {
	companys := make([]models.Company, 0)
	err := app.Engine.Where("owner_id = ?",c.User.Id).OrderBy("updated_at desc").Find(&companys)
	if err != nil{
		return c.Err(err.Error())
	}
	return c.OK(companys)
}

func (c Company) Id(id int) revel.Result {
	company := &models.Company{}
	_, err := app.Engine.Id(id).Get(company)
	if err != nil{
		return c.Err(err.Error())
	}
	return c.OK(company)
}

func (c Company) Add(name ,info ,phone,address string) revel.Result {
	result := c.validateName(name)

	if nil != result{
		return result
	}

	now := time.Now()
	cmp := &models.Company{
		OwnerId:c.User.Id,
		Name:name,
		Info:info,
		Phone:phone,
		Address:address,
		CreatedAt:now,
		UpdatedAt:now,
	}

	_, err := app.Engine.Insert(cmp)
	if err != nil{
		return c.Err("添加失败，请联系管理员")
	}

	//add company users
	cu := &models.CompanyUsers{
		CompanyId: cmp.Id,
		UserId:c.User.Id,
		Status:utils.COMPANY_USER_STATUS_CHECK_YES,
		UpdatedAt:now,
		CreatedAt:now,
	}
	_, err = app.Engine.Insert(cu)
	if err != nil{
		app.Engine.Id(cmp.Id).Delete(cmp)
		return c.Err("添加失败，数据已存在")
	}

	return c.OK(cmp)
}

func (c Company) Update(id int,name,info,phone,address string) revel.Result {
	c.Validation.Required(id).Message("参数错误")

	result := c.validateName(name)
	if result != nil{
		return result;
	}

	isOwner,_ := lib.IsCompanyOwner(nil,id,c.User.Id)
	if !isOwner{
		return c.Err("没有权限")
	}

	cmp := &models.Company{
		Name:name,
		Phone:phone,
		Address:address,
		UpdatedAt:time.Now(),
	}

	_, err := app.Engine.Id(id).Cols("name").Cols("phone").Cols("address").Update(cmp)
	if err != nil{
		return c.Err("更新失败")
	}
	return c.OK("")
}

func (c Company) UpdateLogo(id int,logo string) revel.Result {

	c.Validation.Required(id).Message("参数错误")
	c.Validation.Required(logo).Message("参数错误")

	if c.Validation.HasErrors() {
		return c.Err(utils.ValidationErrorToString(c.Validation.Errors))
	}
	isOwner,_ := lib.IsCompanyOwner(nil,id,c.User.Id)
	if !isOwner{
		return c.Err("没有权限")
	}

	cmp := &models.Company{
		Logo:logo,
	}
	_, err := app.Engine.Id(id).Cols("logo").Update(cmp)
	if err != nil{
		return c.Err("更新失败")
	}

	return c.OK("")
}

func (c Company) validateName(name string) revel.Result{
	c.Validation.Required(name).Message("名称不能为空")
	c.Validation.MinSize(name,utils.COMPANY_NAME_MINSIZE).Message("名称不能太短")

	if c.Validation.HasErrors() {
		return c.Err(utils.ValidationErrorToString(c.Validation.Errors))
	}
	return nil
}