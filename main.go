package main

import (
	"fmt"
	"genesis-test-task/services/currency/rate/sources"
)

func main() {
	rateProvider := &sources.CoinGeckoRateProvider{}
	rate, err := rateProvider.GetExchangeRate("bitcoin", "UAH")
	if err != nil {
		fmt.Printf("Failed to get exchange rate: %v\n", err)
		return
	}

	fmt.Printf("Exchange rate: %f\n", rate)
}
