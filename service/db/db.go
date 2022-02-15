package db

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"
	"os/exec"

	pb "github.com/kofeebrian/capamass/protos/amass/db"
)

type DBService struct {
	pb.UnimplementedDBServiceServer
}

func runDBCommand(ctx *context.Context, domain string) error {
	cmd := exec.Command("amass", "db")
	cmd.Args = append(cmd.Args, "-d", domain)
	cmd.Args = append(cmd.Args, "-json", "/data/out.json")

	return cmd.Run()
}

func (*DBService) Run(ctx context.Context, req *pb.DBRequest) (*pb.DBResponse, error) {
	id := req.Id
	domain := req.Domain

	if err := runDBCommand(&ctx, domain); err != nil {
		log.Panicf("fail to run db command: %v", err)
		return nil, err
	}

	file, err := os.Open("/data/out.json")
	if err != nil {
		log.Panicf("fail to open file: %v", err)
		return nil, err
	}
	defer file.Close()

	var result *pb.DBResult
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	m := scanner.Text()

	err = json.Unmarshal([]byte(m), &result)
	if err != nil {
		log.Panicf("fail to parse output: %v", err)
		return nil, err
	}

	return &pb.DBResponse{
		Id:     id,
		Domain: domain,
		Result: result,
	}, err
}
