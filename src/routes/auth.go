package routes

import (
	"JetChatClientGo/utils"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	ShowTemplateOnGet(w, r, "login")

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			utils.LogError(err)
		}

		credentials := Credentials{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
		}

		utils.Logger.Println(credentials)
	}
}
