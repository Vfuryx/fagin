package services

import "fadmin/app/models/admin/domain/services"

type adminUserService struct {
}

func (adminUserService) UserAdd(username string, nickName string) error {
	return services.UserAdd(username, nickName)
}
