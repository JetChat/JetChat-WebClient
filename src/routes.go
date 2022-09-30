package main

import (
	"JetChatClientGo/routes"
	"JetChatClientGo/server"
	"JetChatClientGo/utils"
	"net/http"
	"strings"
)

func registerRoutes(server *server.Server) {
	server.ServeFiles("/static/", http.Dir("./static/"))

	server.AddMiddleware(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			utils.Logger.Printf("%s %s", r.Method, r.RequestURI)
			next.ServeHTTP(w, r)
		})
	})

	server.AddMiddleware(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api/") || r.URL.Path == "/app" {
				if routes.CheckConnection(w, r) {
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	})

	server.AddRoute("/login", routes.Login).Methods("GET", "POST")
	server.AddRoute("/app", routes.App).Methods("GET")
}
