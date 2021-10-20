package service

import (
	"fagin/app/models/admin_department"
	"fagin/pkg/errorw"
	"github.com/gin-gonic/gin"
)

type adminDepartmentService struct{}

var AdminDepartment adminDepartmentService

func (*adminDepartmentService) Index(params gin.H, columns []string, with gin.H) (ms []admin_department.AdminDepartment, err error) {
	err = admin_department.NewDao().Query(params, columns, with).Find(&ms)
	return ms, errorw.UP(err)
}

func (*adminDepartmentService) Show(id uint, columns []string) (*admin_department.AdminDepartment, error) {
	b := admin_department.New()
	err := b.Dao().FindById(id, columns)
	return b, errorw.UP(err)
}

func (*adminDepartmentService) Create(m *admin_department.AdminDepartment) error {
	err := admin_department.NewDao().Create(m)
	return errorw.UP(err)
}

func (*adminDepartmentService) Update(id uint, data gin.H) error {
	err := admin_department.NewDao().Update(id, data)
	return errorw.UP(err)
}

func (*adminDepartmentService) Delete(id uint) error {
	err := admin_department.NewDao().Destroy(id)
	return errorw.UP(err)
}

func (*adminDepartmentService) Deletes(ids []uint) error {
	err := admin_department.NewDao().Deletes(ids)
	return errorw.UP(err)
}
