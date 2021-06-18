package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/support"
)

var ld246 = support.Ld246{}

func init() {
	ld246.ISupport = ld246
	registerRoute("ld246", func(c *fiber.App) {
		c.Group("/ld246").
			Post("/user", func(c *fiber.Ctx) error {
				return getLd246Body(c)
			})
	})
}

func getLd246Body(c *fiber.Ctx) error {
	var user Ld246User
	_ = c.BodyParser(&user)
	if user.Username == "" || user.Password == "" {
		return c.Status(200).JSON(failed(201, "username/password is null"))
	}
	yaml := config.GetViper()
	yaml.Set(support.GetLd246YamlKey(), user)
	go Run(ld246)
	return c.Status(200).JSON(success())

}

type Ld246User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
