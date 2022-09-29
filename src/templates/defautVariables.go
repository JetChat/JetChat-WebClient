package templates

import (
	"JetChatClientGo/utils"
)

var defaultVariables = map[string]any{
	"Title":      "JetChat",
	"EmailRegex": utils.EmailRegex,
	"Connected":  false,
}

func SetDefaultVariable(name string, value any) {
	defaultVariables[name] = value
}
