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

// CreateServer 创建数据
func (s Server) CreateServer(ctx kratosx.Context, server *entity.Server) (uint32, error) {
	return server.Id, ctx.DB().Create(server).Error
}

func (s Server) GetServerByKeyword(ctx kratosx.Context, keyword string) (*entity.Server, error) {
	return nil, nil
}
