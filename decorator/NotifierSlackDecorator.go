package decorator

import "fmt"

type NotifierSlackDecorator struct {
	wrapee NotifierInterface
}

func NewNotifierSlackDecorator(wrapee NotifierInterface) *NotifierSlackDecorator {
	return &NotifierSlackDecorator{wrapee}
}

func (ns NotifierSlackDecorator) Send(message string) {
	fmt.Printf("Send slack message %s\n", message)

	ns.wrapee.Send(message)
}
