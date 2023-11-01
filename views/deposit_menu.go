package views

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/controllers"
)

func DepositMenu() {
	var amount int
	var pin string

	fmt.Print("Insert amount you want to deposit: Rp ")
	fmt.Scanln(&amount)

	fmt.Print("Insert your pin: ")
	fmt.Scanln(&pin)

	// check pin & increase balance

	fmt.Println("Success deposit Rp", amount)

	controllers.UserInput("[Press enter to go back] ", &controllers.Input)
	MainMenu()
}
