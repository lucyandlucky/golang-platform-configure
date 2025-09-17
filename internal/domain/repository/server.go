package repository

import (
	"configure/internal/domain/entity"
	"github.com/limes-cloud/kratosx"
)

type Server interface {
	// CreateServer 创建服务信息
	CreateServer(ctx kratosx.Context, server *entity.Server) (uint32, error)
}
