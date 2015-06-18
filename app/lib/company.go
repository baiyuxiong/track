package lib

import (
	"github.com/baiyuxiong/track/app/models"
	"github.com/baiyuxiong/track/app/utils"
	"github.com/baiyuxiong/track/app"
	"strings"
	"fmt"
)
func IsCompanyOwner(company *models.Company, id, userId int) (bool, *models.Company) {
	//fmt.Println("----- id " + strconv.Itoa(id) + " --- userId " + strconv.Itoa(userId) )
	if nil == company {
		company = &models.Company{}
		_, err := app.Engine.Id(id).Get(company)
		if nil==company || err != nil {
			return false, company
		}
	}

	if company.OwnerId != userId {
		return false, company
	}
	return true, company
}

type MyCompanies struct {
	CompaniesOwned   []models.Company    `json:"companiesOwned"`
	CompaniesJoined  []models.Company    `json:"companiesJoined"`
	CompaniesJoining []models.Company    `json:"companiesJoining"`
}

//获取创建的、加入的、等待审核的团队列表
func GetMyCompanies(userId int) *MyCompanies {
	companiesOwned := make([]models.Company, 0);
	err := app.Engine.Where("owner_id = ?", userId).OrderBy("updated_at desc").Find(&companiesOwned)
	if err != nil {
		return nil
	}

	ownedCompanyIds := make([]string, 0)
	for _, company := range companiesOwned {
		ownedCompanyIds = append(ownedCompanyIds, fmt.Sprintf("%d", company.Id))
	}

	//.Where("company_id not in ?", "("+strings.Join(ownedCompanyIds, ",")+")")
	companyUsers := make([]models.CompanyUsers, 0)
	err = app.Engine.Where("user_id = ? and status <> ? and company_id not in (?)",userId, utils.COMPANY_USER_STATUS_DELETE,strings.Join(ownedCompanyIds, ",")).OrderBy("updated_at desc").Find(&companyUsers)
	if err != nil {
		return nil
	}

	//获取团队ID列表
	joinedCompanyIds := make([]string, 0)
	joiningCompanyIds := make([]string, 0)
	for _, companyUser := range companyUsers {
		if companyUser.Status == utils.COMPANY_USER_STATUS_CHECK_NO {
			joiningCompanyIds = append(joiningCompanyIds, fmt.Sprintf("%d", companyUser.CompanyId))
		} else if companyUser.Status == utils.COMPANY_USER_STATUS_CHECK_YES {
			joinedCompanyIds = append(joinedCompanyIds, fmt.Sprintf("%d", companyUser.CompanyId))
		}
	}

	//根据ID查询团队信息
	companiesJoined := make([]models.Company, 0);
	companiesJoining := make([]models.Company, 0);
	for companyId := range joinedCompanyIds {
		company := &models.Company{}
		has,_ := app.Engine.Id(companyId).Get(company)
		if has{
			companiesJoined = append(companiesJoined, *company)
		}
	}

	for companyId := range joiningCompanyIds {
		company := &models.Company{}
		has,_ := app.Engine.Id(companyId).Get(company)
		if has {
			companiesJoining = append(companiesJoining, *company)
		}
	}

	return &MyCompanies{
		CompaniesOwned:companiesOwned,
		CompaniesJoined:companiesJoined,
		CompaniesJoining:companiesJoining,
	};
}