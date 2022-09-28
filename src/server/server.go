package server

import (
	"JetChatClientGo/utils"
	"net/http"
	"strconv"
)

type Server struct {
	NativeServer   *http.Server
	Mux            *http.ServeMux
	port           int
	DefaultHandler http.Handler
	Routes         map[string]http.Handler
	RecoverHandler func(w http.ResponseWriter, r *http.Request)
}

func NewServer(port string) *Server {
	intPort, err := strconv.Atoi(port)
	if err != nil {
		utils.LogError(err)
	}

	return &Server{
		NativeServer: &http.Server{
			Addr:         ":" + port,
			Handler:      nil,
			IdleTimeout:  0,
			ReadTimeout:  0,
			WriteTimeout: 0,
		},
		Mux:            http.NewServeMux(),
		port:           intPort,
		DefaultHandler: nil,
		Routes:         make(map[string]http.Handler),
		RecoverHandler: nil,
	}
}

func (s *Server) AddRoute(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	handlerFunc := http.HandlerFunc(handler)
	s.Routes[path] = handlerFunc
	s.Mux.Handle(path, handlerFunc)
	s.NativeServer.Handler = s.Mux

	utils.Logger.Println("Added route: " + path)
}

func (s *Server) AddHandler(path string, handler http.Handler) {
	s.Routes[path] = handler
	s.Mux.Handle(path, handler)
	s.NativeServer.Handler = s.Mux

	utils.Logger.Println("Added route: " + path)
}

func (s *Server) Start() error {
	return s.NativeServer.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.NativeServer.Close()
}

func (s *Server) SetDefaultHandler(handler func(w http.ResponseWriter, r *http.Request)) {
	s.DefaultHandler = http.HandlerFunc(handler)
}
