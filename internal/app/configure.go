package app

import (
	pb "configure/api/configure/configure/v1"
)

type Configure struct {
	pb.UnimplementedConfigureServer
}
