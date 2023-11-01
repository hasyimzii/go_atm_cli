package views

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/controllers"
)

func TransferMenu() {
	var amount int
	var id, pin string

	fmt.Print("Insert id you want to transfer: ")
	fmt.Scanln(&id)

	fmt.Print("Insert amount you want to transfer: Rp ")
	fmt.Scanln(&amount)

	fmt.Print("Insert your pin: ")
	fmt.Scanln(&pin)

	// check id, pin, decrease this balance, & increase other balance

	fmt.Println("Success transfer Rp", amount, "to", "name")

	controllers.UserInput("[Press enter to go back] ", &controllers.Input)
	MainMenu()
}