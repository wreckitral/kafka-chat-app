package route

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/swagger"
)

func GeneralRoute(a *fiber.App) {
    a.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "msg": "Hello world",
            "docs": "/swagger/index.html",
            "status": "/h34l7h",
        })
    })
}

func SwaggerRoute(a *fiber.App) {
    route := a.Group("/swagger")

    route.Get("*", swagger.HandlerDefault)
}

func NotFoundRoute(a *fiber.App) {
    a.Use(
        func(c *fiber.Ctx) error {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "msg": "endpoint not found",
            })
        },
    )
}
