package app

import (
	pb "configure/api/configure/server/v1"
	"configure/internal/conf"
	"configure/internal/domain/entity"
	"configure/internal/domain/service"
	"configure/internal/infra/dbs"
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
)

type Server struct {
	pb.UnimplementedServerServer
	srv *service.Server
}

func NewServer(conf *conf.Config) *Server {
	return &Server{
		srv: service.NewServer(conf, dbs.NewServer()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewServer(c)
		pb.RegisterServerHTTPServer(hs, srv)
		pb.RegisterServerServer(gs, srv)
	})
}

// CreateServer 创建服务信息
func (s *Server) CreateServer(c context.Context, req *pb.CreateServerRequest) (*pb.CreateServerReply, error) {
	id, err := s.srv.CreateServer(kratosx.MustContext(c), &entity.Server{
		Keyword:     req.Keyword,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateServerReply{Id: id}, nil
}
