package main

import (
	"context"
	"log"
	"sync"

	enumpb "github.com/kofeebrian/capamass/protos/amass/enum"
	"google.golang.org/grpc"
)

func main() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
  conn, err := grpc.Dial("enum_service:3000", opts...)
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
				DomainName: "chula.ac.th",
			}

			res, err := c.BasicEnumerate(ctx, &req)
			if err != nil {
				log.Fatalf("error to enumerate: %v", err)
			}

			log.Printf("iter: %d finished with response: %v", i, res)

		}(i)
	}

	wg.Wait()

}
