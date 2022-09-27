package main

import (
	server2 "JetChatClientGo/server"
	"JetChatClientGo/templates"
	"JetChatClientGo/utils"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strconv"
)

var DB *sql.DB

const port = 8080

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		utils.LogError(err)
	}

	templates.AddDefaultVariable("Env", os.Getenv("ENV"))
	templates.AddDefaultVariable("Debug", os.Getenv("ENV") == "debug")
	templates.AddDefaultVariable("Envs", os.Environ())

	server := server2.NewServer(port)
	server.SetDefaultHandler(func(w http.ResponseWriter, r *http.Request) {
		utils.Logger.Println("Route not found: " + r.URL.Path)
	})

	staticFolder := http.FileServer(http.Dir("static"))
	server.AddHandler("/static/", http.StripPrefix("/static/", staticFolder))

	utils.Logger.Println("Starting server on port " + strconv.Itoa(port))

	RegisterRoutes(server)

	err = server.Start()
	if err != nil {
		utils.LogError(err)
	}

	dbConfig := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}

	DB, err = sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		utils.LogError(err)
	}
}
