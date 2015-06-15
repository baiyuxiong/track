package lib

import (
	"github.com/baiyuxiong/track/app"
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
)

func IsCompanyUser(companyId, userId int) bool {
	var s = &models.CompanyUsers{CompanyId: companyId,UserId:userId}
	has, _ := app.Engine.Get(s)
	return has
}

func IsCompanyCheckedUser(companyId ,userId int) bool {
	var s = &models.CompanyUsers{CompanyId: companyId,UserId:userId,Status:utils.COMPANY_USER_STATUS_CHECK_YES}
	has, _ := app.Engine.Get(s)
	return has
}
