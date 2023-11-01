package views

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/controllers"
	"github.com/hasyimzii/go_atm_cli/models"
)

func AccountMenu() {
	var account models.Account = models.Account{Id: "0001", Name: "Asd", Balance: 10000}

	fmt.Println("Your Account")
	fmt.Println("Id:", account.Id)
	fmt.Println("Name:", account.Name)
	fmt.Println("Balance: Rp", account.Balance)

	controllers.UserInput("[Press enter to go back] ", &controllers.Input)
	MainMenu()
}
