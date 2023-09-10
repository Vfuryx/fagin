package services

import (
	"fadmin/app/models/admin/domain/entities/user"
	"fadmin/app/models/admin/domain/services"
	"fadmin/app/models/admin/repositories"
)

type AdminUserService struct{}

func (AdminUserService) UserAdd(username string, nickName string) error {
	return services.UserAdd(username, nickName)
}

func (AdminUserService) Find(id uint) (*user.AdminUser, error) {
	return repositories.User().Find(id)
}
