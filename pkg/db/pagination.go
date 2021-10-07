package db

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
)

// Paginator 分页管理器
type Paginator struct {
	CurrentPage int   `json:"current_page"` // 当前页
	PageSize    int   `json:"page_size"`    // 每页数量
	TotalPage   int   `json:"total_page"`   // 总页数
	TotalCount  int64 `json:"total_count"`  // 总数量
}

// NewPaginatorWithCtx 新建一个分页管理器
// 获取请求的 page, limit 参数
// 返回分页管理器
func NewPaginatorWithCtx(ctx *gin.Context, defaultPage, defaultLimit int) *Paginator {
	page := ctx.DefaultQuery("page", strconv.Itoa(defaultPage))
	limit := ctx.DefaultQuery("pageSize", strconv.Itoa(defaultLimit))
	currentPage, err := strconv.Atoi(page)
	if err != nil {
		currentPage = defaultPage
	}
	pageSize, err := strconv.Atoi(limit)
	if err != nil {
		pageSize = defaultLimit
	}
	// 分页
	return NewPaginator(currentPage, pageSize)
}

// NewPaginator 新建一个分页管理器
// 获取请求的 page, limit 参数
// 返回分页管理器
func NewPaginator(currentPage, pageSize int) *Paginator {
	// 分页
	return &Paginator{
		CurrentPage: currentPage,
		PageSize:    pageSize,
	}
}

// Paginate 分页
func (p *Paginator) Paginate(query *gorm.DB, model interface{}) (err error) {
	var count int64
	if query == nil {
		query = ORM()
	}
	if err = query.Model(model).Select([]string{}).Count(&count).Error; err != nil {
		return err
	}
	query = query.Limit(p.PageSize).Offset(getPageOffset(p.CurrentPage, p.PageSize))
	if err = query.Find(model).Error; err != nil {
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
