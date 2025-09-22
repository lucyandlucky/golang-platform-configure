package repository

import (
	"github.com/limes-cloud/kratosx"
	"github.com/lucyandlucky/golang-platform-configure/internal/domain/entity"
	"github.com/lucyandlucky/golang-platform-configure/types"
)

type Server interface {
	// GetServerByKeyword  获取指定的服务信息
	GetServerByKeyword(ctx kratosx.Context, keyword string) (*entity.Server, error)

	// CreateServer 创建服务信息
	CreateServer(ctx kratosx.Context, server *entity.Server) (uint32, error)

	//	ListServer 获取服务列表
	ListServer(ctx kratosx.Context, req *types.ListServerRequest) ([]*entity.Server, uint32, error)
}
