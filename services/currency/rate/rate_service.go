package rate

import (
	"currency.com/currency/rate/messages"
)

type RateService interface {
	GetRate(currencies messages.RateRequest) (rate messages.RateResult, err error)
}
