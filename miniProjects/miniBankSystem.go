package miniProjects

import (
	"fmt"
	"os"
)

type BankStruct struct {
	balances   int
	operations int
}

var Denominations = []int{1000, 500, 100, 50, 20, 10, 5, 2, 1}

func BankSystem(bs *BankStruct) {
	balances := &BankStruct{balances: 5000} //Initial balance

	for {

		fmt.Println("1. Withdraw")
		fmt.Println("2. Balance")
		fmt.Println("3. Deposit")
		fmt.Println("4. Exit")

		fmt.Print("Choose Operations: ")
		fmt.Scanln(&bs.operations)
		switch bs.operations {
		case 1:
			Withdraw(balances)
		case 2:
			Balance(balances)
		case 3:
			Deposit(balances)
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid operation")
		}
	}
}

func Withdraw(bs *BankStruct) {
	var withdraw int
	fmt.Print("Enter Ammount to Withdraw: ")
	fmt.Scanln(&withdraw)

	if withdraw > bs.balances {
		fmt.Println("Insufficient balance")
		return
	}

	bs.balances -= withdraw
	fmt.Println("New Balances: ", bs.balances)

	for _, denomination := range Denominations {
		bill := bs.balances / denomination
		fmt.Printf("%d bills of %d\n", bill, denomination)
		bs.balances %= denomination
	}

}

func Balance(bs *BankStruct) {
	fmt.Println("Balance: ", bs.balances)

	for _, denomination := range Denominations {
		bill := bs.balances / denomination
		fmt.Printf("%d bills of %d\n", bill, denomination)
		bs.balances %= denomination
	}
}

func Deposit(bs *BankStruct) {
	var deposit int

	fmt.Print("Enter Ammount to Deposit: ")
	fmt.Scanln(&deposit)

	bs.balances += deposit
	fmt.Println("New Balances: ", bs.balances)
}
