package main

import (
    "log"
	"os"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{
            "message": "Hello world!",
        })
    })

    log.Fatal(app.Listen("0.0.0.0:" + os.Getenv("PORT")))
}