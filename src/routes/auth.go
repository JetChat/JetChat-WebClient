package routes

import (
	"JetChatClientGo/api"
	"JetChatClientGo/templates"
	"JetChatClientGo/utils"
	"errors"
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

		connectionCookie, err := Connect(credentials)
		if err != nil {
			if errors.Is(err, api.ErrInvalidCredentials) {
				templates.SetDefaultVariable("Error", "Incorrect identifiant or password")
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			} else {
				utils.LogError(err)
			}
		}
		utils.Logger.Println("Connection cookie: " + connectionCookie.String())
		templates.SetDefaultVariable("Connected", err == nil)
		api.ConnectionCookie = connectionCookie

		self, err := api.GetSelf(r)
		if err != nil {
			utils.LogError(err)
		}

		templates.SetDefaultVariable("Self", self)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Connect(credentials Credentials) (*http.Cookie, error) {
	route := "/login"
	request := api.NewRequest[any](route)
	request.Json()
	request.SetBody(credentials)

	res, err := request.Send()
	if err != nil {
		if res != nil && res.StatusCode == http.StatusNotFound {
			return nil, api.ErrInvalidCredentials
		}

		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, api.ErrInvalidCredentials
	}

	return res.Cookies()[0], nil
}
