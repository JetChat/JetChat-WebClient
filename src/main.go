package main

import (
	"JetChatClientGo/api"
	server2 "JetChatClientGo/server"
	"JetChatClientGo/templates"
	"JetChatClientGo/utils"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		utils.LogError(err)
	}

	port := os.Getenv("PORT")
	api.ApiURL = os.Getenv("API_URL")

	templates.SetDefaultVariable("Env", os.Getenv("ENV"))
	templates.SetDefaultVariable("Debug", os.Getenv("ENV") == "debug")
	templates.SetDefaultVariable("Envs", os.Environ())

	server := server2.NewServer(port)
	server.SetDefaultHandler(func(w http.ResponseWriter, r *http.Request) {
		utils.Logger.Println("Route not found: " + r.URL.Path)
	})

	staticFolder := http.FileServer(http.Dir("static"))
	server.AddHandler("/static/", http.StripPrefix("/static/", staticFolder))

	utils.Logger.Println("Starting server on port " + port)

	RegisterRoutes(server)

	err = server.Start()
	if err != nil {
		utils.LogError(err)
	}
}
