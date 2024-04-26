package noticesrv

import (
	"log"

	"github.com/dkZzzz/quality_hub/config"
	"github.com/dkZzzz/quality_hub/proto/noticepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Client noticepb.NoticeClient
)

func Init_client() {
	conn, err := grpc.Dial(config.Cfg.NoticeServerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	Client = noticepb.NewNoticeClient(conn)

	log.Println("client connected")
	select {}
}
