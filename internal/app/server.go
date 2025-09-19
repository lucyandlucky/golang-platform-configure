package app

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	pb "github.com/lucyandlucky/golang-platform-configure/api/configure/server/v1"
	"github.com/lucyandlucky/golang-platform-configure/internal/conf"
	"github.com/lucyandlucky/golang-platform-configure/internal/domain/entity"
	"github.com/lucyandlucky/golang-platform-configure/internal/domain/service"
	"github.com/lucyandlucky/golang-platform-configure/internal/infra/dbs"
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
