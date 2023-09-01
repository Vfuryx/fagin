package searches

import (
	"fadmin/app/models/admin/domain/entities/user"
	"fadmin/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type UserSearch struct {
	id   *uint
	name *string

	page      int
	pageSize  int
	paginator *db.Paginator
	res       []int
}

func (s *UserSearch) Search() ([]uint, error) {
	query := s.getQuery()

	s.paginator = db.NewPaginator(s.page, s.pageSize, 10)
	var userItems []user.AdminUser
	err := s.paginator.Paginate(query, userItems)
	if err != nil {
		return nil, err
	}

	var res []uint
	for _, item := range userItems {
		res = append(res, item.ID)
	}

	return res, nil
}

func (s *UserSearch) GetPagination() *db.Paginator {
	return s.paginator
}

func (s *UserSearch) getQuery() *gorm.DB {
	columns := []string{"id"}
	model := db.ORM().Select(columns)

	if s.name != nil {
		model = model.Where("username = ?", *s.name)
	}

	if s.id != nil {
		model = model.Where("id = ?", *s.id)
	}

	return model
}

func (s *UserSearch) Load(ctx *fiber.Ctx) *UserSearch {
	params := ctx.AllParams()
	if v, ok := params["page"]; ok {
		s.setPage(v)
	}
	if v, ok := params["page_size"]; ok {
		s.setPageSize(v)
	}
	if v, ok := params["name"]; ok {
		s.setName(v)
	}
	if v, ok := params["id"]; ok {
		s.setId(v)
	}

	return s
}

func (s *UserSearch) setName(name interface{}) {
	var v = cast.ToString(name)
	s.name = &v
}

func (s *UserSearch) setId(id interface{}) {
	var v = cast.ToUint(id)
	s.id = &v
}

func (s *UserSearch) setPage(page interface{}) {
	s.page = cast.ToInt(page)
}

func (s *UserSearch) setPageSize(pageSize interface{}) {
	s.pageSize = cast.ToInt(pageSize)
}
