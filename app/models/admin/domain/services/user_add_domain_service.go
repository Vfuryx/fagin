package services

import (
	"fadmin/app/models/admin/domain/entities/user"
	"fadmin/pkg/db"
)

func UserAdd(username string, nickName string) error {
	var err error
	u := new(user.AdminUser)
	err = u.Add(username, nickName)
	if err != nil {
		return err
	}
	err = db.ORM().Create(u).Error
	if err != nil {
		return err
	}

	return err
}
