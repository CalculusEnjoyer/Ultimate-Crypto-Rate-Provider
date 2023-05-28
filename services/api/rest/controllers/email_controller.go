package controllers

import (
	"genesis-test-task/services/api/grpc/client/currency"
	"genesis-test-task/services/api/grpc/client/email"
	"genesis-test-task/services/api/grpc/client/storage"
	"genesis-test-task/services/api/rest/utils"
	rateProto "genesis-test-task/services/currency/rate/messages/proto"
	emailProto "genesis-test-task/services/email/dispatcher/messages/proto"
	"genesis-test-task/services/storage/emails/errors"
	"genesis-test-task/services/storage/emails/messages/proto"
	"net/http"
	"strconv"
)

var rateGRPCClient currency.CurrencyGRPCClient
var emailGRPCClient email.EmailGRPCClient
var storageGRPCClient storage.StorageGRPCClient

func init() {
	rateGRPCClient = currency.CurrencyGRPCClient{}
	emailGRPCClient = email.EmailGRPCClient{}
	storageGRPCClient = storage.StorageGRPCClient{}
}

func AddEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	if !utils.ValidateEmail(email) {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	_, err = storageGRPCClient.AddEmail(proto.AddEmailRequest{
		Email: email,
	})

	if err == errors.EmailAlreadyExist {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	return
}

func SendEmails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rateResp, _ := rateGRPCClient.GetRate(rateProto.RateRequest{
		BaseCurrency:   "bitcoin",
		TargetCurrency: "uah",
	})
	rate := rateResp.Rate

	emailsResponse := storageGRPCClient.GetAllEmails(proto.GetAllEmailsRequest{})
	emails := emailsResponse.Email

	for i := range emails {
		emailGRPCClient.SendEmail(emailProto.SendEmailRequest{
			Body:    utils.BtcRateString + strconv.FormatFloat(rate, 'f', -1, 64),
			Subject: utils.BtcRateSubject,
			To:      emails[i],
		})
	}

	return
}
