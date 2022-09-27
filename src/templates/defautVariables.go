package templates

var defaultVariables = map[string]any{
	"Title": "JetChat",
}

func AddDefaultVariable(name string, value any) {
	defaultVariables[name] = value
}
