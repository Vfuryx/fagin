package response

import (
	errNo "fadmin/app/errno"
	"fadmin/pkg/errno"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

const DefCap = 20

type Response interface {
	Serialize() []map[string]interface{}
}

type Collect[T Response] struct {
	Response T
}

func (c *Collect[T]) SetCollect(res T) *Collect[T] {
	c.Response = res

	return c
}

func (c Collect[T]) Handle() []map[string]interface{} {
	return c.Response.Serialize()
}

func (c Collect[T]) Item() map[string]interface{} {
	return c.Handle()[0]
}

func (c Collect[T]) Collection() []map[string]interface{} {
	return c.Handle()
}

type Model struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func JSONSuccess(ctx *fiber.Ctx, result interface{}) error {
	return JSON(ctx, errNo.OK, result, nil, http.StatusOK)
}

func JSONErr(ctx *fiber.Ctx, err error, errors interface{}) error {
	return JSON(ctx, err, nil, errors, http.StatusOK)
}

func JSONWithStatus(ctx *fiber.Ctx, statusCode int, err error, result, errors interface{}) error {
	return JSON(ctx, err, result, errors, statusCode)
}

func JSON(ctx *fiber.Ctx, err error, result, errors interface{}, statusCode int) error {
	code, msg := errno.Decode(err)
	res := Model{
		Code:    code,
		Message: msg,
		Result:  result,
		Errors:  errors,
	}

	return ctx.JSON(res)
}
