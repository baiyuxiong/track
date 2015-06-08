package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app/lib"
	"github.com/baiyuxiong/track/app"
	"time"
)

type CompanyUsers struct {
	BaseController
}



func (c CompanyUsers) List(company_id int) revel.Result {

	if lib.IsCompanyCheckedUser(company_id,c.User.Id){
		companyUsers := make([]models.CompanyUsers, 0)
		err := app.Engine.Where("company_id = ?",company_id).OrderBy("created_at desc").Find(&companyUsers)
		if err != nil{
			return c.Err(err.Error())
		}
		return c.OK(companyUsers)
	}
	return c.Err("没有权限")
}

func (c CompanyUsers) Add(company_id,user_id int) revel.Result {
	isOwner,_ := lib.IsCompanyOwner(nil,company_id,c.User.Id)
	if !isOwner{
		return c.Err("没有权限")
	}

	has := lib.IsCompanyUser(company_id,user_id)
	if has{
		return c.Err("用户已存在")
	}

	now := time.Now()
	cu := &models.CompanyUsers{
		CompanyId: company_id,
		UserId:user_id,
		Status:utils.COMPANY_USER_STATUS_CHECK_NO,
		UpdatedAt:now,
		CreatedAt:now,
	}
	_, err := app.Engine.Insert(cu)
	if err != nil{
		return c.Err("添加用户失败")
	}
	return c.OK(nil)
}

func (c CompanyUsers) Check(company_id,user_id int) revel.Result {
	isOk,message := c.changeStatus(company_id,user_id,utils.COMPANY_USER_STATUS_CHECK_YES)
	if !isOk{
		return c.Err(message)
	}
	return c.OK(nil)
}

func (c CompanyUsers) Delete(company_id,user_id int) revel.Result {
	isOk,message := c.changeStatus(company_id,user_id,utils.COMPANY_USER_STATUS_DELETE)
	if !isOk{
		return c.Err(message)
	}
	return c.OK(nil)
}

func (c CompanyUsers) changeStatus(company_id, user_id, status int) (isOk bool,message string) {
	isOk = false
	message = ""

	isOwner,_ := lib.IsCompanyOwner(nil,company_id,c.User.Id)
	if !isOwner{
		message = "没有权限"
		return
	}

	has := lib.IsCompanyUser(company_id,user_id)
	if !has{
		message = "该用户未申请"
		return
	}

	_, err := app.Engine.Where("company_id =  ? and user_id = ?", company_id, user_id).Cols("status").Update(&models.CompanyUsers{Status:status})

	if err != nil{
		message = err.Error()
		return
	}

	isOk = true
	return
}