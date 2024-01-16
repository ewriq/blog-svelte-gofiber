package main 

import (
	Router "basic-blog/Router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	Router.Initalize(app)

	app.Listen(":3000")
}