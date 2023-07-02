package routes

import (
	"fadmin/app/controllers/api"
	"fadmin/pkg/router/no_router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

const apiMax = 1000
const apiDuration = 60 * time.Second
const apiPrefix = "/api"

var apiRoute routeFunc = func(Api fiber.Router) {
	// 404
	no_router.NoRoute(apiPrefix, func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"m": "api 404"})
	})

	//// 测试限流中间件
	Api.Use(limiter.New(limiter.Config{
		Max:        apiMax,
		Expiration: apiDuration,
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusTooManyRequests).SendString("服务器繁忙")
		},
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	V1 := Api.Group("/v1")
	{
		V1.Get("/index", api.IndexController.Index)
	}
}
