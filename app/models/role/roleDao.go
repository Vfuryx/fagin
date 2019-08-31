package role

import (
	"fagin/pkg/db"
)

func Create(role *Role) (ok bool, err error) {

	if err := db.ORM.Model(&Role{}).Create(role).Error; err != nil {
		return false, err
	}

	return true, nil
}

func Update(role *Role) (ok bool, err error) {

	if err := db.ORM.Model(&Role{}).Update(role).Error; err != nil {
		return false, err
	}

	return true, nil
}
