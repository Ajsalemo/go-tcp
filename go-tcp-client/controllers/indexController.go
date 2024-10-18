package controllers

import (
    "github.com/gofiber/fiber/v2"
)

type IndexControllerMessage struct {
	Msg string
  }

func IndexController(c *fiber.Ctx) error {
	res := IndexControllerMessage{
		Msg: "go-tcp-client",
	}
	return c.JSON(res)
}