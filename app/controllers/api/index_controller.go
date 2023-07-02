package api

import (
	"fadmin/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type indexController struct {
}

var IndexController indexController

func (ctr *indexController) Index(ctx *fiber.Ctx) error {
	logger.Log().Error("test index err")
	return ctx.JSON(fiber.Map{"m": "index"})
}
