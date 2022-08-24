package library

import "net/http"

type Server struct {
	httpHandler *http.Handler
}

func (server *Server) Run(address) {

}
