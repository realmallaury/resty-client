package main

import (
	"log"
	"os"

	rc "github.com/push-er/resty-client/cmd/client"
)

func main() {
	logger := log.New(os.Stdout, "Demo: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	accountRestClient, err := rc.New("http://localhost:8080")
	if err != nil {
		logger.Fatalf("Cannot initialize rest client: %v", err)
	}

	rsp, err := accountRestClient.Fetch("test")
	if err != nil {
		logger.Fatalf("Cannot fetch account data: %v", err)
	}
	logger.Print(rsp)
	logger.Print(rsp.Status())
}
