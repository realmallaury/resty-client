package main

import (
	"fmt"
	"log"
	"os"

	"github.com/realmallaury/resty-client/cmd/client"
	"github.com/realmallaury/resty-client/cmd/account"
)

func main() {
	logger := log.New(os.Stdout, "Demo: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	accountRestClient, err := client.New("http://localhost:8080")
	if err != nil {
		logger.Fatalf("Cannot initialize rest client: %v", err)
	}

	account := account.GetTestCreateAccount()

	createdAccount, err := accountRestClient.Create(account)
	if err != nil {
		logger.Fatalf("Cannot create account: %v", err)
	}

	fmt.Printf("Account:\n %+v\n\n", account)
	fmt.Printf("Created account:\n %+v\n\n", createdAccount)

	// account, err = accountRestClient.Fetch("cd27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	// if err != nil {
	// 	logger.Fatalf("Cannot fetch account data: %v", err)
	// }

	// fmt.Printf("Account:\n %+v\n\n", account)

	// err = accountRestClient.Delete("cd27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)
	// if err != nil {
	// 	logger.Fatalf("Cannot delete account: %v", err)
	// }

	// for i := 1; i <= 10; i++ {
	// 	account := account.GetTestCreateAccount()
	// 	account.ID = uuid.New().String()

	// 	_, err := accountRestClient.Create(account)
	// 	if err != nil {
	// 		logger.Fatalf("Cannot create account: %v", err)
	// 	}
	// }

	// accounts, err := accountRestClient.List(0, 5)
	// if err != nil {
	// 	logger.Fatalf("Cannot list accounts: %v", err)
	// }

	// fmt.Printf("Account:\n %+v\n\n", accounts)

	// accounts, err = accountRestClient.List(1, 5)
	// if err != nil {
	// 	logger.Fatalf("Cannot list accounts: %v", err)
	// }

	// fmt.Printf("Account:\n %+v\n\n", accounts)
}
