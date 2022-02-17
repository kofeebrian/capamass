package db

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"

	pb "github.com/kofeebrian/capamass/protos/amass/db"
)

type DBService struct {
	pb.UnimplementedDBServiceServer
}

func runDBCommand(ctx *context.Context, domain string, config *pb.DBConfig) ([]byte, error) {
	cmd := exec.Command("amass", "db")
	cmd.Args = append(cmd.Args, "-dir", "/.config/amass")

	if latest := config.GetLatest(); latest {
		cmd.Args = append(cmd.Args, "-enum", "1")
	}

	if domain != "" {
		cmd.Args = append(cmd.Args, "-d", domain)
	}

	cmd.Args = append(cmd.Args, "-json", "-")

	return cmd.Output()
}

func (*DBService) Run(ctx context.Context, req *pb.DBRequest) (*pb.DBResponse, error) {
	id := req.GetId()
	domain := req.GetDomain()
	config := req.GetConfig()

	out, err := runDBCommand(&ctx, domain, config)
	if err != nil {
		log.Panicf("fail to run db command: %v", err)
		return nil, err
	}

	var result *pb.DBResult
	err = json.Unmarshal(out, &result)
	if err != nil {
		log.Panicf("fail to parse output: %v", err)
		return nil, err
	}

	return &pb.DBResponse{
		Id:     id,
		Domain: &domain,
		Result: result,
	}, err
}
