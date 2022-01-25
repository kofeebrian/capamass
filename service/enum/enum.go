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

type Address struct {
	IP   string `json:"ip"`
	Cidr string `json:"cidr"`
	ASN  uint32 `json:"asn"`
	Desc string `json:"desc"`
}

type EnumResult struct {
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	Addresses []Address `json:"addresses"`
	Tag       string    `json:"cert"`
	Sources   []string  `json:"srouces"`
}

type EnumServer struct {
	pb.UnimplementedEnumServiceServer
}

var (
	stdout, stderr bytes.Buffer
)

func enumerate(names []string) error {
	cmd := exec.Command("amass", "enum")
  for _, name := range names {
    log.Println(name)
    cmd.Args = append(cmd.Args, "-d", name)                       
  }
	cmd.Args = append(cmd.Args, "-timeout", "5")                       // set timeout 2 min.
	cmd.Args = append(cmd.Args, "-config", "utils/config/default.ini") // use config TODO: user can choose
	cmd.Args = append(cmd.Args, "-ip")                                 // show ip
	cmd.Args = append(cmd.Args, "-json", fmt.Sprintf("/data/save.json", ))          

  return cmd.Run()
}

func extractResult() (*[]*pb.EnumResponse_Result, error) {
  f, err := os.Open("/data/save.json")
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

	err := enumerate(req.Domains)
	if err != nil {
		log.Panicf("failed to enumerate: %v", err)
		return &pb.EnumResponse{}, err
	}

  results, err := extractResult()
  if err != nil {
    log.Panicf("extract results failed: %v", err)
    return &pb.EnumResponse{}, err
  }

	return &pb.EnumResponse{
    Results: *results,
	}, nil
}
