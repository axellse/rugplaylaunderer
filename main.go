package main

//entry point

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func makeTrade(action string, coin string, amount float64, token string) error {
	body := strings.NewReader("{\"type\":\"" + action + "\",\"amount\":" + strconv.FormatFloat(amount, 'g', -1, 64) + "}") //ik i should probably do json marshaling here but lazy
	req, err := http.NewRequest("GET", "https://rugplay.com/api/coin/"+coin+"/trade", body)
	req.AddCookie(&http.Cookie{
		Name:  "__Secure-better-auth.session_token",
		Value: token,
	})
	if err != nil {
		return err
	}

	resp, reqerr := http.DefaultClient.Do(req)
	if reqerr != nil {
		return reqerr
	}

	ba, rderr := io.ReadAll(resp.Body)
	if rderr != nil {
		return rderr
	}

	var jsonresp tradeResponse
	jsonerr := json.Unmarshal(ba, &jsonresp)
	if jsonerr != nil {
		return jsonerr
	}

	if !jsonresp.Success {
		return errors.New("request response did not indicate success")
	}
	return nil
}

func main() {
	fmt.Println("welcome to the program! 💸💸💸")
	fmt.Println("note: this program is used to generate income and launder money on the platform rugplay, created by the youtuber facedev. it's an online playground for investments that does not involve real money at all, everything is made up.")
	fmt.Println("----------------------------------------")
	fmt.Print("loading accounts... ")
	loadAccounts()
	fmt.Println("done!")
	fmt.Print("filling accounts... ")
	earned := fillAccounts()
	go startFillingAccounts()
	fmt.Println("done! (earned " + strconv.FormatFloat(earned, 'g', -1, 64) + " across all accounts 🤑🤑🤑)")


}
