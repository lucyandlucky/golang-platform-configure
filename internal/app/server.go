package app

import (
	pb "configure/api/configure/server/v1"
)

type Server struct {
	pb.UnimplementedServerServer
}
