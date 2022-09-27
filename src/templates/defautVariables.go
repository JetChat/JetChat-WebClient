package templates

import (
	"JetChatClientGo/utils"
)

var defaultVariables = map[string]any{
	"Title":      "JetChat",
	"EmailRegex": utils.EmailRegex,
}

func AddDefaultVariable(name string, value any) {
	defaultVariables[name] = value
}
