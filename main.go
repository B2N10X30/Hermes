package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

func BankApp() {
	var (
		//User details and fixed Balance
		userName        string
		balance         float64
		option          string
		recipient       string
		bankName        string
		amount          float64
		continum        string
		checker         string
		usrName         = "^[a-zA-Z0-9]{5,10}$"
		userCheker, err = regexp.Compile(usrName)
	)
	if err != nil {
		log.Println("System error:", err)
		return
	}
	fmt.Println("Welcome to Hermes bank , Thank YOU for Banking with us.")
	for {
		fmt.Print("username: ")
		fmt.Scanf("%s", &userName)
		if userCheker.MatchString(userName) {
		start:
			fmt.Println("Hello, ", userName)
			balance = 100
			fmt.Printf("Available Balance: $%f\n", balance)
			fmt.Println("||Select an option for a Transaction||\n ------------------------------------")
			fmt.Println(" [T]\t||   transfer")
			fmt.Println(" [W]\t||   withdrawal")
			fmt.Println(" [D]\t||   deposit")
			fmt.Println(" [E]\t||   exit")

			fmt.Printf("Input an option: ")
			fmt.Scanf("%s", &option)

			switch option {
			case "T":
				for {
					fmt.Print("Transfer: press T or t to continue, Press E or any key to exit: ")
					fmt.Scanf("%s", &continum)
					if continum == "T" || continum == "t" {
						fmt.Print("Recipient Account: ")
						fmt.Scanf("%s", &recipient)
						fmt.Print("Bank Name: ")
						fmt.Scanf("%s", &bankName)
						fmt.Print("Amount: ")
						fmt.Scanf("%f", &amount)
						fmt.Printf("You have decided to transfer $%f to %s,%s. Press C to continue: ", amount, recipient, bankName)
						fmt.Scanf("%s", &checker)
						if checker == "C" || checker == "c" {
							if amount < balance {
								balance -= amount
								fmt.Println("Processing...")
								time.Sleep(5 * time.Second)
								fmt.Println("Transaction Successful.")
								fmt.Printf("Cleared Balance: $%f\n", balance)
							} else {
								fmt.Println("Insufficient Balance to complete Transaction. Cancelling Transaction...")
								time.Sleep(5 * time.Second)
							}
						} else {
							fmt.Println("Cancelling Transaction...")
							goto start
						}
					} else {
						goto start
					}
				}

			case "W":
				for {
					fmt.Println("Withdrawal: Press W or w to continue, Press E or any key to exit")
					fmt.Scanf("%s", &continum)
					if continum == "W" || continum == "w" {
						fmt.Print("Amount: ")
						fmt.Scanf("%f", &amount)
						fmt.Printf("You have decided to withdraw $%f from your account. Press C to continue: ", amount)
						fmt.Scanf("%s", &checker)
						if checker == "C" || checker == "c" {
							if amount < balance {
								balance -= amount
								fmt.Println("Processing")
								time.Sleep(5 * time.Second)
								fmt.Println("Transaction Successful.")
								fmt.Printf("Cleared Balance: $%f\n", balance)
								fmt.Printf("Collect your cash,Card and scram ;¬]\n")
							} else {
								fmt.Println("Insufficient Balance to complete Transaction. Cancelling Transaction...")
							}
						} else {
							fmt.Println("Cancelling Transaction...")
							goto start
						}
					} else {
						goto start
					}
				}

			case "D":
				for {
					fmt.Println("Deposit: Press D or d to continue, Press E or any key to exit: ")
					fmt.Scanf("%s", &continum)
					if continum == "D" || continum == "d" {
						fmt.Printf("Amount: ")
						fmt.Scanf("%f", &amount)
						fmt.Printf("$%f will be Added to your Balance. Press c to continue: ", amount)
						if checker == "C" || checker == "c" {
							balance += amount
							fmt.Println("Processing...")
							time.Sleep(5 * time.Second)
							fmt.Printf("$%f has been added to your Account\n", amount)
							fmt.Printf("Cleared Balance: $%f\n", balance)
						} else {
							fmt.Println("Cancelling...")
							time.Sleep(5 * time.Second)
							goto start
						}
					} else {
						fmt.Println("Cancelling...")
						time.Sleep(5 * time.Second)
						goto start
					}
				}

			default:
				fmt.Println("Thank you for Banking with us. Exiting...")
				time.Sleep(5 * time.Second)
				break
			}
		} else {
			fmt.Println("Please Provide a valid username!(not less than 5 combinations of letters and numbers)")
		}
	}
}

func main() {
	BankApp()
}
