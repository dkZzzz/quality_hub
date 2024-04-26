package noticesrv

import (
	"log"
	"net"

	"github.com/dkZzzz/quality_hub/cmd/notice"
	"github.com/dkZzzz/quality_hub/config"
	"github.com/dkZzzz/quality_hub/db/etcd"
	"github.com/dkZzzz/quality_hub/proto/noticepb"
	"google.golang.org/grpc"
)

func Init_server() {
	lis, err := net.Listen("tcp", config.Cfg.NoticeServerHost)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	noticepb.RegisterNoticeServer(s, &notice.NoticeServerImpl{})
	log.Printf("server listening at %v", lis.Addr())
	etcd.Register("notice")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	select {}
}
