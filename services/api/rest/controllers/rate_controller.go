package controllers

import (
	"encoding/json"
	"genesis-test-task/services/api/grpc/client/currency"
	"genesis-test-task/services/currency/rate/messages/proto"
	"net/http"
)

var currencyGRPCClient currency.CurrencyGRPCClient

func init() {
	currencyGRPCClient = currency.CurrencyGRPCClient{}
}

func GetRate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := currencyGRPCClient.GetRate(proto.RateRequest{BaseCurrency: "bitcoin", TargetCurrency: "uah"})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(response.Rate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
