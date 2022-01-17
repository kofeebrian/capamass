package service

import (
	"fmt"
	"log"
	"net"

	"github.com/kofeebrian/capamass/config"
	enumpb "github.com/kofeebrian/capamass/protos/amass/enum"
	"github.com/kofeebrian/capamass/service/enum"
	"google.golang.org/grpc"
)

func Init(config *config.ServiceConfig) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register Service Servers
	enumpb.RegisterEnumServiceServer(s, &enum.EnumServer{})

	log.Printf("serve at %s", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
