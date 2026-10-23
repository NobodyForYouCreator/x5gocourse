package main

import (
	"errors"
	"fmt"
)

const (
	StatusNew  = "new"
	StatusPaid = "paid"
	StatusDone = "done"
)

type Item struct {
	Name  string
	Price int
}

type Order struct {
	ID     int
	Status string
	Items  []Item
}

func NewOrder(id int) (*Order, error) {
	if id <= 0 {
		return nil, errors.New("id должен быть положительным")
	}
	return &Order{ID: id, Status: StatusNew}, nil
}

func AddItem(order *Order, item Item) error {
	if order == nil {
		return errors.New("заказ nil")
	}
	if item.Name == "" {
		return errors.New("пустое имя позиции")
	}
	if item.Price < 0 {
		return errors.New("отрицательная цена")
	}
	order.Items = append(order.Items, item)
	return nil
}

func Total(order *Order) int {
	sum := 0
	for _, item := range order.Items {
		sum += item.Price
	}
	return sum
}

func main() {
	statuses := make(map[string]string)
	statuses[StatusNew] = "Новый"
	statuses[StatusPaid] = "Оплачен"
	statuses[StatusDone] = "Выполнен"

	order, err := NewOrder(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	AddItem(order, Item{Name: "Молоко", Price: 8900})
	AddItem(order, Item{Name: "Хлеб", Price: 5500})
	if err := AddItem(order, Item{Name: "", Price: 100}); err != nil {
		fmt.Println("ошибка:", err)
	}

	fmt.Println("Заказ", order.ID, statuses[order.Status], "сумма:", Total(order))
}
