package templates

import (
	"encoding/json"
	"html/template"
	"strings"
)

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func Repeat(s string, n int) template.HTML {
	return template.HTML(strings.Repeat(s, n))
}

func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

func Json(v any) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func JsonPretty(v any) string {
	marshal, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return ""
	}
	return string(marshal)
}

func HTML(s string) template.HTML {
	return template.HTML(s)
}

func JS(s string) template.JS {
	return template.JS(s)
}

func PutOrElse(v, def any) any {
	if v == nil {
		return def
	}

	return v
}

var DefaultFunctions = template.FuncMap{
	"add":        Add,
	"sub":        Sub,
	"mul":        Mul,
	"div":        Div,
	"repeat":     Repeat,
	"join":       Join,
	"split":      Split,
	"json":       Json,
	"jsonPretty": JsonPretty,
	"html":       HTML,
	"js":         JS,
	"putOrElse":  PutOrElse,
}
