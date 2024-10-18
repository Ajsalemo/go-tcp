package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	controllers "go-tcp-server/controllers"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func handleRequest(conn net.Conn) {
	// close conn
	defer conn.Close()
	// write data to response
	time := time.Now().Format(time.ANSIC)
	res := "Recieved connection at: " + time
	zap.L().Info("TCP server [1]: Recieved connection at: " + time)
	conn.Write([]byte(res))

}

func main() {
	app := fiber.New()
	app.Get("/", controllers.IndexController)
	// Notify the application of the below signals to be handled on shutdown
	s := make(chan os.Signal, 1)
	signal.Notify(s,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	// Goroutine to clean up prior to shutting down
	go func() {
		sig := <-s
		switch sig {
		case os.Interrupt:
			zap.L().Warn("CTRL+C / os.Interrupt recieved, shutting down the application..")
			app.Shutdown()
		case syscall.SIGTERM:
			zap.L().Warn("SIGTERM recieved.., shutting down the application..")
			app.Shutdown()
		case syscall.SIGQUIT:
			zap.L().Warn("SIGQUIT recieved.., shutting down the application..")
			app.Shutdown()
		case syscall.SIGINT:
			zap.L().Warn("SIGINT recieved.., shutting down the application..")
			app.Shutdown()
		}
	}()
	go func() {
		const (
			HOST = "0.0.0.0"
			PORT = "8080"
			TYPE = "tcp"
		)

		zap.L().Info("TCP server [1] is listening on port 8080")
		listen, err := net.Listen(TYPE, HOST+":"+PORT)
		if err != nil {
			zap.L().Fatal(err.Error())
		}
		// close listener
		defer listen.Close()
		for {
			conn, err := listen.Accept()
			if err != nil {
				zap.L().Fatal(err.Error())
			}
			go handleRequest(conn)
		}
	}()

	zap.L().Info("Fiber (HTTP) server is running on port 3000")
	fiberErr := app.Listen(":3000")

	if fiberErr != nil {
		zap.L().Fatal(fiberErr.Error())
	}
}
