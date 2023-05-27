package executor

import (
	"genesis-test-task/services/email/dispatcher/executor/templates"
)

type SendEmailRequest struct {
	To      string
	Content templates.EmailContent
}
