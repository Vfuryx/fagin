package routes

import (
	adminController "fadmin/app/controllers/admin"
	"fadmin/pkg/router/no_router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

const adminMax = 1000
const adminDuration = 60 * time.Second
const adminPrefix = "/admin"

var adminRoute routeFunc = func(admin fiber.Router) {
	// 404
	no_router.NoRoute(adminPrefix, func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"m": "admin 404"})
	})

	//// 测试限流中间件
	admin.Use(limiter.New(limiter.Config{
		Max:        adminMax,
		Expiration: adminDuration,
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusTooManyRequests).SendString("服务器繁忙")
		},
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	V1 := admin.Group("/v1")
	{
		user := V1.Group("/user")
		{
			user.Post("/index", adminController.UserController.Index)
			user.Post("/add", adminController.UserController.Add)
		}
	}
}
