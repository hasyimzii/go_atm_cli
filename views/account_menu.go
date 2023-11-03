package views

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/controllers"
)

func AccountMenu() {
	account := controllers.GetAccount()

	fmt.Println("Your Account")
	fmt.Println("Id:", account.Id)
	fmt.Println("Name:", account.Name)
	fmt.Println("Balance: Rp", account.Balance)

	controllers.UserInput("[Press enter to go back] ", &controllers.Input)
	MainMenu()
}
