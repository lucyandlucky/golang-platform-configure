package repository

import (
	"configure/internal/domain/entity"
	"github.com/limes-cloud/kratosx"
)

type Server interface {
	// GetServerByKeyword  获取指定的服务信息
	GetServerByKeyword(ctx kratosx.Context, keyword string) (*entity.Server, error)
	// CreateServer 创建服务信息
	CreateServer(ctx kratosx.Context, server *entity.Server) (uint32, error)
}
