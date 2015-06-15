package lib

import (
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app"
)
func IsTaskBelongToCompany(task *models.Task, id ,companyId int) (bool,*models.Task){
	//fmt.Println("----- id " + strconv.Itoa(id) + " --- userId " + strconv.Itoa(userId) )
	if nil == task{
		task = &models.Task{}
		_, err := app.Engine.Id(id).Get(task)
		if nil==task || err != nil{
			return false,task
		}
	}

	if task.CompanyId != companyId{
		return false,task
	}
	return true,task
}