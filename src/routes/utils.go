package routes

import (
	"JetChatClientGo/templates"
	"JetChatClientGo/utils"
	"net/http"
)

func ShowTemplateOnGet(w http.ResponseWriter, r *http.Request, templateName string, additionalData ...map[string]any) {
	if r.Method == "GET" {
		totalData := make(map[string]any)
		for _, data := range additionalData {
			for key, value := range data {
				totalData[key] = value
			}
		}

		err := templates.RenderTemplate(templateName, w, totalData)
		if err != nil {
			utils.LogError(err)
		}
	}
}
