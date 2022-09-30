package templates

import (
	"bytes"
	"html/template"
	"net/http"
	"path"
)

const templatesDirectory = "templates"
const mainTemplate = "main"

type Template struct {
	Variables map[string]any
	Functions template.FuncMap
	FileName  string
}

func (t *Template) Render(w http.ResponseWriter) error {
	tmpl := template.New("").Funcs(defaultFunctions).Funcs(t.Functions)
	tmpl, err := tmpl.ParseGlob(path.Join(templatesDirectory, "*.tmpl"))
	if err != nil {
		return err
	}

	rendered := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(rendered, t.FileName, t.Variables)
	if err != nil {
		return err
	}

	const renderedContentKey = "RenderedContent"
	t.Variables[renderedContentKey] = template.HTML(rendered.String())

	err = tmpl.ExecuteTemplate(w, mainTemplate, t.Variables)
	if err != nil {
		return err
	}

	return nil
}

func (t *Template) AddFunction(name string, function interface{}) {
	t.Functions[name] = function
}

func RenderTemplate(name string, w http.ResponseWriter, variables map[string]any) error {
	t := Template{
		Variables: variables,
		Functions: make(template.FuncMap),
		FileName:  name,
	}

	t.Variables["TemplateName"] = name

	for key, value := range defaultVariables {
		t.Variables[key] = value
	}

	return t.Render(w)
}
