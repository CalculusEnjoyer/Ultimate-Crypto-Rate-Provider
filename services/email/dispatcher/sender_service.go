package dispatcher

import (
	"email.com/email/dispatcher/executor"
)

type EmailService interface {
	SendEmail(req executor.SendEmailRequest) (err error)
}
