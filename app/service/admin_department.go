package service

import (
	"fagin/app/models/admin_department"
	"github.com/gin-gonic/gin"
)

type adminDepartmentService struct{}

var AdminDepartment adminDepartmentService

func (*adminDepartmentService) Index(params gin.H, columns []string, with gin.H) (ms []admin_department.AdminDepartment, err error) {
	err = admin_department.NewDao().Query(params, columns, with).Find(&ms)
	return
}

func (*adminDepartmentService) Show(id uint, columns []string) (*admin_department.AdminDepartment, error) {
	b := admin_department.New()
	err := b.Dao().FindById(id, columns)
	return b, err
}

func (*adminDepartmentService) Create(m *admin_department.AdminDepartment) error {
	return admin_department.NewDao().Create(m)
}

func (*adminDepartmentService) Update(id uint, data gin.H) error {
	return admin_department.NewDao().Update(id, data)
}

func (*adminDepartmentService) Delete(id uint) error {
	return admin_department.NewDao().Destroy(id)
}

func (*adminDepartmentService) Deletes(ids []uint) error {
	return admin_department.NewDao().Deletes(ids)
}
