package main

import "testing"

func TestNewOrderBadID(t *testing.T) {
	if _, err := NewOrder(0); err == nil {
		t.Error("ожидали ошибку для id = 0")
	}
}

func TestEmptyOrderTotal(t *testing.T) {
	order, _ := NewOrder(1)
	if Total(order) != 0 {
		t.Error("сумма пустого заказа должна быть 0")
	}
}

func TestTotalTwoItems(t *testing.T) {
	order, _ := NewOrder(1)
	AddItem(order, Item{Name: "Молоко", Price: 8900})
	AddItem(order, Item{Name: "Хлеб", Price: 5500})

	if Total(order) != 14400 {
		t.Errorf("ожидали 14400, получили %d", Total(order))
	}
}

func TestAddItemBadData(t *testing.T) {
	if err := AddItem(nil, Item{Name: "Хлеб", Price: 100}); err == nil {
		t.Error("ожидали ошибку для nil заказа")
	}

	order, _ := NewOrder(1)
	if err := AddItem(order, Item{Name: "", Price: 100}); err == nil {
		t.Error("ожидали ошибку для пустого имени")
	}
	if err := AddItem(order, Item{Name: "Сыр", Price: -5}); err == nil {
		t.Error("ожидали ошибку для отрицательной цены")
	}
	if len(order.Items) != 0 {
		t.Error("заказ не должен был измениться")
	}
}

func TestCopyOrder(t *testing.T) {
	order, _ := NewOrder(1)
	AddItem(order, Item{Name: "Хлеб", Price: 5500})

	copyOrder := *order
	copyOrder.Items = make([]Item, len(order.Items))
	copy(copyOrder.Items, order.Items)

	copyOrder.ID = 2
	copyOrder.Items[0].Price = 1

	if order.ID != 1 || order.Items[0].Price != 5500 {
		t.Error("исходный заказ изменился")
	}
}

func TestMapWithMake(t *testing.T) {
	statuses := make(map[string]string)
	statuses[StatusNew] = "Новый"

	if statuses[StatusNew] != "Новый" {
		t.Error("статус не записался")
	}
}
