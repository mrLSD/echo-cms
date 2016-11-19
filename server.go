package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/middleware"
	"github.com/mrlsd/echo-cms/config"
	"github.com/mrlsd/echo-cms/modules/backend"
)

func main() {
	config.LoadConfig()

	e := echo.New()

	// Routers
	backend.UrlRules(e)
	e.Static("/", "static")

	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.NonWWWRedirect())
	e.Use(middleware.Secure())
	e.Use(middleware.CSRF())
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2M"))

	e.Logger.Fatal(e.Start(":3000"))
}
