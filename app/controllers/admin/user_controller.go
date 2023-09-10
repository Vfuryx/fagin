package admin

import (
	"fadmin/app/controllers/admin/transform/user_transform"
	"fadmin/app/models/admin/searches"
	adminServices "fadmin/app/models/admin/services"
	"fadmin/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	BaseController
}

var UserController userController

func (ctr *userController) Index(ctx *fiber.Ctx) error {
	logger.Log().Error("test index err")

	search := new(searches.UserSearch)
	uints, err := search.Load(ctx).Search()
	if err != nil {
		return ctr.ResponseJSONErr(ctx, err, nil)
	}

	transform, err := user_transform.AdminUserListTransform(uints, search.GetPagination())
	if err != nil {
		return ctr.ResponseJSONErr(ctx, err, nil)
	}

	return ctr.ResponseJSONOK(ctx, transform)
}

func (ctr *userController) Add(ctx *fiber.Ctx) error {
	logger.Log().Error("test index err")

	err := adminServices.User().UserAdd("12", "1212")
	if err != nil {
		return err
	}

	return ctr.ResponseJSONOK(ctx, fiber.Map{"m": "index"})
}
