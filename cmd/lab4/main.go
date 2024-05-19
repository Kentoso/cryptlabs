package main

import (
	"fmt"
	"github.com/Kentoso/cryptlabs/blockchain"
	"time"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock([]blockchain.Transaction{
		{"Alice", "Bob", 50, time.Now()},
		{"Bob", "Charlie", 30, time.Now()},
	})

	bc.AddBlock([]blockchain.Transaction{
		{"Charlie", "Alice", 20, time.Now()},
	})

	err := bc.SaveToFile("blockchain.json")
	if err != nil {
		fmt.Println("Error saving blockchain:", err)
		return
	}

	loadedBc, err := blockchain.LoadFromFile("blockchain.json")
	if err != nil {
		fmt.Println("Error loading blockchain:", err)
		return
	}

	valid, reason := loadedBc.IsValid()
	if valid {
		fmt.Println("Blockchain is valid")
	} else {
		fmt.Println("Blockchain is invalid:", reason)
	}

	balances, minBalances, maxBalances := loadedBc.GetAccountsBalances(len(loadedBc.Blocks) - 1)
	fmt.Println("Accounts balances:", balances)
	fmt.Println("Accounts min balances:", minBalances)
	fmt.Println("Accounts max balances:", maxBalances)
}
