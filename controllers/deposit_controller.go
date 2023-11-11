package controllers

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/models"
	"github.com/hasyimzii/go_atm_cli/repository"
	"github.com/sirupsen/logrus"
)

func Deposit(amount int) error {
	for index, account := range Accounts {
		if account.Id == LoginId {
			(&Accounts[index]).Balance += amount

			// update data
			accountRepository := repository.NewAccountRepo()
			err := accountRepository.Update(Accounts)
			if err != nil {
				return err
			}

			msg := fmt.Sprint("Success deposit Rp ", amount)
			logrus.Info("[Account: ", LoginName, "] ", msg)
			fmt.Println(msg)
			return nil
		}
	}

	return models.NotFoundError("Deposit failed, id not found!")
}
