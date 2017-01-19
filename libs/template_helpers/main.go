package libs

import (
	"html/template"
)

func Helpers() template.FuncMap {
	return template.FuncMap{
		"css": CssHelper,
		"js": JsHelper,
	}
}