package controllers

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/repository"
)

func Transfer(id string, amount int) error {
	// decrease account balance
	for index, account := range Accounts {
		if account.Id == LoginId {
			if amount > account.Balance {
				fmt.Println("Transfer failed, insufficient balance!")
			} else {
				(&Accounts[index]).Balance -= amount
			}
		}
	}

	// increase target balance
	var targetName string
	for index, account := range Accounts {
		if account.Id == id {
			(&Accounts[index]).Balance += amount
			targetName = account.Name
		}
	}

	accountRepository := repository.NewAccountRepository()

	err := accountRepository.Update(Accounts)
	if err != nil {
		return err
	}

	fmt.Println("Success transfer Rp", amount, "to", targetName)
	return nil
}
