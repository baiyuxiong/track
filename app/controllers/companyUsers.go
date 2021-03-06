package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app/lib"
	"github.com/baiyuxiong/track/app"
	"time"
	"fmt"
)

type CompanyUsers struct {
	BaseController
}

type users struct {
	Users map[string]models.UserProfiles `json:"users"`
}

func (c CompanyUsers) List(companyId int) revel.Result {

	users := users{
		Users: make(map[string]models.UserProfiles),
	}

	if lib.IsCompanyCheckedUser(companyId, c.User.Id) {
		companyUsers := make([]models.CompanyUsers, 0)
		err := app.Engine.Where("company_id = ?", companyId).OrderBy("created_at desc").Find(&companyUsers)
		if err != nil {
			return c.Err(err.Error())
		}

		for _,companyUser := range companyUsers{
			userProfile := new(models.UserProfiles)
			has, _ := app.Engine.Id(companyUser.UserId).Get(userProfile)
			if has {
				users.Users[fmt.Sprint(companyUser.UserId)] =*userProfile
			}
		}

		return c.OK(users)
	}

	return c.Err("没有权限")
}

func (c CompanyUsers) Add(companyId, userId int) revel.Result {
	isOwner, _ := lib.IsCompanyOwner(nil, companyId, c.User.Id)
	if !isOwner {
		return c.Err("没有权限")
	}

	has := lib.IsCompanyUser(companyId, userId)
	if has {
		return c.Err("用户已存在")
	}

	now := time.Now()
	cu := &models.CompanyUsers{
		CompanyId: companyId,
		UserId:userId,
		Status:utils.COMPANY_USER_STATUS_CHECK_NO,
		UpdatedAt:now,
		CreatedAt:now,
	}
	_, err := app.Engine.Insert(cu)
	if err != nil {
		return c.Err("添加用户失败")
	}

	c.updateUserCheckCount(companyId,false,true)
	return c.OK(nil)
}
func (c CompanyUsers) AddByCompanyName(companyName string, userId int) revel.Result {

	if len(companyName) == 0{
		return c.Err("请输入完整的单位名称，不能为空")
	}
	company := &models.Company{}
	has, err := app.Engine.Where("name = ?",companyName).Get(company)
	if err != nil{
		return c.Err(err.Error())
	}

	if !has{
		return c.Err("单位不存在，请输入完整的单位名称")
	}

	return c.Add(company.Id,userId);
}


func (c CompanyUsers) Check(companyId, userId int) revel.Result {
	isOk, message := c.changeStatus(companyId, userId, utils.COMPANY_USER_STATUS_CHECK_YES)
	if !isOk {
		return c.Err(message)
	}
	c.updateUserCheckCount(companyId,true,true)
	return c.OK(nil)
}

func (c CompanyUsers) Delete(companyId, userId int) revel.Result {
	isOk, message := c.changeStatus(companyId, userId, utils.COMPANY_USER_STATUS_DELETE)
	if !isOk {
		return c.Err(message)
	}
	c.updateUserCheckCount(companyId,true,true)
	return c.OK(nil)
}

func (c CompanyUsers) changeStatus(companyId, userId, status int) (isOk bool, message string) {
	isOk = false
	message = ""

	isOwner, _ := lib.IsCompanyOwner(nil, companyId, c.User.Id)
	if !isOwner {
		message = "没有权限"
		return
	}

	has := lib.IsCompanyUser(companyId, userId)
	if !has {
		message = "该用户未申请"
		return
	}

	_, err := app.Engine.Where("company_id =  ? and user_id = ?", companyId, userId).Cols("status").Update(&models.CompanyUsers{Status:status})

	if err != nil {
		message = err.Error()
		return
	}

	isOk = true
	return
}

func (c *CompanyUsers) updateUserCheckCount(companyId int,isUpdateCheckedCount,isUpdateUnCheckedCount bool) {

	checkedCount := 0;
	if isUpdateCheckedCount{
		user := new(models.CompanyUsers)
		total, err := app.Engine.Where("company_id =  ? and status = ?", companyId,utils.COMPANY_USER_STATUS_CHECK_YES).Count(user)
		if err == nil{
			checkedCount = int(total)
		}else{
			isUpdateCheckedCount = false
		}
	}

	unCheckedCount := 0
	if isUpdateUnCheckedCount{
		user := new(models.CompanyUsers)
		total, err := app.Engine.Where("company_id =  ? and status = ?", companyId,utils.COMPANY_USER_STATUS_CHECK_NO).Count(user)
		if err == nil{
			unCheckedCount = int(total)
		}else{
			isUpdateUnCheckedCount = false
		}
	}

	session := app.Engine.Id(companyId)

	cmp := &models.Company{}
	if isUpdateCheckedCount {
		cmp.UserCheckCount = checkedCount;
		session.Cols("user_check_count")
	}
	if isUpdateUnCheckedCount {
		cmp.UserUncheckCount = unCheckedCount;
		session.Cols("user_uncheck_count")
	}

	session.Update(cmp)
}