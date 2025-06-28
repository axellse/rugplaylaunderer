package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"time"
)

var deadCoinsCache []string //list of coin symbols
var deadCoinsCacheExpiry int64 //unix time

func concealTransactions(acc account) {
  fmt.Print("picking coins to launder for " + acc.Name + "...")
  
  if time.Now().Unix() > deadCoinsCacheExpiry {
    resp, err := makeRequest("GET", "https://rugplay.com/api/market?search=&sortBy=volume24h&sortOrder=asc&priceFilter=all&changeFilter=all&page=1&limit=80", nil, acc.Token)
    if err == nil {return}
    
    var fresp []coinEntry
    jerr := json.Unmarshal(resp, &fresp)
    if jerr != nil {return}

    deadCoinsCache = []string{}
    for _, ce := range fresp {
      deadCoinsCache = append(deadCoinsCache, ce.Symbol)
    }
    deadCoinsCacheExpiry = time.Now().Unix()
  }
  rand.Intn(len(deadCoinsCache))
  
}