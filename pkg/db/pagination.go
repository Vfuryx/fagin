package db

import (
	"gorm.io/gorm"
	"math"
)

// Paginator 分页管理器
type Paginator struct {
	CurrentPage int   `json:"current_page"` // 当前页
	PageSize    int   `json:"page_size"`    // 每页数量
	TotalPage   int   `json:"total_page"`   // 总页数
	TotalCount  int64 `json:"total_count"`  // 总数量
}

// NewPaginator 新建一个分页管理器
// 获取请求的 page, limit 参数
// 返回分页管理器
func NewPaginator(currentPage, pageSize, defPageSize int) *Paginator {
	if currentPage <= 0 {
		currentPage = 1
	}
	if pageSize <= 0 {
		pageSize = defPageSize
	}

	// 分页
	return &Paginator{
		CurrentPage: currentPage,
		PageSize:    pageSize,
	}
}

// Paginate 分页
func (p *Paginator) Paginate(query *gorm.DB, model interface{}) error {
	var (
		count int64
	)

	if query == nil {
		query = ORM()
	}

	if err := query.Model(model).Select([]string{}).Count(&count).Error; err != nil {
		return err
	}

	query = query.Limit(p.PageSize).Offset(getPageOffset(p.CurrentPage, p.PageSize))

	if err := query.Find(model).Error; err != nil {
		return err
	}

	p.TotalCount = count
	p.TotalPage = int(math.Ceil(float64(count) / float64(p.PageSize)))

	return nil
}

func getPageOffset(currentPage, pageSize int) int {
	result := 0
	if currentPage > 0 {
		result = (currentPage - 1) * pageSize
	}

	return result
}
