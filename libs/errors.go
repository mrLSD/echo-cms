package libs

import (
	"github.com/labstack/echo"
	"net/http"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}
	e := c.Echo()

	if e.Debug {
		msg = err.Error()
	}
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD { // Issue #608
			c.NoContent(code)
		} else {
			data := echo.HTTPError{Code: code, Message: msg}
			if err_render := c.Render(code, "errors/main", data); err_render != nil {
				e.Logger.Error(err_render)
			}
		}
	}
	e.Logger.Error(err)
}
