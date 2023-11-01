package views

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/controllers"
)

func MainMenu() {
	fmt.Println("MENU")
	fmt.Println("[1] Transfer")
	fmt.Println("[2] Deposit")
	fmt.Println("[3] Withdraw")
	fmt.Println("[4] Account")
	fmt.Println("[0] Exit")

	controllers.UserInput("Type menu number: ", &controllers.Input)

	switch controllers.Input {
	case "1":
		TransferMenu()
	case "2":
		DepositMenu()
	case "3":
		WithdrawMenu()
	case "4":
		AccountMenu()
	case "0":
		ExitMenu()
	default:
		controllers.WrongInput()
		MainMenu()
	}
}
