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
	e.GET("/admin/pages", controllers.GetPage)
	e.GET("/admin/pages/create", controllers.GetPageCreate)

	e.POST("/admin/form", controllers.PostForm)
}
