package dispatcher

import (
	"genesis-test-task/services/email/dispatcher/executor"
)

type EmailService interface {
	SendEmail(req executor.SendEmailRequest) (err error)
}
