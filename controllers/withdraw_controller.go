package controllers

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/repository"
)

func Withdraw(amount int) error {
	for index, account := range Accounts {
		if account.Id == LoginId {
			if amount > account.Balance {
				fmt.Println("Withdraw failed, insufficient balance!")
			} else {
				(&Accounts[index]).Balance -= amount
			}
		}
	}

	accountRepository := repository.NewAccountRepository()

	err := accountRepository.Update(Accounts)
	if err != nil {
		return err
	}

	fmt.Println("Success withdraw Rp", amount)
	return nil
}
