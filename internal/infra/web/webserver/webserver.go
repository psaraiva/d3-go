package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(route string, handler http.HandlerFunc, method string) {
	// @todo: move to valid method
	if method != http.MethodGet &&
		method != http.MethodPost &&
		method != http.MethodPut &&
		method != http.MethodPatch &&
		method != http.MethodDelete {
		panic("unsupported method")
	}

	if s.Handlers[route] == nil {
		s.Handlers[route] = make(map[string]http.HandlerFunc)
	}
	s.Handlers[route][method] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, methods := range s.Handlers {
		for method, handler := range methods {
			switch method {
			case http.MethodGet:
				s.Router.Get(path, handler)
			case http.MethodPost:
				s.Router.Post(path, handler)
			case http.MethodPut:
				s.Router.Put(path, handler)
			case http.MethodPatch:
				s.Router.Patch(path, handler)
			case http.MethodDelete:
				s.Router.Delete(path, handler)
			default:
				panic("unsupported method")
			}
		}
	}
	err := http.ListenAndServe(":"+s.WebServerPort, s.Router)
	if err != nil {
		panic(err)
	}
}
