package user

import "github.com/Vfuryx/fagin/database"

func Create(user *User) (ok bool, err error) {

	if err := database.ORM.Model(&User{}).Create(user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func Update(user *User) (ok bool, err error) {

	if err = database.ORM.Model(&User{}).Update(user).Error; err != nil{
		return false, err
	}

	return true, nil
}