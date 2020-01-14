package db

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math"
	"strconv"
)

type IDao interface {
	FindById(id uint, columns []string) error
	Create(data interface{}) error
	Update(id uint, data map[string]interface{}) error
	Destroy(id uint) error
	Query(params map[string]interface{}, columns []string, with map[string]interface{}) IDao
	Find(model interface{}) error
	First(model interface{}) error
	Paginator(model interface{}, p *Paginator) error
}

type Dao struct {
	M  interface{}
	DB *gorm.DB
}

func (d *Dao) Find(model interface{}) error {
	if d.DB == nil {
		return ORM.Find(model).Error
	}
	return d.DB.Find(model).Error
}

func (d *Dao) First(model interface{}) error {
	if d.DB == nil {
		return ORM.First(model).Error
	}
	return d.DB.First(model).Error
}

func (d *Dao) FindById(id uint, columns []string) error {
	return ORM.Select(columns).Where("id = ?", id).First(d.M).Error
}

func (d *Dao) Create(data interface{}) error {
	return ORM.Create(data).Error
}

func (d *Dao) Update(id uint, data map[string]interface{}) error {
	err := d.FindById(id, []string{"id"})
	if err != nil {
		return err
	}
	return ORM.Model(d.M).Where("id = ?", id).Update(data).Error
}

func (d *Dao) Destroy(id uint) error {
	return ORM.Where("id = ?", id).Delete(d.M).Error
}

// 分页管理器
type Paginator struct {
	CurrentPage int `json:"current_page"` // 当前页
	PageSize    int `json:"page_size"`    // 每页数量
	TotalPage   int `json:"total_page"`   // 总页数
	TotalCount  int `json:"total_count"`  // 总数量
}

// 新建一个分页管理器
// 获取请求的 page, limit 参数
// 返回分页管理器
func NewPaginator(ctx *gin.Context, defaultPage, defaultLimit int) Paginator {
	page := ctx.DefaultQuery("page", strconv.Itoa(defaultPage))
	limit := ctx.DefaultQuery("limit", strconv.Itoa(defaultLimit))
	currentPage, err := strconv.Atoi(page)
	if err != nil {
		currentPage = defaultPage
	}
	pageSize, err := strconv.Atoi(limit)
	if err != nil {
		pageSize = defaultLimit
	}
	// 分页
	return Paginator{
		CurrentPage: currentPage,
		PageSize:    pageSize,
	}
}

// 分页处理器
func (d *Dao) Paginator(model interface{}, p *Paginator) error {
	var (
		err   error
		query *gorm.DB
		count int
	)

	if query = d.DB; query == nil {
		query = ORM
	}

	if err = query.Model(model).Count(&count).Error; err != nil {
		return err
	}

	query = query.Limit(p.PageSize).Offset((p.CurrentPage - 1) * p.PageSize)
	if err = query.Find(model).Error; err != nil {
		return err
	}

	p.TotalCount = count
	p.TotalPage = int(math.Ceil(float64(count) / float64(p.PageSize)))
	return nil
}