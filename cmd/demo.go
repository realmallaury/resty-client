package main

import (
	"fmt"
	"log"
	"os"

	rc "github.com/push-er/resty-client/cmd/client"
	"github.com/push-er/resty-client/cmd/model"
)

func main() {
	logger := log.New(os.Stdout, "Demo: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	accountRestClient, err := rc.New("http://localhost:8080")
	if err != nil {
		logger.Fatalf("Cannot initialize rest client: %v", err)
	}

	account := model.GetTestCreateAccount()
	createdAccount, err := accountRestClient.Create(account)
	if err != nil {
		logger.Fatalf("Cannot create account: %v", err)
	}

	fmt.Printf("Account:\n %+v\n\n", account)
	fmt.Printf("Created account:\n %+v\n\n", createdAccount)

	account, err = accountRestClient.Fetch("cd27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	if err != nil {
		logger.Fatalf("Cannot fetch account data: %v", err)
	}

	fmt.Printf("Account:\n %+v\n\n", account)
}
