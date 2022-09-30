package templates

import (
	"encoding/json"
	"html/template"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func repeat(s string, n int) template.HTML {
	return template.HTML(strings.Repeat(s, n))
}

func join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func split(s, sep string) []string {
	return strings.Split(s, sep)
}

func jsonify(v any) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func jsonPretty(v any) string {
	marshal, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return ""
	}
	return string(marshal)
}

func html(s string) template.HTML {
	return template.HTML(s)
}

func js(s string) template.JS {
	return template.JS(s)
}

func putOrElse(v, def any) any {
	if v == nil {
		return def
	}

	return v
}

func jsImport(s string) template.HTML {
	return template.HTML("<script type=\"module\" src=\"/static/js/" + s + "\"></script>")
}

func cssImport(s string) template.HTML {
	return template.HTML("<link rel=\"stylesheet\" href=\"/static/css/" + s + "\">")
}

var defaultFunctions = template.FuncMap{
	"add":        add,
	"sub":        sub,
	"mul":        mul,
	"div":        div,
	"repeat":     repeat,
	"join":       join,
	"split":      split,
	"json":       jsonify,
	"jsonPretty": jsonPretty,
	"html":       html,
	"js":         js,
	"putOrElse":  putOrElse,
	"jsImport":   jsImport,
	"cssImport":  cssImport,
}
