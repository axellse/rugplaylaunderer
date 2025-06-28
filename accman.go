package main

//manages accounts and generates revenue for them

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"
)

var managedAccounts []*account
var activePromo string

func loadAccounts() {
  rawAccs := os.Getenv("ACCOUNTS")
  var accounts []*account
  unmarshalerr := json.Unmarshal([]byte(rawAccs), accounts)
  if unmarshalerr != nil {
    panic("cant unmarshal accounts")
  }
  managedAccounts = accounts
}

func fillAccounts() { //🤑🤑🤑
	for _, acc := range managedAccounts {
		claimDailyReward(acc.Token)
    if acc.latestClaimedPromo != activePromo {
      claimCurrentPromo(acc.Token)
      acc.latestClaimedPromo = activePromo
    }
		
	}
}

func startFillingAccounts() {
  for {
    time.Sleep(3 * time.Hour)
    fillAccounts()
  }
}

func claimDailyReward(token string) {
	//daily reward
	req, err := http.NewRequest("POST", "https://rugplay.com/api/rewards/claim", nil)
	req.AddCookie(&http.Cookie{
		Name:  "__Secure-better-auth.session_token",
		Value: token,
	})

	if err != nil {
		return
	}
	http.DefaultClient.Do(req)
}

func claimCurrentPromo(token string) {
	body := strings.NewReader("{code: \"mustard\"}")
	promoreq, rcerr := http.NewRequest("POST", "https://rugplay.com/api/promo/verify", body)
	promoreq.AddCookie(&http.Cookie{
		Name:  "__Secure-better-auth.session_token",
		Value: token,
	})
	if rcerr != nil {
		return
	}

	http.DefaultClient.Do(promoreq)
}
