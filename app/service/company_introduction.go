package service

import (
	ci "fagin/app/models/company_introduction"
	"fagin/pkg/errorw"
)

type companyIntroduction struct{}

// CompanyIntroduction
var CompanyIntroduction companyIntroduction

func (*companyIntroduction) ShowCompanyIntroduction(id uint, columns []string) (*ci.CompanyIntroduction, error) {
	introduction := ci.New()
	err := introduction.Dao().FindById(id, columns)
	return introduction, errorw.UP(err)
}

func (*companyIntroduction) UpdateCompanyIntroduction(id uint, data map[string]interface{}) error {
	return errorw.UP(ci.NewDao().Update(id, data))
}
