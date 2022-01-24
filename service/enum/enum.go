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

func enumerate(name string) error {
	cmd := exec.Command("amass", "enum", "-d", name)
	cmd.Args = append(cmd.Args, "-timeout", "5")                       // set timeout 2 min.
	cmd.Args = append(cmd.Args, "-config", "utils/config/default.ini") // use config TODO: user can choose
	cmd.Args = append(cmd.Args, "-ip")                                 // show ip
	cmd.Args = append(cmd.Args, "-json", fmt.Sprintf("/data/%s.json", name))          

  return cmd.Run()
}

func extractResult() (*[]EnumResult, error) {
  f, err := os.Open("/data/result.json")
  if err != nil {
    log.Panicf("failed to open result file: %v", err)
    return nil, err
  }

  var results []EnumResult
  scanner := bufio.NewScanner(f)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    m := scanner.Text()
    var result EnumResult
    err := json.Unmarshal([]byte(m), &result)
    if err != nil {
      log.Panicf("failed to parse json: %v", err)
      continue
    }
    results = append(results, result)
  }
  return &results, nil
}

func (*EnumServer) BasicEnumerate(ctx context.Context, req *pb.EnumRequest) (*pb.EnumResponse, error) {
	log.Printf("starting enumeration...")

	err := enumerate(req.DomainName)
	if err != nil {
		log.Panicf("failed to enumerate: %v", err)
		return &pb.EnumResponse{
			Result: "failed to enumerate",
		}, err
	}

  results, err := extractResult()
  if err != nil {
    log.Panicf("extract results failed: %v", err)
    return &pb.EnumResponse{
      Result: "failed to extract results",
    }, err
  }

  fmt.Println(results)

	return &pb.EnumResponse{
		Result: "Test Success",
	}, nil
}
