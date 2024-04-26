package chatsrv

import (
	"log"

	"github.com/dkZzzz/quality_hub/config"
	"github.com/dkZzzz/quality_hub/proto/chatpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Client chatpb.ChatClient
)

func Init_client() {
	conn, err := grpc.Dial(config.Cfg.SonarqubeServerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	Client = chatpb.NewChatClient(conn)

	log.Println("client connected")
	select {}
}
