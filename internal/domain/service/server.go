package service

import (
	"github.com/limes-cloud/kratosx"
	"github.com/lucyandlucky/golang-platform-configure/api/configure/errors"
	"github.com/lucyandlucky/golang-platform-configure/internal/conf"
	"github.com/lucyandlucky/golang-platform-configure/internal/domain/entity"
	"github.com/lucyandlucky/golang-platform-configure/internal/domain/repository"
	"github.com/lucyandlucky/golang-platform-configure/types"
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
func (u *Server) CreateServer(ctx kratosx.Context, req *entity.Server) (uint32, error) {
	id, err := u.repo.CreateServer(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// ListServer 获取服务列表
func (u *Server) ListServer(ctx kratosx.Context, req *types.ListServerRequest) ([]*entity.Server, uint32, error) {
	list, total, err := u.repo.ListServer(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}
