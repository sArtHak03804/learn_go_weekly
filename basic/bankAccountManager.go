package main

import "fmt"

type Account struct {
	accountNumber string
	balance       float64
}

func deposit(account *Account, amount float64) {
	account.balance += amount
	fmt.Printf("Deposited %.2f. New balance: %.2f\n", amount, account.balance)
}

func withdraw(account *Account, amount float64) {
	if amount > account.balance {
		fmt.Println("Insufficient balance!")
	} else {
		account.balance -= amount
		fmt.Printf("Withdrew %.2f. New balance: %.2f\n", amount, account.balance)
	}
}

func checkBalance(account *Account) {
	fmt.Printf("Current balance: %.2f\n", account.balance)
}

func bankAccountManager() {
	account := Account{accountNumber: "1234567890", balance: 1000.00}
	var choice int
	var amount float64

	for {
		fmt.Println("\nBank Account Manager:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter amount to deposit: ")
			fmt.Scanln(&amount)
			deposit(&account, amount)
		case 2:
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scanln(&amount)
			withdraw(&account, amount)
		case 3:
			checkBalance(&account)
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}
