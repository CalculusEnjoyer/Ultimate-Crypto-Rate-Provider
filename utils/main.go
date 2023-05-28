package main

import (
	"genesis-test-task/services/api/grpc/client/currency"
	"genesis-test-task/services/api/grpc/client/email"
	"genesis-test-task/services/api/grpc/client/storage"
	"genesis-test-task/services/currency/rate/messages/proto"
	emailproto "genesis-test-task/services/email/dispatcher/messages/proto"
	storageproto "genesis-test-task/services/storage/emails/messages/proto"
)

func main() {
	var check = currency.CurrencyGRPCClient{}

	res, _ := check.GetRate(proto.RateRequest{
		BaseCurrency:   "bitcoin",
		TargetCurrency: "uah",
	})

	print(res.Rate)

	var emailCheck = email.EmailGRPCClient{}

	resnew := emailCheck.SendEmail(emailproto.SendEmailRequest{
		Body:    "Check",
		Subject: "fjdljf",
		To:      "kravchukzxy@gmail.com",
	})

	print(resnew.Error)

	var storageCheck = storage.StorageGRPCClient{}

	resnew1 := storageCheck.AddEmail(storageproto.AddEmailRequest{
		Email: "CHECKk",
	})

	print(resnew1.Error)

	resnew2 := storageCheck.GetAllEmails(storageproto.GetAllEmailsRequest{})

	emails := resnew2.GetEmail()

	for i := range emails {
		print(emails[i])
	}
}
