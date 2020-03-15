package main

import (
	"fmt"

	"github.com/taybur/Golang_practice/funding"
)

func main() {
	funding := funding.NewFund(10000)
	totalbalance := funding.BalanceAmmount()
	fmt.Println("Starting from here", totalbalance)
	funding.Withdraw(1000)
	newbalance := funding.BalanceAmmount()
	fmt.Println("new money after with from here", newbalance)

}
