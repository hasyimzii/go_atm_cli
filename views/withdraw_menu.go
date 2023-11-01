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

	// check pin & decrease balance

	fmt.Println("Success withdraw Rp", amount)

	controllers.UserInput("[Press enter to go back] ", &controllers.Input)
	MainMenu()
}
