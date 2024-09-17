package main

import "sync"

type Account struct {
	number int
	money int
	Lock sync.Locker
}

func NewAccount(number int, money int) *Account {
	return &Account{
		number: number,
		money: money,
		Lock: NewSpinlock(), // &sync.Mutex{},
	}
}

func (a *Account) transact(amount int) bool {
	if a.money + amount < 0 {
		return false
	}

	a.money += amount
	return true
}

func (a *Account) DepositTo(account *Account, amount int) bool {
	if amount <= 0 || a.number == account.number {
		return false
	}

	firstLock := a.Lock
	secondLock := account.Lock

	if account.number < a.number {
		firstLock = account.Lock
		secondLock = a.Lock
	}

	firstLock.Lock()
	secondLock.Lock()
	defer secondLock.Unlock()
	defer firstLock.Unlock()
	if retrieve_succeed := a.transact(amount *-1); retrieve_succeed {
		return account.transact(amount)
	}

	return false
}