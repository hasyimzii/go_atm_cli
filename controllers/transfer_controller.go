package controllers

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/models"
	"github.com/hasyimzii/go_atm_cli/repositories"
	"github.com/sirupsen/logrus"
)

func Transfer(id string, amount int) error {
	// decrease account balance
	for index, account := range Accounts {
		if account.Id == LoginId {
			if amount > account.Balance {
				msg := "Transfer failed, insufficient balance!"
				logrus.Warn("[Account: ", LoginName, "] ", msg)
				return models.ValidationError(msg)
			} else {
				(&Accounts[index]).Balance -= amount
				break
			}
		}
	}

	// increase target balance
	var targetName string
	for index, account := range Accounts {
		if account.Id == id {
			(&Accounts[index]).Balance += amount
			targetName = account.Name

			// update data
			accountRepository := repositories.NewAccountRepo()
			err := accountRepository.Update(Accounts)
			if err != nil {
				return err
			}

			msg := fmt.Sprint("Success transfer Rp ", amount, " to ", targetName)
			logrus.Info("[Account: ", LoginName, "] ", msg)
			fmt.Println(msg)
			return nil
		}
	}

	return models.NotFoundError("Transfer failed, id not found!")
}
