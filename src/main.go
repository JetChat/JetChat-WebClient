package main

import (
	"JetChatClientGo/api"
	server2 "JetChatClientGo/server"
	"JetChatClientGo/templates"
	"JetChatClientGo/utils"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		utils.FatalError(err)
	}

	port := os.Getenv("PORT")
	api.ApiURL = os.Getenv("API_URL")

	templates.SetDefaultVariable("Env", os.Getenv("ENV"))
	templates.SetDefaultVariable("Debug", os.Getenv("ENV") == utils.DebugMode)
	templates.SetDefaultVariable("Envs", os.Environ())

	server := server2.NewServer(port)
	registerRoutes(server)

	err = server.Start()
	if err != nil {
		utils.FatalError(err)
	}
}
