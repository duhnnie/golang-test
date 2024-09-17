package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

const accounts_number = 10000
const initialMoney = 100
const threads = 4

var transactions_processed int32 = 0
var accounts []*Account

func randomTransactions() {
	for {
		originIndex := rand.Intn(accounts_number)
		destinyIndex := rand.Intn(accounts_number)
		amount := rand.Intn(100)

		for originIndex == destinyIndex {
			destinyIndex = rand.Intn(accounts_number)
		}

		originAccount := accounts[originIndex]
		destinyAccount := accounts[destinyIndex]

		originAccount.DepositTo(destinyAccount, amount)
		atomic.AddInt32(&transactions_processed, 1)
	}
}

func main() {
	for i := 0; i < accounts_number; i++ {
		newAccount := NewAccount(i + 1, initialMoney)
		accounts = append(accounts, newAccount)
	}

	for i := 0; i < threads; i++ {
		go randomTransactions()
	}

	for {
		sum := 0

		for i := 0; i < len(accounts); i++ {
			accounts[i].Lock.Lock()
		}

		for i := 0; i < len(accounts); i++ {
			sum += accounts[i].money
		}

		for i := 0; i < len(accounts); i++ {
			accounts[i].Lock.Unlock()
		}

		fmt.Printf("Money in bank: %d - transactions: %d\n", sum, transactions_processed)
		time.Sleep(2 * time.Second)
	}
}