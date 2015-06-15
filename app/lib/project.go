package lib

import (
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app"
)
func IsProjectBelongToCompany(project *models.Project, id ,companyId int) (bool,*models.Project){
	if nil == project{
		project = &models.Project{}
		_, err := app.Engine.Id(id).Get(project)
		if nil==project || err != nil{
			return false,project
		}
	}

	if project.CompanyId != companyId{
		return false,project
	}
	return true,project
}