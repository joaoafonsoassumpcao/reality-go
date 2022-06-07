package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Static("/", "./public/")

	app.Static("/", "./views/index.tpl")

	app.Listen(":3000")

}
