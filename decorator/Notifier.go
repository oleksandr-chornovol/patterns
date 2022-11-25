package decorator

import "fmt"

type Notifier struct {
}

func (n Notifier) Send(message string) {
	fmt.Printf("Send default message %s\n", message)
}
