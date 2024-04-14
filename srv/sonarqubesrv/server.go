package sonarqubesrv

import (
	"log"
	"net"

	"github.com/dkZzzz/quality_hub/cmd/sonarqube"
	"github.com/dkZzzz/quality_hub/config"
	"github.com/dkZzzz/quality_hub/db/etcd"
	"github.com/dkZzzz/quality_hub/proto/sonarqubepb"
	"google.golang.org/grpc"
)

func Init_server() {
	lis, err := net.Listen("tcp", config.Cfg.SonarqubeServerHost)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sonarqubepb.RegisterSonarQubeServer(s, &sonarqube.SonarQubeServerImpl{})
	log.Printf("server listening at %v", lis.Addr())
	etcd.Register("sonarqube")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	select {}
}
