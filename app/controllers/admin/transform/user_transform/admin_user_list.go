package user_transform

import (
	"fadmin/app/models/admin/repositories"
	"fadmin/pkg/db"
	"fadmin/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func AdminUserListTransform(ids []uint, p *db.Paginator) (fiber.Map, error) {
	sm := make([]map[string]interface{}, 0, response.DefCap)
	items, err := repositories.User().FindItems(ids)
	if err != nil {
		return nil, err
	}

	for _, u := range items {
		m := map[string]interface{}{
			"id":        u.ID,
			"username":  u.Username,
			"nick_name": u.NickName,
			"email":     u.Email,
			"remark":    u.Remark,
			"phone":     u.Phone,
			"status":    u.Status,
			"create_at": u.CreatedAt,
		}
		sm = append(sm, m)
	}

	res := fiber.Map{
		"items": sm,
		"pagination": fiber.Map{
			"page":        p.CurrentPage,
			"page_size":   p.PageSize,
			"page_count":  p.TotalPage,
			"total_count": p.TotalCount,
		},
	}

	return res, nil
}
