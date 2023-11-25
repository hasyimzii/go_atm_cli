package repositories

import (
	"encoding/json"
	"os"

	"github.com/hasyimzii/go_atm_cli/models"
)

type AccountRepoInterface interface {
	FindAll() ([]models.Account, error)
	Update(accounts []models.Account) error
}

type accountRepository struct {
	file string
}

func NewAccountRepo() AccountRepoInterface {
	return &accountRepository{file: "./data/data.json"}
}

func (repo *accountRepository) FindAll() ([]models.Account, error) {
	file, err := os.Open(repo.file)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var accounts []models.Account
	err = json.NewDecoder(file).Decode(&accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (repo *accountRepository) Update(accounts []models.Account) error {
	file, err := os.Create(repo.file)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(accounts)
	if err != nil {
		return err
	}

	return nil
}
