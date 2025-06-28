package main

//manages accounts and generates revenue for them

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

var managedAccounts *[]account
var activePromo string

func loadAccounts() {
	rawAccs := os.Getenv("ACCOUNTS")
	var accounts []account
	unmarshalerr := json.Unmarshal([]byte(rawAccs), &accounts)
	if unmarshalerr != nil {
		log.Fatal("failed unmarshaling acccounts: ", unmarshalerr)
	}
	managedAccounts = &accounts
}

func fillAccounts() float64 { //🤑🤑🤑
	var moneyEarned float64
	newAccs := []account{}
	for _, acc := range *managedAccounts {
		moneyEarned += claimDailyReward(acc.Token)
		if acc.latestClaimedPromo != activePromo {
			moneyEarned += claimCurrentPromo(acc.Token)
			acc.latestClaimedPromo = activePromo
		}
		newAccs = append(newAccs, acc)
	}
	managedAccounts = &newAccs

	return moneyEarned
}

func startFillingAccounts() {
	for {
		time.Sleep(3 * time.Hour)
		fillAccounts()
	}
}

func claimDailyReward(token string) float64 {
	//daily reward
	_, err := makeRequest("POST", "https://rugplay.com/api/rewards/claim", nil, token)
	//TODO: get earned money

	if err != nil {
		return 0
	}

	return 0
}

func claimCurrentPromo(token string) float64 {
	_, err := makeRequest("POST", "https://rugplay.com/api/promo/verify", "{code: \"mustard\"}", token)
	//TODO: get earned money

	if err != nil {
		return 0
	}

	return 0
}
