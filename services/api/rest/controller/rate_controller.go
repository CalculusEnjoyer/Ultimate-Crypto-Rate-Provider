package controller

import (
	"api/domain"
	"api/rest"
	"context"
	"net/http"
)

type RateErrorPresenter interface {
	PresentHTTPErr(err error, w http.ResponseWriter)
}

type RatePresenter interface {
	SuccessfulRateResponse(w http.ResponseWriter, response domain.RateResponse)
}

type RateService interface {
	GetRate(request domain.RateRequest, cnx context.Context) (*domain.RateResponse, error)
}

type RateController struct {
	rateService  RateService
	errPresenter RateErrorPresenter
	presenter    RatePresenter
}

func NewRateController(rateService RateService, errPresenter RateErrorPresenter, presenter RatePresenter) *RateController {
	return &RateController{rateService: rateService, errPresenter: errPresenter, presenter: presenter}
}

func (rc *RateController) GetBTCRate(w http.ResponseWriter, r *http.Request) {
	response, err := rc.rateService.GetRate(domain.RateRequest{BaseCurrency: domain.BTC, TargetCurrency: domain.UAH}, r.Context())
	if err != nil {
		rc.errPresenter.PresentHTTPErr(err, w)
		return
	}

	rc.presenter.SuccessfulRateResponse(w, *response)
}

func (rc *RateController) GetRate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		rc.errPresenter.PresentHTTPErr(err, w)
		return
	}

	target := r.Form.Get(rest.KeyTargetCurrency)
	base := r.Form.Get(rest.KeyBaseCurrency)

	response, err := rc.rateService.GetRate(domain.RateRequest{
		BaseCurrency:   domain.Currency(base),
		TargetCurrency: domain.Currency(target)}, r.Context())
	if err != nil {
		rc.errPresenter.PresentHTTPErr(err, w)
		return
	}

	rc.presenter.SuccessfulRateResponse(w, *response)
}
