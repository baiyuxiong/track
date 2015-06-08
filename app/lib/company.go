package lib

import (
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app"
	"fmt"
	"strconv"
)
func IsCompanyOwner(company *models.Company, id ,userId int) (bool,*models.Company){
	fmt.Println("----- id " + strconv.Itoa(id) + " --- userId " + strconv.Itoa(userId) )
	if nil == company{
		company = &models.Company{}
		_, err := app.Engine.Id(id).Get(company)
		if nil==company || err != nil{
			return false,company
		}
	}

	if company.OwnerId != userId{
		return false,company
	}
	return true,company
}