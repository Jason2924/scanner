package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Start() error
	Stop(context.Context) error
}

type server struct {
	Server *http.Server
}

func NewServer(port string, ngin *gin.Engine) Server {
	serv := &http.Server{
		Addr:    ":" + port,
		Handler: ngin,
	}
	return &server{
		Server: serv,
	}
}

func (srv *server) Start() error {
	return srv.Server.ListenAndServe()
}

func (srv *server) Stop(ctxt context.Context) error {
	return srv.Server.Shutdown(ctxt)
}
