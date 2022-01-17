package enum

import (
	"bytes"
	"context"
	"log"
	"os/exec"

	pb "github.com/kofeebrian/capamass/protos/amass/enum"
)

type EnumServer struct {
	pb.UnimplementedEnumServiceServer
}

var (
	stdout, stderr bytes.Buffer
)

func enumerate(name string) (string, error) {
	// TODO: use config.ini
	cmd := exec.Command("amass", "enum")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (*EnumServer) BasicEnumerate(ctx context.Context, req *pb.EnumRequest) (*pb.EnumResponse, error) {
	log.Printf("starting enumeration...")

	result, err := enumerate(req.DomainName)
	if err != nil {
		log.Panicf("failed to enumerate: %v", err)
		return &pb.EnumResponse{
			Result: "failed to enumerate",
		}, err
	}

	return &pb.EnumResponse{
		Result: result,
	}, nil
}
