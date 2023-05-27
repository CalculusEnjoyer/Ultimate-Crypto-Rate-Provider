package dispatcher

import (
	sender2 "genesis-test-task/services/email/dispatcher/executor"
)

var dispatcher = sender2.GomailSender{}

type emailService struct{}

func NewService() EmailService { return &emailService{} }

func (e emailService) SendEmail(request sender2.SendEmailRequest) (err error) {
	return dispatcher.Send(request.Content, request.To)
}
