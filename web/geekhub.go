package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/support"
)

var hub = support.Geekhub{}

func init() {
	hub.ISupport = hub
	registerRoute("geekhub", func(c *fiber.App) {
		c.Group("/geekhub").
			Post("/cookie", func(c *fiber.Ctx) error {
				return getGeekhubCookies(c)
			}).
			Get("/start", func(c *fiber.Ctx) error {
				return GeekhubStart(c)
			})
	})
}

// getGeekhubCookies 获取Cookie
func getGeekhubCookies(c *fiber.Ctx) error {
	var cookies map[string]string
	_ = c.BodyParser(&cookies)
	if len(cookies) == 0 {
		return c.Status(200).JSON(failed(201, "Cookies size 0"))
	}

	yaml := config.GetViper()
	yaml.Set(support.GeekhubYamlKey(), cookies)
	_ = yaml.WriteConfig()
	go Run(hub)

	return c.Status(200).JSON(success())
}

func GeekhubStart(c *fiber.Ctx) error {
	return hub.DoRun()
}
