package service

import (
	"fmt"
	"log"
	"net"

	"github.com/kofeebrian/capamass/config"
	dbpb "github.com/kofeebrian/capamass/protos/amass/db"
	enumpb "github.com/kofeebrian/capamass/protos/amass/enum"
	vizpb "github.com/kofeebrian/capamass/protos/amass/viz"
	"github.com/kofeebrian/capamass/service/db"
	"github.com/kofeebrian/capamass/service/enum"
	"github.com/kofeebrian/capamass/service/viz"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Init(config *config.ServiceConfig) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register Service Servers
	enumpb.RegisterEnumServiceServer(s, &enum.EnumServer{})
	dbpb.RegisterDBServiceServer(s, &db.DBService{})
	vizpb.RegisterVizServiceServer(s, &viz.VizService{})

	reflection.Register(s)

	log.Printf("serve at %s", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
