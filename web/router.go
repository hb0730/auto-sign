package web

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hb0730/auto-sign/utils"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

type routerFunc struct {
	Name   string
	Weight int
	Func   func(router *fiber.App)
}

type routeSlice []routerFunc

func (r routeSlice) Len() int { return len(r) }

func (r routeSlice) Less(i, j int) bool { return r[i].Weight > r[j].Weight }

func (r routeSlice) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

var routes routeSlice
var routerOnce sync.Once

func registerRoute(name string, f func(router *fiber.App)) {
	registerRouteWithWeight(name, 50, f)
}

func registerRouteWithWeight(name string, weight int, f func(route *fiber.App)) {
	if weight > 100 || weight < 0 {
		utils.WarnF("route [%s] weight must be >= 0 and <=100", name)
	}

	for _, route := range routes {
		if strings.ToLower(name) == strings.ToLower(route.Name) {
			utils.WarnF("route [%s] already registered", route.Name)
		}
	}
	routes = append(routes, routerFunc{
		Name:   name,
		Weight: weight,
		Func:   f,
	})
}

func RouterSetup(router *fiber.App) {
	routerOnce.Do(func() {
		router.Use(logger.New(logger.Config{
			Format:     "${time}     INFO    ${ip} -> [${status}] ${method} ${latency} ${route} => ${url} ${body}\n",
			TimeFormat: "2006-01-02 15:04:05",
			Output:     os.Stdout,
		}))
		router.Use(fiberecover.New())
		sort.Sort(routes)
		for _, r := range routes {
			r.Func(router)
			utils.InfoF("load route [%s] success...", r.Name)
		}
	})
}

// Response

func success() Response {
	return Response{
		Code:      200,
		Message:   "Success",
		Timestamp: time.Now().Unix(),
	}
}

func failed(code int, message string, args ...interface{}) Response {
	return Response{
		Code:      code,
		Message:   fmt.Sprintf(message, args...),
		Timestamp: time.Now().Unix(),
	}
}

func data(data interface{}) Response {
	return Response{
		Code:      200,
		Message:   "success",
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
}
