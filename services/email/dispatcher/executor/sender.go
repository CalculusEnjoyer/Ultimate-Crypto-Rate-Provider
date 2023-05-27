package executor

import (
	"genesis-test-task/services/email/dispatcher/executor/templates"
)

type Sender interface {
	Send(content templates.EmailContent, email string) (err error)
}
