package usersrv

import (
	"log"
	"net"

	"github.com/dkZzzz/quality_hub/cmd/user"
	"github.com/dkZzzz/quality_hub/config"
	"github.com/dkZzzz/quality_hub/db/etcd"
	"github.com/dkZzzz/quality_hub/proto/userpb"
	"google.golang.org/grpc"
)

func Init_server() {
	lis, err := net.Listen("tcp", config.Cfg.UserServerHost)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &user.UserServerImpl{})
	log.Printf("server listening at %v", lis.Addr())
	etcd.Register("user")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	select {}
}
