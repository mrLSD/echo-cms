package controllers

import (
	"github.com/mrlsd/echo-cms/config"
	"github.com/labstack/echo"
	"html/template"
	"net/http"
)

func GetPage(c echo.Context) error {
	head := config.GetSiteHeader("backend")
	head.SetTitle("Pages")
	data := struct {
		Name template.HTML
		Head config.SiteHeader
	}{
		Name: "Test pages",
		Head: head,
	}
	return c.Render(http.StatusOK, "admin/pages/main", data)
}

func GetPageCreate(c echo.Context) error {
	head := config.GetSiteHeader("backend")
	head.SetTitle("Pages")
	data := struct {
		Name template.HTML
		Head config.SiteHeader
	}{
		Name: "Test pages",
		Head: head,
	}
	return c.Render(http.StatusOK, "admin/pages/create", data)
}
