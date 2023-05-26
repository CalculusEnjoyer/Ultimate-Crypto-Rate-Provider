package rate

import (
	"genesis-test-task/services/currency/rate/messages"
	"genesis-test-task/services/currency/rate/sources"
)

type RateService interface {
	GetBtcUahRate(provider *sources.RateProvider)(rate messages.RateMessage)
}
