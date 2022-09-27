package routes

import (
	"JetChatClientGo/utils"
	"net/http"
	"regexp"
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
			Password: r.FormValue("password"),
		}

		identifiant := r.FormValue("identifiant")

		if regexp.MustCompile(utils.EmailRegex).MatchString(identifiant) {
			credentials.Email = identifiant
		} else {
			credentials.Username = identifiant
		}

		utils.Logger.Println(credentials)
	}
}
