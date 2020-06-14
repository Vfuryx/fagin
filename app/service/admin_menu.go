package service

import (
	"fagin/app/models/admin_menu"
	"github.com/gin-gonic/gin"
	"strconv"
)

type adminMenuService struct{}

var AdminMenuService adminMenuService

func (adminMenuService) Index(params gin.H, columns []string, with gin.H) ([]admin_menu.AdminMenu, error) {
	var menus []admin_menu.AdminMenu

	err := admin_menu.Dao().Query(params, columns, with).Find(&menus)
	if err != nil {
		return nil, err
	}

	// tree
	menus = getMenuTree(menus, 0)
	return menus, err
}

func getMenuTree(data []admin_menu.AdminMenu, pID uint) []admin_menu.AdminMenu {
	menus := make([]admin_menu.AdminMenu, 0, 20)
	for _, m := range data {
		if m.ParentId == pID {
			mc := getMenuTree(data, m.ID)
			m.Children = mc
			menus = append(menus, m)
		}
	}
	return menus
}

func (adminMenuService) Show(id uint, columns []string) (*admin_menu.AdminMenu, error) {
	m := admin_menu.New()
	err := m.Dao().FindById(id, columns)
	return m, err
}

func (adminMenuService) Create(m *admin_menu.AdminMenu) error {
	err := admin_menu.Dao().Create(m)
	if err != nil {
		return err
	}
	// 设置路径
	return setPaths(m)
}

// 设置路径
func setPaths(menus *admin_menu.AdminMenu) error {
	adminMenu := admin_menu.New()

	paths := ""

	// 判断父ID是否为0
	if menus.ParentId == 0 {
		paths = "0-" + strconv.FormatUint(uint64(menus.ID), 10)
	}else {
		err := adminMenu.Dao().FindById(menus.ParentId, []string{"id", "paths"})
		if err != nil {
			return err
		}
		paths = adminMenu.Paths + "-" + strconv.FormatUint(uint64(menus.ID), 10)
	}

	return menus.Dao().Update(menus.ID, gin.H{"paths": paths})
}

func (adminMenuService) Update(id uint, data gin.H) error {
	return admin_menu.Dao().Update(id, data)
}


func (adminMenuService) Delete(id uint) error {
	return admin_menu.Dao().Delete(id)
}
