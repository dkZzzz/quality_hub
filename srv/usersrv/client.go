package usersrv

import (
	"log"

	"github.com/dkZzzz/quality_hub/config"
	"github.com/dkZzzz/quality_hub/proto/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Client userpb.UserServiceClient
)

func Init_client() {
	conn, err := grpc.Dial(config.Cfg.UserServerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	Client = userpb.NewUserServiceClient(conn)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// _, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	log.Println("client connected")
	select {}
}
