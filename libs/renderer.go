package libs

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func SetRenderer(e *echo.Echo, path string) {
	t := &Template{
		templates: template.Must(template.ParseGlob(path)),
	}
	e.SetRenderer(t)
}
