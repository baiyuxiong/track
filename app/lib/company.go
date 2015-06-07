package lib

import (
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app"
)
func IsCompanyOwner(company models.Company, id ,user_id int) (bool,models.Company){
	if nil == company{
		company := &models.Company{}
		_, err := app.Engine.Id(id).Get(company)
		if err != nil{
			return false,company
		}
	}

	if company.OwnerId != user_id{
		return false,company
	}
	return true,company
}