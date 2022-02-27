package enum

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	pb "github.com/kofeebrian/capamass/protos/amass/enum"
)

type EnumServer struct {
	pb.UnimplementedEnumServiceServer
}

func runEnumCommand(ctx *context.Context, domain string, config *pb.EnumConfig) error {
	cmd := exec.Command("amass", "enum")
	cmd.Stdout = os.Stdout // debug
	cmd.Args = append(cmd.Args, "-d", domain)

	cmd.Args = append(cmd.Args, "-dir", "/.config/amass")

	/* Timeout */
	if timeout := config.GetTimeout(); timeout > 0 {
		cmd.Args = append(cmd.Args, "-timeout", strconv.FormatUint(uint64(timeout), 10)) // set timeout in min.
	} else {
		cmd.Args = append(cmd.Args, "-timeout", "10") // default timeout = 10 mins
	}

	/* Mode */
	switch mode := config.GetMode(); mode {
	case pb.EnumConfig_ACTIVE:
		cmd.Args = append(cmd.Args, "-config", "utils/config/active.ini")
	case pb.EnumConfig_PASSIVE:
		cmd.Args = append(cmd.Args, "-config", "utils/config/passive.ini")
	case pb.EnumConfig_DEFAULT:
	default:
		cmd.Args = append(cmd.Args, "-config", "utils/config/default.ini")
	}

	/* DNS Resolver */
	if resolvers := config.GetDnsResolvers(); resolvers != nil {
		cmd.Args = append(cmd.Args, "-r", strings.Join(resolvers, ","))
	}

	if config.GetMode() != pb.EnumConfig_PASSIVE {
		cmd.Args = append(cmd.Args, "-ip") // show ip
	}

	cmd.Args = append(cmd.Args, "-src") // show source
	// cmd.Args = append(cmd.Args, "-silent") // do not print output to stdout

	return cmd.Run()
}

func (*EnumServer) Run(ctx context.Context, req *pb.EnumRequest) (*pb.EnumResponse, error) {
	id := req.GetId()
	domain := req.GetDomain()
	config := req.GetConfig()

	err := runEnumCommand(&ctx, domain, config)
	if err != nil {
		log.Panicf("failed to enumerate: %v", err)
		return nil, err
	}

	return &pb.EnumResponse{
		Id:     id,
		Domain: domain,
	}, nil
}
