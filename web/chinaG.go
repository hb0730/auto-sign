package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/support"
)

var gg = support.ChinaG{}

func init() {
	gg.ISupport = gg
	registerRoute("chinaG", func(c *fiber.App) {
		c.Group("/chinag").
			Post("/user", func(c *fiber.Ctx) error {
				return getChinaGBody(c)
			})
	})
}

func getChinaGBody(c *fiber.Ctx) error {
	var user ChinaGUser
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}
	if user.Username == "" || user.Password == "" {
		return c.Status(200).
			JSON(failed(201, "username/password is null"))
	}
	yaml := config.ReadYaml()
	yaml.Set(support.GetChinaGYamlKey(), user)
	err = yaml.WriteConfig()
	if err != nil {
		return err
	}

	config.LoadYaml()
	go func() {
		gg.Run()
	}()

	return c.Status(200).
		JSON(success())
}

type ChinaGUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
