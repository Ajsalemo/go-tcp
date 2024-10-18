package main

import (
	"github.com/gofiber/fiber/v2"

	"go.uber.org/zap"
	controllers "go-tcp-client/controllers"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	app := fiber.New()
	app.Get("/", controllers.IndexController)
	app.Get("/api/tcp/:host/:port", controllers.ExecuteTcpConnectionController)
	
	zap.L().Info("Fiber (HTTP) server is running on port 3000")
	fiberErr := app.Listen(":3080")

	if fiberErr != nil {
		zap.L().Fatal(fiberErr.Error())
	}
}
