package tools

import (
	"time"
)

type mockDb struct {}

var mockLoginDetails = map[string]*LoginDetails{
	"alex": {
		AuthToken: "ABC123",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "ABC456",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "ABC789",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]*CoinDetails{
	"alex": {
		Coins:   100,
		Username: "alex",
	},
	"jason": {
		Coins:   300,
		Username: "jason",
	},
	"marie": {
		Coins:   200,
		Username: "marie",
	},
}

func (db *mockDb) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate a delay for database access
	time.Sleep(time.Second * 1)

	if details, exists := mockLoginDetails[username]; exists {
		return details
	}
	return nil
}

func (db *mockDb) GetUserCoins(username string) *CoinDetails {
	// Simulate a delay for database access
	time.Sleep(time.Second * 1)

	if details, exists := mockCoinDetails[username]; exists {
		return details
	}
	return nil
}

func (db *mockDb) SetupDatabase() error {
	return nil
}