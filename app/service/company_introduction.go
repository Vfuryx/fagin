package service

import (
	ci "fagin/app/models/company_introduction"
)

type companyIntroduction struct{}

var CompanyIntroduction companyIntroduction

func (*companyIntroduction) ShowCompanyIntroduction(id uint, columns []string) (*ci.CompanyIntroduction, error) {
	introduction := ci.New()
	err := introduction.Dao().FindById(id, columns)
	return introduction, err
}

func (*companyIntroduction) UpdateCompanyIntroduction(id uint, data map[string]interface{}) error {
	return ci.NewDao().Update(id, data)
}
