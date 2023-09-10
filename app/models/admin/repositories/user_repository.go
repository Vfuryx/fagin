package repositories

import (
	"fadmin/app/models/admin/domain/entities/user"
	"fadmin/pkg/db"
)

type UserRepository struct{}

func (UserRepository) Find(id uint) (*user.AdminUser, error) {
	columns := []string{"*"}
	model := new(user.AdminUser)
	err := db.ORM().Select(columns).Where("id = ?", id).Find(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (UserRepository) FindItems(ids []uint) ([]user.AdminUser, error) {
	columns := []string{"*"}
	var model []user.AdminUser
	err := db.ORM().Select(columns).Where("id = ?", ids).Find(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}
