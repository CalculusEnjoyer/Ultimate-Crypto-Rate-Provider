package executor

import (
	"email.com/email/dispatcher/executor/templates"
)

type Sender interface {
	Send(content templates.EmailContent, email string) (err error)
}
