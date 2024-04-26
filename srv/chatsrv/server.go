package chatsrv

import (
	"log"
	"net"

	"github.com/dkZzzz/quality_hub/cmd/chat"
	"github.com/dkZzzz/quality_hub/config"
	"github.com/dkZzzz/quality_hub/db/etcd"
	"github.com/dkZzzz/quality_hub/proto/chatpb"
	"google.golang.org/grpc"
)

func Init_server() {
	lis, err := net.Listen("tcp", config.Cfg.ChatServerHost)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	chatpb.RegisterChatServer(s, &chat.ChatServerImpl{})
	log.Printf("server listening at %v", lis.Addr())
	etcd.Register("chat")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	select {}
}
