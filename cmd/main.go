package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"shorturl/internal/config"
	"shorturl/internal/controller"
	"shorturl/internal/middleware"
	"shorturl/internal/model"
	"shorturl/internal/pkg/logging"
)

func main() {
	cfg := config.GetConfig()

	logging.Init()

	model.InitDB(cfg.Sqlite.StoragePath)

	engine := html.New("./assets/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./assets/static")

	app.Use(middleware.Logger)

	app.Get("/", controller.Index)
	app.Post("/shorter", controller.Shorter)
	app.Get("/redirect/:short_url", controller.Redirect)

	app.Use(middleware.PageNotFound)

	app.Listen(fmt.Sprintf(":%s", cfg.App.Port))
}
