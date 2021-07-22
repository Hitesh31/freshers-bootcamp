package main

import (
	"fmt"
	"sync"
)

var (
	mutx   sync.Mutex
	balance int
)

func deposit(val int, wg *sync.WaitGroup) {
	mutx.Lock()
	fmt.Printf("Depositing %d to account with\n", val)
	balance += val
	mutx.Unlock()
	wg.Done()
}

func withdraw(val int, wg *sync.WaitGroup) {
	mutx.Lock()
	if val > balance {
		fmt.Println("Error, Can't withdraw")
		mutx.Unlock()
		wg.Done()
		return
	}

	fmt.Printf("Withdrawing %d from account\n", val)
	balance -= val
	mutx.Unlock()
	wg.Done()
}

func main() {
	balance=500
	var wg sync.WaitGroup
	wg.Add(2)
	go deposit(500, &wg)
	go withdraw(700, &wg)
	wg.Wait()

	fmt.Printf("New Balance %d\n", balance)
}
