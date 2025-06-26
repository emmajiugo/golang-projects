package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database collection
type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockDb{}

	if err := database.SetupDatabase(); err != nil {
		log.Error("Failed to setup database: ", err)
		return nil, err
	}

	return &database, nil
}

