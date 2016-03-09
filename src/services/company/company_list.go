package company

import (
	"models"
)

type CompanyList struct {
	companies []*models.Company
}

func NewCompanyList() *CompanyList {
	companyList := CompanyList{}
	companyList.do()

	return &companyList
}

func (companyList *CompanyList) do() *CompanyList {
	var err error
	_, err = models.GetDB().QueryTable("companies").All(&companyList.companies)
	if err != nil {
		return companyList
	}

	return companyList
}

func (companyList *CompanyList) GetList() []*models.Company {
	return companyList.companies
}
