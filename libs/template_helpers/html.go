package libs

import ("fmt"
"html/template"
)

func CssHelper(args ...interface{}) template.HTML {
	var content string
	content = ""
	if len(args) > 0 {
		content = fmt.Sprintf("<link rel=\"stylesheet\" href=\"%s\">", args[0])
	}
	return  template.HTML(content)
}

func JsHelper(args ...interface{}) template.HTML {
	var content string
	content = ""
	if len(args) > 0 {
		content = fmt.Sprintf("<script src=\"%s\"></script>", args[0])
	}
	return  template.HTML(content)
}
