package routes

import (
	"JetChatClientGo/api"
	"JetChatClientGo/utils"
	"net/http"
)

var CachedAppVariables = make(map[uint64]AppVariables)

type AppVariables struct {
	Self   api.User
	Guilds []api.Guild
}

func App(w http.ResponseWriter, r *http.Request) {
	id, err := api.GetSelfId(r)
	if err != nil {
		utils.LogError(err)
	}

	vars := CachedAppVariables[id]
	ShowTemplateOnGet(w, r, "app", utils.StructToMap(vars))
}
