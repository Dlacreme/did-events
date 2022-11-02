package didgrpc

import (
	did_proto "github.com/dlacreme/did-proto"
)

type server struct {
	did_proto.UnimplementedDidServer
}

func newServer() *server {
	s := server{}
	return &s
}
