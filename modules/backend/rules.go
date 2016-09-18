package backend

import (
	"github.com/labstack/echo"

	"github.com/mrlsd/echo-go/libs"
	"github.com/mrlsd/echo-go/modules/backend/controllers"
)

func UrlRules(e *echo.Echo) {
	libs.SetRenderer(e, "modules/backend/views/**/*.html")

	e.GET("/admin/", controllers.GetMain)
	e.GET("/admin/login/", controllers.GetLogin)
}
