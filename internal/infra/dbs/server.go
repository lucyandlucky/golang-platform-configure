package dbs

import (
	"configure/internal/domain/entity"
	"github.com/limes-cloud/kratosx"
	"sync"
)

type Server struct {
}

var (
	serverIns  *Server
	serverOnce sync.Once
)

func NewServer() *Server {
	serverOnce.Do(func() {
		serverIns = &Server{}
	})
	return serverIns
}

func (s Server) CreateServer(ctx kratosx.Context, server *entity.Server) (uint32, error) {
	//TODO implement me
	panic("implement me")
}
