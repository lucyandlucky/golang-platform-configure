package service

import (
	"configure/api/configure/errors"
	"configure/internal/conf"
	"configure/internal/domain/entity"
	"configure/internal/domain/repository"
	"github.com/limes-cloud/kratosx"
)

type Server struct {
	conf *conf.Config
	repo repository.Server
}

func NewServer(conf *conf.Config, repo repository.Server) *Server {
	return &Server{
		conf: conf,
		repo: repo,
	}
}

// CreateServer 创建服务信息
func (s *Server) CreateServer(ctx kratosx.Context, req *entity.Server) (uint32, error) {
	id, err := s.repo.CreateServer(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}

	return id, nil
}
