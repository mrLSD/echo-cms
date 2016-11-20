package main

import (
	"github.com/labstack/echo"

	"github.com/mrlsd/echo-cms/config"
	"github.com/mrlsd/echo-cms/modules/backend"
)

func main() {
	config.LoadConfig()

	e := echo.New()
	backend.UrlRules(e)
	e.Run(standard.New(":3000"))
}
