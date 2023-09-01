package admin

import (
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
		return err
	}
	search.GetPagination()
	return ctr.ResponseJSONOK(ctx, fiber.Map{"m": uints})
}

func (ctr *userController) Add(ctx *fiber.Ctx) error {
	logger.Log().Error("test index err")

	err := adminServices.User().UserAdd("12", "1212")
	if err != nil {
		return err
	}

	return ctr.ResponseJSONOK(ctx, fiber.Map{"m": "index"})
}
