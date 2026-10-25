package main

import (
	"errors"
	"fmt"
)

var ErrEmptyMessage = errors.New("пустое сообщение")

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct{}

type BrokenNotifier struct{}

func (EmailNotifier) Send(message string) error {
	if message == "" {
		return ErrEmptyMessage
	}
	fmt.Println("email отправлен:", message)
	return nil
}

func (BrokenNotifier) Send(string) error {
	panic("драйвер канала повреждён")
}

func safeSend(notifier Notifier, message string) (err error) {
	defer func() {
		if value := recover(); value != nil {
			err = fmt.Errorf("паника при отправке: %v", value)
		}
	}()

	return notifier.Send(message)
}

func notifyAll(notifiers []Notifier, order *Order) error {
	var errs []error
	for _, n := range notifiers {
		err := n.Send(fmt.Sprintf("заказ %d готов, сумма %d", order.ID, Total(order)))
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}
