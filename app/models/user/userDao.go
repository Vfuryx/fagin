package user

import "fagin/app/models"

func Create(user *User) (ok bool, err error) {

	if err := models.ORM.Model(&User{}).Create(user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func Update(user *User) (ok bool, err error) {

	if err = models.ORM.Model(&User{}).Update(user).Error; err != nil{
		return false, err
	}

	return true, nil
}