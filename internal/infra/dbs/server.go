package dbs

import (
	"fmt"
	"github.com/limes-cloud/kratosx"
	"github.com/lucyandlucky/golang-platform-configure/internal/domain/entity"
	"github.com/lucyandlucky/golang-platform-configure/types"
	"google.golang.org/protobuf/proto"
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

// ListServer 获取服务列表
func (s Server) ListServer(ctx kratosx.Context, req *types.ListServerRequest) ([]*entity.Server, uint32, error) {
	var (
		list  []*entity.Server
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.Server{}).Select(fs)
	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Ids != nil {
		db = db.Where("id IN ?", req.Ids)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}
	return list, uint32(total), db.Find(&list).Error
}
