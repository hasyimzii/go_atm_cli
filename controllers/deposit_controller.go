package controllers

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/repository"
)

func Deposit(amount int) error {
	for index, account := range Accounts {
		if account.Id == LoginId {
			(&Accounts[index]).Balance += amount
		}
	}

	accountRepository := repository.NewAccountRepository()

	err := accountRepository.Update(Accounts)
	if err != nil {
		return err
	}

	fmt.Println("Success deposit Rp", amount)
	return nil
}
