package app

import (
	pb "github.com/lucyandlucky/golang-platform-configure/api/configure/configure/v1"
)

type Configure struct {
	pb.UnimplementedConfigureServer
}
