package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/support"
)

func init() {
	registerRoute("geekhub", func(c *fiber.App) {
		c.Group("/geekhub").
			Post("/cookie", func(c *fiber.Ctx) error {
				return getCookie(c)
			})
	})
}

// getCookie 获取Cookie
func getCookie(c *fiber.Ctx) error {
	var cookies map[string]string
	c.BodyParser(&cookies)
	if len(cookies) == 0 {
		return c.Status(200).JSON(failed(201, "Cookies size 0"))
	} else {
		return writerYaml(cookies, c)
	}
}

// writerYaml 将cookie写入yaml并触发sign
func writerYaml(cookies map[string]string, c *fiber.Ctx) error {
	yaml := config.ReadYaml()
	yaml.Set(support.GeekhubYamlKey(), cookies)
	_ = yaml.WriteConfig()

	config.LoadYaml()

	hub := support.Geekhub{}
	hub.ISupport = hub
	hub.DoRun()

	return c.Status(200).JSON(success())
}
