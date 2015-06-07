package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/lib"
	"github.com/baiyuxiong/track/app"
)

type CompanyUsers struct {
	BaseController
}

func (c CompanyUsers) List(company_id int) revel.Result {

	if lib.IsCompanyUser(company_id,c.User.Id){
		companyUsers := make([]models.CompanyUsers, 0)
		err := app.Engine.Where("company_id = ?",company_id).OrderBy("created_at desc").Find(&companyUsers)
		if err != nil{
			return c.Err(err.Error())
		}
		return c.OK(companyUsers)
	}
	return c.Err("没有权限")
}

func (c CompanyUsers) Add(company_id int) revel.Result {

}

func (c CompanyUsers) Update(company_id int) revel.Result {
	isOwner,_ := lib.IsCompanyOwner(nil,company_id,c.User.Id)
	if !isOwner{

	}
}

func (c CompanyUsers) Delete(company_id int) revel.Result {
	isOwner,_ := lib.IsCompanyOwner(nil,company_id,c.User.Id)
	if !isOwner{

	}
}
