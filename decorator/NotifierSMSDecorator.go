package decorator

import "fmt"

type NotifierSMSDecorator struct {
	wrapee NotifierInterface
}

func NewNotifierSMSDecorator(wrapee NotifierInterface) *NotifierSMSDecorator {
	return &NotifierSMSDecorator{wrapee}
}

func (nsms NotifierSMSDecorator) Send(message string) {
	fmt.Printf("Send sms message %s\n", message)

	nsms.wrapee.Send(message)
}
