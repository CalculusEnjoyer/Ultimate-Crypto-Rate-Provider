package sources

const coingeckoUrl = "https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s"

func GetExchangeRate(baseCurrency, targetCurrency string) (rate float64, err error) {
	return 100500, nil
}
