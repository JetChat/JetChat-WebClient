package main

import (
	"JetChatClientGo/routes"
	server2 "JetChatClientGo/server"
)

func RegisterRoutes(server *server2.Server) {
	server.AddRoute("/login", routes.Login)
}
