package admin_menu

import (
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func New() *AdminMenu {
	return &AdminMenu{}
}

type dao struct {
	db.Dao
}

var _ db.IDao = &dao{}

func (m *AdminMenu) Dao() *dao {
	dao := &dao{}
	dao.Dao.M = m
	return dao
}

func Dao() *dao {
	dao := &dao{}
	dao.Dao.M = New()
	return dao
}

func (dao) All(params gin.H, columns []string) (*[]AdminMenu, error) {
	var model []AdminMenu
	m := db.ORM().Select(columns)
	if v, ok := params["type"]; ok {
		m = m.Where("type = ?", v)
	}
	err := m.Find(&model).Error
	return &model, err
}

func (d *dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.IDao {
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

	if v, ok = params["type"]; ok {
		model = model.Where("type = ?", v)
	}

	if v, ok = params["in_type"]; ok {
		model = model.Where("type in (?)", v)
	}

	if v, ok = params["like_title"]; ok && v.(string) != "" {
		model = model.Where("`title` LIKE ?", v)
	}

	if v, ok = params["like_paths"]; ok && v.(string) != "" {
		model = model.Where("`paths` LIKE ?", v)
	}

	if v, ok = params["is_show"]; ok {
		model = model.Where("is_show = ?", v)
	}

	if v, ok = params["path"]; ok {
		model = model.Where("path = ?", v)
	}

	if v, ok = params["method"]; ok {
		model = model.Where("method = ?", v)
	}

	if v, ok = params["orderBy"]; ok {
		model = model.Order(v)
	}
	if v, ok = params["status"]; ok {
		model = model.Where("status = ?", v)
	}

	d.DB = d.With(model, with)
	return d
}

func (d *dao) Update(id uint, data map[string]interface{}) error {
	var m AdminMenu
	err := db.ORM().Select([]string{"id", "parent_id", "paths"}).Where("id = ?", id).First(&m).Error
	if err != nil {
		return err
	}

	if v, ok := data["parent_id"]; ok {
		if v != m.ParentId {
			paths := ""
			if v.(uint) == 0 {
				paths = "0-" + strconv.FormatUint(uint64(m.ID), 10)
			} else {
				var pm AdminMenu
				err = db.ORM().Select([]string{"id", "parent_id", "paths"}).Where("id = ?", v).First(&pm).Error
				paths = pm.Paths + "-" + strconv.FormatUint(uint64(m.ID), 10)
			}
			var mm []AdminMenu
			err := db.ORM().Where(`paths like ?`, m.Paths+"%").Find(&mm).Error
			if err != nil {
				return err
			}
			for i := range mm {
				mm[i].Paths = strings.Replace(mm[i].Paths, m.Paths, paths, -1)
				err = db.ORM().Model(&AdminMenu{}).Where("id = ?", mm[i].ID).Updates(mm[i]).Error
				if err != nil {
					return err
				}
			}
		}
	}

	return db.ORM().Model(d.M).Where("id = ?", id).Updates(data).Error
}

func (d *dao) Delete(id uint) error {
	var m AdminMenu
	err := db.ORM().Select([]string{"id", "paths"}).Where("id = ?", id).First(&m).Error
	if err != nil {
		return err
	}

	return db.ORM().Where(`paths like ?`, m.Paths+"%").Delete(&AdminMenu{}).Error
}
