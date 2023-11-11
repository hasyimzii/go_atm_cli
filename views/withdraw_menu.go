package views

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/controllers"
)

func WithdrawMenu() {
	var amount int
	var pin string

	fmt.Print("Insert amount you want to withdraw: Rp ")
	fmt.Scanln(&amount)

	fmt.Print("Insert your pin: ")
	fmt.Scanln(&pin)

	login := controllers.Login(controllers.LoginId, pin)

	if login {
		err := controllers.Withdraw(amount)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Wrong pin!")
	}

	controllers.UserInput("[Press enter to go back] ", &controllers.Input)
	MainMenu()
}
