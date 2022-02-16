package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	enumpb "github.com/kofeebrian/capamass/protos/amass/enum"
	"google.golang.org/grpc"
)

var (
	name = flag.String("name", "", "domain name")
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to get .env: %v", err)
	}

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	flag.Parse()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	var wg sync.WaitGroup
	var c enumpb.EnumServiceClient

	c = enumpb.NewEnumServiceClient(conn)

	for i := 0; i < 1; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			req := enumpb.EnumRequest{
				Domain: *name,
			}

			res, err := c.Run(ctx, &req)
			if err != nil {
				log.Fatalf("error to enumerate: %v", err)
			}

			log.Printf("iter: %d finished with response: %v", i, res)

		}(i)
	}

	wg.Wait()

}
