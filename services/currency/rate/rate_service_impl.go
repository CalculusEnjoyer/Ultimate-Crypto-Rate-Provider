package rate

import (
	"genesis-test-task/services/currency/rate/messages"
	"genesis-test-task/services/currency/rate/providsers"
	"time"
)

var Provider = providsers.CoinGeckoRateProvider{}

type rateService struct{}

func NewService() RateService { return &rateService{} }

func (r *rateService) GetRate(currencies messages.RateRequest) (rate messages.RateResult, err error) {
	btcRate, err := Provider.GetExchangeRate(currencies.BaseCurrency, currencies.TargetCurrency)

	return messages.RateResult{Rate: btcRate, Timestamp: time.Now().String()}, err
}
