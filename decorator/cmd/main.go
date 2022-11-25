package main

import "patterns/decorator"

func main() {
	dc := decorator.NewNotifierSMSDecorator(decorator.NewNotifierSlackDecorator(decorator.Notifier{}))

	dc.Send("Hello!")
}
