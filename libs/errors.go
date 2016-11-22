package libs

import (
	"net/http"
	"github.com/labstack/echo"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}
	e := c.Echo()
	println("\t=> Debug mode: ", e.Debug)
	println("\t=> Code: ", code, c.Request().Method)
	println("\t=> Msg: ", msg)

	if e.Debug {
		msg = err.Error()
	}
	if !c.Response().Committed {
		println("\t=> Commited: ", c.Response().Committed)
		if c.Request().Method == echo.HEAD { // Issue #608
			c.NoContent(code)
		} else {
			c.Render(code, "errors/main", msg)
		}
	}
	e.Logger.Error(err)
}
