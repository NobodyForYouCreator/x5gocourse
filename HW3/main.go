package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if value := recover(); value != nil {
			fmt.Println("main перехватил панику:", value)
		}
	}()

	_, err := NewOrder(0)
	fmt.Println(err, "| ErrBadID:", errors.Is(err, ErrBadID))

	order, _ := NewOrder(1)
	AddItem(order, Item{Name: "Молоко", Price: 8900})
	AddItem(order, Item{Name: "Хлеб", Price: 5500})
	fmt.Println("сумма:", Total(order))

	err = AddItem(order, Item{Name: "Сыр", Price: -1})
	fmt.Println(err)
	fmt.Println("ErrNegativePrice:", errors.Is(err, ErrNegativePrice))
	fmt.Println("ошибка про Сыр:", errors.Is(err, &ItemError{Name: "Сыр"}))
	if itemErr, ok := errors.AsType[*ItemError](err); ok {
		fmt.Println("позиция из ошибки:", itemErr.Name)
	}

	err = order.ValidateBad()
	fmt.Println("ValidateBad: err == nil?", err == nil)
	err = order.Validate()
	fmt.Println("Validate: err == nil?", err == nil)

	err = safeSend(BrokenNotifier{}, "Заказ готов")
	fmt.Println("safeSend:", err)

	notifyAll([]Notifier{EmailNotifier{}, BrokenNotifier{}}, order)
	fmt.Println("сюда не дойдём")
}
