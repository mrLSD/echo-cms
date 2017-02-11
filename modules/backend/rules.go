package backend

import (
	"github.com/labstack/echo"

	"github.com/mrlsd/echo-cms/libs"
	"github.com/mrlsd/echo-cms/modules/backend/controllers"
)

func UrlRules(e *echo.Echo) {
	libs.SetRenderer(e, "modules/backend/views/**/*.html")

	e.GET("/admin", controllers.GetMain)
	e.GET("/admin/form", controllers.GetForm)
	e.GET("/admin/login", controllers.GetLogin)

	e.POST("/admin/form", controllers.PostForm)
}
