package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/support"
)

var apple = support.AppleTuan{}

func init() {
	apple.ISupport = apple
	registerRoute("appletuan", func(c *fiber.App) {
		c.Group("/appletuan").
			Post("/cookie", func(c *fiber.Ctx) error {
				return getAppleTuanBody(c)
			})
	})
}

func getAppleTuanBody(c *fiber.Ctx) error {
	var Cookie map[string]string
	_ = c.BodyParser(&Cookie)
	if len(Cookie) == 0 {
		return c.Status(200).JSON(failed(200, "cookie size 0"))
	}
	yaml := config.ReadYaml()
	yaml.Set(support.GetAppleTuanYamlKey(), Cookie)
	_ = yaml.WriteConfig()
	config.LoadYaml()
	go func() {
		apple.Run()
	}()
	return c.Status(200).JSON(success())
}
