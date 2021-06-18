package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hb0730/auto-sign/web"
	"github.com/mritd/logger"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger.Info("[main] start ....")
	app := &cli.App{
		Name:  "auto-sign server",
		Usage: "auto sign ",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "addr",
				Usage:   "服务监听端口",
				EnvVars: []string{"SERVER_ADDRESS"},
				Value:   ":8080",
			},
			&cli.StringFlag{
				Name:    "cron",
				Usage:   "定时任务表达式",
				EnvVars: []string{"SERVER_CRON"},
				Value:   "30 * * * *",
			},
		},
		Authors: []*cli.Author{
			{
				Name: "hb0730", Email: "huangbing0730@gmail.com",
			},
		},
		Action: func(c *cli.Context) error {
			app := fiber.New(
				fiber.Config{
					ServerHeader: "auto-sign",
					ErrorHandler: func(c *fiber.Ctx, err error) error {
						code := fiber.StatusInternalServerError
						if e, ok := err.(*fiber.Error); ok {
							code = e.Code
						}
						return c.Status(code).JSON(web.Response{
							Code:      code,
							Message:   err.Error(),
							Timestamp: time.Now().Unix(),
						})
					},
				})
			web.RouterSetup(app)
			//启动 cron 监听 shutdown指令
			go func() {
				err := StartCron(c.String("cron"))
				if err != nil {
					logger.Warn("[main] Cron start error ,Http Server shutdown")
					if err := app.Shutdown(); err != nil {
						logger.Errorf("[main] Server forced to shutdown error: %v", err)
					}
				}

				sigs := make(chan os.Signal)
				signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
				for range sigs {
					logger.Warn("[main] Received a termination signal, bark server shutdown...")
					if err := app.Shutdown(); err != nil {
						logger.Errorf("[main] Server forced to shutdown error: %v", err)
					}
				}
			}()
			return app.Listen(c.String("addr"))
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.Errorf("[main] start error,error message:【 %s 】", err.Error())
		os.Exit(-1)
	}
}
