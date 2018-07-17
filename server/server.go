package server

import "crypto/tls"

type Server struct {
	Port int
	tlsConfig *tls.Config

}

func NewServer() (server *Server, err error) {
	server = &Server{}
	return
}

func (*Server) Start() {

}