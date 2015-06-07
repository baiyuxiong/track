package lib

import (
	"github.com/baiyuxiong/track/app"
	"github.com/baiyuxiong/track/app/models"
)

func IsCompanyUser(company_id ,user_id int) bool {
	var s = &models.CompanyUsers{CompanyId: company_id,UserId:user_id}
	has, _ := app.Engine.Get(s)
	return has
}
