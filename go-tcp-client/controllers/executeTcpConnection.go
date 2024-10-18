package controllers

import (
	"bufio"
	"net"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type TcpMessage struct {
	Msg string
}

func ExecuteTcpConnectionController(c *fiber.Ctx) error {
	host := c.Params("host")
	port := c.Params("port")
	TYPE := "tcp"

	tcpServer, err := net.ResolveTCPAddr(TYPE, host+":"+port)

	if err != nil {
		zap.L().Fatal(err.Error())
	}
	zap.L().Info("Attempting to connect to TCP server " + host + ":" + port)
	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		zap.L().Error(err.Error())
	}

	// Defer closing the connection
	defer conn.Close()
	// buffer to get data
	message, _ := bufio.NewReader(conn).ReadString('\n')
	zap.L().Info("Received message from TCP server " + host + ":" + port + " " + message)

	res := TcpMessage{
		Msg: message,
	}
	return c.JSON(res)
}
