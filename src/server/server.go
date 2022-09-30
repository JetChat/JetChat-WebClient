package server

import (
	"JetChatClientGo/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	NativeServer   *http.Server
	Router         *mux.Router
	port           int
	DefaultHandler http.Handler
}

func NewServer(port string) *Server {
	intPort, err := strconv.Atoi(port)
	if err != nil {
		utils.FatalError(err)
	}

	return &Server{
		Router:         mux.NewRouter(),
		port:           intPort,
		DefaultHandler: nil,
	}
}

func (s *Server) AddMiddleware(middleware mux.MiddlewareFunc) {
	s.Router.Use(middleware)
}

func (s *Server) AddRoute(path string, handler func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	handlerFunc := http.HandlerFunc(handler)
	utils.Logger.Println("Added route: " + path)
	return s.Router.Handle(path, handlerFunc)
}

func (s *Server) ServeFiles(path string, dir http.FileSystem) {
	s.Router.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(dir)))
}

func (s *Server) Start() error {
	s.NativeServer = &http.Server{
		Addr:         ":" + strconv.Itoa(s.port),
		Handler:      s.Router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	utils.Logger.Println("Starting server on port " + strconv.Itoa(s.port))
	return s.NativeServer.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.NativeServer.Close()
}

func (s *Server) SetDefaultHandler(handler func(w http.ResponseWriter, r *http.Request)) {
	s.DefaultHandler = http.HandlerFunc(handler)
	s.Router.NotFoundHandler = s.DefaultHandler

	utils.Logger.Println("Added default route /")
}
