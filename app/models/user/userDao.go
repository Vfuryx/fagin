package user

import (
	"fagin/pkg/db"
)

func Create(user *User) (ok bool, err error) {

	if err := db.ORM.Model(&User{}).Create(user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func Update(user *User) (ok bool, err error) {

	if err := db.ORM.Model(&User{}).Update(user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func GetUserByUsername(username string) (User, error) {
	var user User
	err := db.ORM.Where(&User{Username: username}).First(&user).Error
	return user, err
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := db.ORM.Where(&User{Email: email}).First(&user).Error
	return user, err
}

//func Login(username string, password string) (err exception) {
//	var user User
//	return models.ORM.Where(&User{Username:username,Password:password}).First(&user).Error
//}
