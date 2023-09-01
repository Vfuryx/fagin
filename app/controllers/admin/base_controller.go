package admin

import (
	"fadmin/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// BaseController 基础控制器
type BaseController struct {
}

const DefaultLimit = 15

func (ctr BaseController) ResponseJSONOK(ctx *fiber.Ctx, data interface{}) error {
	return response.JSONSuccess(ctx, data)
}

func (ctr BaseController) ResponseJSONErr(ctx *fiber.Ctx, err error, errors interface{}) error {
	return response.JSONErr(ctx, err, errors)
}
