package main

import (
	"github.com/labstack/echo"

	"github.com/mrlsd/echo-cms/config"
	"github.com/mrlsd/echo-cms/modules/backend"
	"github.com/gpmgo/gopm/modules/log"
)

func main() {
	config.LoadConfig()

	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)
	e.Static("/", "static")

	backend.UrlRules(e)
	e.Start(":3000")
}
