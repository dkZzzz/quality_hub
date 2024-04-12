package sonarqubesrv

import (
	"log"
	"net"

	"github.com/dkZzzz/quality_hub/cmd/sonarqube"
	"github.com/dkZzzz/quality_hub/db/etcd"
	"github.com/dkZzzz/quality_hub/proto/sonarqubepb"
	"google.golang.org/grpc"
)

func Init_server() {
	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sonarqubepb.RegisterSonarQubeServer(s, &sonarqube.SonarQubeServerImpl{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	etcd.Register("sonarqube")
	select {}
}
