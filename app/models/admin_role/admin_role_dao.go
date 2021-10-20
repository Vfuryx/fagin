package admin_role

import (
	"fagin/app/models/admin_menu"
	"fagin/app/models/admin_role_menu"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func New() *AdminRole {
	return &AdminRole{}
}

type Dao struct {
	db.Dao
}

var _ db.DAO = &Dao{}

func (m *AdminRole) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

// All 所有
func (d *Dao) All(columns []string) (*[]AdminRole, error) {
	var model []AdminRole
	err := db.ORM().Select(columns).Find(&model).Error
	return &model, err
}

// Query 查询
func (d *Dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.DAO {
	model := db.ORM().Select(columns)

	var (
		v  interface{}
		ok bool
	)
	if v, ok = params["id"]; ok {
		model = model.Where("id = ?", v)
	}

	if v, ok = params["in_id"]; ok {
		model = model.Where("id in (?)", v)
	}

	if v, ok = params["like_name"]; ok && v.(string) != "" {
		model = model.Where("`name` LIKE ?", v)
	}

	if v, ok = params["like_key"]; ok && v.(string) != "" {
		model = model.Where("`key` LIKE ?", v)
	}

	if v, ok = params["key"]; ok && v.(string) != "" {
		model = model.Where("`key` = ?", v)
	}

	if v, ok = params["status"]; ok {
		model = model.Where("`status` = ?", v)
	}

	if v, ok = params["type"]; ok {
		model = model.Where("`type` = ?", v)
	}

	if v, ok = params["orderBy"]; ok {
		model = model.Order(v)
	}

	d.DB = d.With(model, with)
	return d
}

// Show 展示
func (d *Dao) Show(id uint, columns []string) (*AdminRole, error) {
	var role AdminRole
	err := db.ORM().
		Select(columns).
		Where("id = ?", id).
		Preload("Menus").
		First(&role).Error
	return &role, err
}

// Update 更新
func (d *Dao) Update(id uint, data map[string]interface{}) error {
	err := db.ORM().Model(d.GetModel()).Omit("menuIDs", "permissionsIDs").Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}

	// 获取菜单组
	var menus []admin_menu.AdminMenu
	if v, ok := data["menuIDs"]; ok {
		ids := v.([]uint)
		err = admin_menu.NewDao().Query(gin.H{"in_id": ids}, []string{"id", "paths"}, nil).Find(&menus)
		if err != nil {
			return err
		}
		delete(data, "menuIDs")
	}
	menus, err = GetMenuAll(menus)
	if err != nil {
		return err
	}
	err = db.ORM().Model(&AdminRole{ID: id}).Association("Menus").Replace(menus)
	return err
}

func GetMenuAll(menus []admin_menu.AdminMenu) ([]admin_menu.AdminMenu, error) {
	var mIDsExists = make(map[int]struct{})
	var mIDs = make(map[int]struct{})
	var err error

	var i int
	var idSlice []string
	var s string
	var ok, ok2 bool
	for _, menu := range menus {
		mIDsExists[int(menu.ID)] = struct{}{}
		idSlice = strings.Split(menu.Paths, "-")
		idSlice = idSlice[1 : len(idSlice)-1]
		for _, s = range idSlice {
			i, err = strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			_, ok = mIDs[i]
			_, ok2 = mIDsExists[i]
			if !ok && !ok2 {
				mIDs[i] = struct{}{}
			}
		}
	}
	var ids []int
	for i = range mIDs {
		if i != 0 {
			ids = append(ids, i)
		}
	}
	var menus2 []admin_menu.AdminMenu
	err = admin_menu.NewDao().Query(gin.H{"in_id": ids}, []string{"id"}, nil).Find(&menus2)
	if err != nil {
		return nil, err
	}
	// 追加父级菜单
	menus = append(menus, menus2...)
	return menus, nil
}

// Delete 删除
func (d *Dao) Delete(id uint) error {
	var role AdminRole
	err := db.ORM().Where("id = ?", id).Delete(&role).Error
	if err != nil {
		return err
	}
	var roleMenu admin_role_menu.AdminRoleMenu
	return db.ORM().Where("role_id = ?", id).Delete(&roleMenu).Error
}

// Deletes 批量删除
func (d *Dao) Deletes(ids []uint) error {
	var role AdminRole
	err := db.ORM().Where("id in (?)", ids).Delete(&role).Error
	if err != nil {
		return err
	}
	var roleMenu admin_role_menu.AdminRoleMenu
	return db.ORM().Where("role_id in (?)", ids).Delete(&roleMenu).Error
}
