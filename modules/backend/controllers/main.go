package controllers

import (
	"github.com/labstack/echo"
	"html/template"
	"net/http"

	"github.com/mrlsd/echo-cms/config"
)

func GetMain(c echo.Context) error {
	head := config.GetSiteHeader("backend")
	head.SetTitle("Main+")
	data := struct {
		Name template.HTML
		Head config.SiteHeader
	}{
		Name: "Test <b>BOLD</b>",
		Head: head,
	}
	return c.Render(http.StatusOK, "admin/main", data)
}

func GetLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "admin/login", nil)
}
