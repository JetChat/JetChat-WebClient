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

var CredentialsCache = make(map[string]Credentials)

func Login(w http.ResponseWriter, r *http.Request) {
	ShowTemplateOnGet(w, r, "login")

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			utils.FatalError(err)
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
				utils.FatalError(err)
			}
		}
		templates.SetDefaultVariable("Connected", err == nil)
		r.AddCookie(connectionCookie)
		w.Header().Add("Set-Cookie", connectionCookie.String())

		CredentialsCache[r.RemoteAddr] = credentials

		self, err := api.GetSelf(r)
		if err != nil {
			utils.FatalError(err)
		}
		userAppVariables := CachedAppVariables[self.UserID]
		userAppVariables.Guilds, err = api.GetGuilds(r)
		if err != nil {
			utils.LogError(err)
		}

		CachedAppVariables[self.UserID] = userAppVariables

		templates.SetDefaultVariable("Self", self)
		http.Redirect(w, r, "/app", http.StatusSeeOther)
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

func Reconnect(r *http.Request) error {
	credentials := CredentialsCache[r.RemoteAddr]
	cookie, err := Connect(credentials)
	if err != nil {
		return err
	}

	r.AddCookie(cookie)
	return nil
}

func CheckConnection(w http.ResponseWriter, r *http.Request) (redirect bool) {
	sessionCookie, err := r.Cookie(utils.ConnectionCookie)
	if err != nil || sessionCookie.Value == "" {
		err = Reconnect(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return true
		}
	}

	return false
}
