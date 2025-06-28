package main

//type defintions

type tradeResponse struct {
	Success    bool    `json:"success"`
	NewBalance float64 `json:"newBalance"`
}

type account struct {
	Name               string
	Token              string
	CoverCoin          string //name of coin used to cover transaction
	latestClaimedPromo string
}

type coinEntry struct {
	Symbol string `json:"symbol"`
  Price float64
}
