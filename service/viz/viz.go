package viz

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	pb "github.com/kofeebrian/capamass/protos/amass/viz"
)

type VizService struct {
	pb.UnimplementedVizServiceServer
}

func runVizCommand(ctx context.Context, domain string, config *pb.VizConfig) error {
	cmd := exec.CommandContext(ctx, "amass", "viz")
	cmd.Args = append(cmd.Args, "-dir", "/.config/amass")
	cmd.Args = append(cmd.Args, "-d", domain)

	if graphistry := config.GetGraphistry(); graphistry {
		cmd.Args = append(cmd.Args, "-graphistry")
	}

	cmd.Args = append(cmd.Args, "-o", "/data")

	return cmd.Run()
}

func readResult(ctx *context.Context, result *pb.GraphistryResult) error {
	file, err := os.Open("/data/amass_graphistry.json")
	if err != nil {
		log.Panicf("failed to open result file: %v", err)
		return err
	}
	defer file.Close()

	value, _ := ioutil.ReadAll(file)

	return json.Unmarshal(value, &result)
}

func (*VizService) GetGraphistry(ctx context.Context, req *pb.VizRequest) (*pb.VizResponse, error) {
	id := req.GetId()
	domain := req.GetDomain()

	err := runVizCommand(ctx, domain, &pb.VizConfig{
		Graphistry: true,
	})

	if err != nil {
		log.Panicf("failed to run viz: %v", err)
		return nil, err
	}

	var result pb.GraphistryResult
	err = readResult(&ctx, &result)
	if err != nil {
		log.Panicf("failed to read result file: %v", err)
		return nil, err
	}

	return &pb.VizResponse{
		Id:     id,
		Domain: &domain,
		Result: &pb.VizResponse_GResult{
			GResult: &result,
		},
	}, err
}
