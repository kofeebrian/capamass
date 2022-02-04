package enum

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	pb "github.com/kofeebrian/capamass/protos/amass/enum"
)

type EnumServer struct {
	pb.UnimplementedEnumServiceServer
}

var (
	stdout, stderr bytes.Buffer
)

func enumerate(name string) error {
	cmd := exec.Command("amass", "enum")
	cmd.Args = append(cmd.Args, "-d", name)
	cmd.Args = append(cmd.Args, "-timeout", "5")                       // set timeout in min.
	cmd.Args = append(cmd.Args, "-config", "utils/config/default.ini") // use config TODO: user can choose
	cmd.Args = append(cmd.Args, "-ip")                                 // show ip
	cmd.Args = append(cmd.Args, "-src")                                // show source
	cmd.Args = append(cmd.Args, "-json", fmt.Sprintf("/data/%s.json", name))
	cmd.Args = append(cmd.Args, "-silent") // do not print output to stdout
	return cmd.Run()
}

func extractResult(name string) (*[]*pb.EnumResponse_Result, error) {
	f, err := os.Open(fmt.Sprintf("/data/%s.json", name))
	if err != nil {
		log.Panicf("failed to open result file: %v", err)
		return nil, err
	}

	var results []*pb.EnumResponse_Result
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m := scanner.Text()
		var result pb.EnumResponse_Result
		err := json.Unmarshal([]byte(m), &result)
		if err != nil {
			log.Panicf("failed to parse json: %v", err)
			continue
		}
		results = append(results, &result)
	}

	return &results, nil
}

func (*EnumServer) BasicEnumerate(ctx context.Context, req *pb.EnumRequest) (*pb.EnumResponse, error) {
	log.Printf("starting enumeration...")
	fmt.Printf("%v\n", req)

	err := enumerate(req.GetDomain())
	if err != nil {
		log.Panicf("failed to enumerate: %v", err)
	}

	if ctx.Err() != context.Canceled {
		return &pb.EnumResponse{}, ctx.Err()
	}

	results, err := extractResult(req.GetDomain())
	if err != nil {
		log.Panicf("extract results failed: %v", err)
		return &pb.EnumResponse{}, err
	}

	return &pb.EnumResponse{
		Results: *results,
	}, nil
}
