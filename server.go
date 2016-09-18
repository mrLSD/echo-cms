package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

	"github.com/mrlsd/echo-go/modules/backend"
	"github.com/mrlsd/echo-go/config"
)

func main() {
	config.LoadConfig()

	e := echo.New()
	backend.UrlRules(e)
	e.Run(standard.New(":3000"))
}
