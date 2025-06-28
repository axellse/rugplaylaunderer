package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var deadCoinsCache []coinEntry    //list of coin symbols
var deadCoinsCacheExpiry int64 //unix time

func concealTransactions(acc account) {
	fmt.Print("picking coins to launder for " + acc.Name + "...")

	if time.Now().Unix() > deadCoinsCacheExpiry {
		resp, err := makeRequest("GET", "https://rugplay.com/api/market?search=&sortBy=volume24h&sortOrder=asc&priceFilter=all&changeFilter=all&page=1&limit=80", nil, acc.Token)
		if err == nil {
			return
		}

		var fresp []coinEntry
		jerr := json.Unmarshal(resp, &fresp)
		if jerr != nil {
			return
		}

		deadCoinsCache = fresp
		deadCoinsCacheExpiry = time.Now().Unix() + (5 * 60)
	}

	pickedCoins := []string{}
  pickedCoinPrices := []float64{}
	for range 5 {
		coin := deadCoinsCache[rand.Intn(len(deadCoinsCache)-1)]
		pickedCoins = append(pickedCoins, coin.Symbol)
    pickedCoinPrices = append(pickedCoinPrices, coin.Price)
	}

	fmt.Println(" (picked " + strings.Join(pickedCoins, ", ") + ")")

  for i, coin := range pickedCoins {
    buyAmm := (rand.Float64() * 5) + 4
    fmt.Print("buying " + coin + " for " + strconv.FormatFloat(buyAmm, 'g', -1, 64) + "(no. " + strconv.Itoa(i) + ")... ")

    var err error = /*makeTrade("BUY", coin, buyAmm / pickedCoinPrices[i], acc.Token)*/ nil
    if err != nil {
      fmt.Println(" (failed!!!)")
    } else {
      fmt.Println(" (done!)")
    }
  }
}
