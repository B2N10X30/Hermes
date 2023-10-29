package main

import (
	"fmt"
	"regexp"
)

func SLice(my_slice []int) {
	// for index-value pair in range of items in
	for _, v := range my_slice {
		fmt.Printf("%d\n", v+4)
	}
}

func BankApp() {
	var (
		//User details and fixed Balance
		userName string
		balAnce  float64
		option   string
	)
	//REGEX TO CHECK FOR USERNAME .¬) doesn't work though
	usr_name_ := "^[a-zA-Z0-9]{5,10}$"
	user_cheker, err := regexp.Compile(usr_name_)
	if err != nil {
		fmt.Println("System error:", err)
		return
	}
	fmt.Println("Welcome to Hermes bank, Thanks for Banking with us")
	fmt.Print("username: ")
	fmt.Scanf("%s", &userName)
	if user_cheker.MatchString(userName) {
		fmt.Println("Please Provide a valid username!(not less than 5 combinations of letters and numbers)")
		return
	} else {
		for {
			fmt.Println("Hello, ", userName)
			balAnce = 100

			fmt.Printf("Available Balance: $%f\n", balAnce)
			fmt.Println("||Select the corresponding option for a Transaction||\n ------------------------------------------------ ")
			fmt.Println(" [T]\t||   transfer")
			fmt.Println(" [W]\t||   withdrawal")
			fmt.Println(" [D]\t||   deposit")

			fmt.Printf("Input an option: ")
			fmt.Scanf("%s", &option)
			//functions
			if option == "T" {
				var cont string
				fmt.Println("to continue this Transaction press T or any character to exit : ")
				fmt.Scanf("%s", &cont)
				if cont == "T" {
					var (
						Recipient string
						Bank_name string
						Amount    float64
					)
					fmt.Println("recipient account: ")
					fmt.Scanf("%s", &Recipient)
					fmt.Println("Bank Name: ")
					fmt.Scanf("%s", &Bank_name)
					fmt.Print("Amount: ")
					fmt.Scanf("%f", &Amount)
					balAnce -= Amount
					fmt.Printf("$%f has been transfered to %s\n", Amount, Recipient)
					fmt.Printf("Cleared Balance: $%f", balAnce)
				} else {
					fmt.Println("exiting...")
					return
				}

			}
			if option == "W" {
				var cont string
				fmt.Println("to continue this Transaction press W or any character to exit : ")
				fmt.Scanf("%s", &cont)
				if cont == "W" {
					var (
						Amount float64
					)
					fmt.Print("Amount: ")
					fmt.Scanf("%f", &Amount)
					if Amount < balAnce {
						balAnce -= Amount
						fmt.Printf("$%f has been withdrawn from your Balance", Amount)
						fmt.Printf("Cleared Balance: $%f", balAnce)
					} else {
						fmt.Println("Transaction Amount is greater than your Available balance, Top up your Balance to continue")
					}
				} else {
					fmt.Println("exiting...")
					return
				}
			}
			if option == "D" {
				var cont string
				fmt.Println("to continue this Transaction press D or any character to exit : ")
				fmt.Scanf("%s", &cont)
				if cont == "D" {
					var (
						Amount float64
					)
					//Amount = 500
					fmt.Print("Amount to deposit: ")
					fmt.Scanf("%f", &Amount)
					balAnce += Amount
					fmt.Printf("$%f has be added to your account\n", Amount)
					fmt.Printf("Cleared Balance: $%f", balAnce)
				} else {
					fmt.Println("exiting...")
					return
				}
			} else {
				return
			}
		}
	}

}

func main() {
	//SLice([]int{1, 2, 3})
	BankApp()
}
