package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/support"
	"github.com/mritd/logger"
)

var v2ex = support.V2ex{}

func init() {
	v2ex.ISupport = v2ex
	registerRoute("v2ex", func(c *fiber.App) {
		c.Group("/v2ex").
			Post("/cookie", func(c *fiber.Ctx) error {
				return getV2exBody(c)
			})
	})
}
func getV2exBody(c *fiber.Ctx) error {
	var cookies map[string]string
	err := c.BodyParser(&cookies)
	if err != nil {
		return err
	}
	if len(cookies) == 0 {
		logger.Warn("[web v2ex] cookie size 0")
		return c.Status(200).JSON(failed(201, "cookies size 0"))
	}
	yaml := config.GetViper()
	yaml.Set(support.GetV2exYamlKey(), cookies)
	err = yaml.WriteConfig()
	if err != nil {
		return err
	}
	go Run(v2ex)
	return c.Status(200).JSON(success())
}
