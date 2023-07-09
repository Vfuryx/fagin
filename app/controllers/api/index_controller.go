package api

import (
	"fadmin/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type indexController struct {
	BaseController
}

var IndexController indexController

func (ctr *indexController) Index(ctx *fiber.Ctx) error {
	logger.Log().Error("test index err")

	return ctr.ResponseJSONOK(ctx, fiber.Map{"m": "index"})
}
