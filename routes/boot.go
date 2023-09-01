package routes

import (
	"github.com/gofiber/fiber/v2"
)

type routeFunc func(fiber.Router)

func Handle(App *fiber.App) {
	App.Route(apiPrefix, apiRoute, "api")
	App.Route(adminPrefix, adminRoute, "admin")
}
