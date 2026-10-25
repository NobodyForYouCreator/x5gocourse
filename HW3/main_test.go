package main

import (
	"errors"
	"testing"
)

func TestNewOrderBadID(t *testing.T) {
	_, err := NewOrder(0)
	if !errors.Is(err, ErrBadID) {
		t.Errorf("ожидали ErrBadID, получили %v", err)
	}
}

func TestTotal(t *testing.T) {
	order, _ := NewOrder(1)
	if Total(order) != 0 {
		t.Error("сумма пустого заказа должна быть 0")
	}

	AddItem(order, Item{Name: "Молоко", Price: 8900})
	AddItem(order, Item{Name: "Хлеб", Price: 5500})
	if Total(order) != 14400 {
		t.Errorf("ожидали 14400, получили %d", Total(order))
	}
}

func TestAddItemBadData(t *testing.T) {
	if err := AddItem(nil, Item{Name: "Хлеб", Price: 100}); !errors.Is(err, ErrNilOrder) {
		t.Errorf("ожидали ErrNilOrder, получили %v", err)
	}

	order, _ := NewOrder(1)
	if err := AddItem(order, Item{Name: "", Price: 100}); !errors.Is(err, ErrEmptyName) {
		t.Errorf("ожидали ErrEmptyName, получили %v", err)
	}
	if err := AddItem(order, Item{Name: "Сыр", Price: -5}); !errors.Is(err, ErrNegativePrice) {
		t.Errorf("ожидали ErrNegativePrice, получили %v", err)
	}
	if len(order.Items) != 0 {
		t.Error("заказ не должен был измениться")
	}
}

func TestItemErrorIs(t *testing.T) {
	order, _ := NewOrder(1)
	err := AddItem(order, Item{Name: "Сыр", Price: -5})

	if !errors.Is(err, &ItemError{Name: "Сыр"}) {
		t.Error("ожидали ошибку про Сыр")
	}
	if errors.Is(err, &ItemError{Name: "Хлеб"}) {
		t.Error("это не ошибка про Хлеб")
	}
	if itemErr, ok := errors.AsType[*ItemError](err); !ok || itemErr.Name != "Сыр" {
		t.Errorf("AsType не сработал: %v", err)
	}
}

func TestNilReceiver(t *testing.T) {
	order := Order{ID: 1, Items: []Item{{Name: "Хлеб", Price: 5500}}}

	if order.ValidateBad() == nil {
		t.Error("ожидали, что ValidateBad вернёт не-nil интерфейс")
	}
	if order.Validate() != nil {
		t.Error("Validate должен вернуть nil для валидного заказа")
	}

	order.Items = append(order.Items, Item{Name: "", Price: 1}, Item{Name: "Сыр", Price: -1})
	if order.Validate() == nil {
		t.Error("Validate должен вернуть ошибку")
	}
}

func TestSafeSend(t *testing.T) {
	if err := safeSend(EmailNotifier{}, "привет"); err != nil {
		t.Errorf("ожидали nil, получили %v", err)
	}
	if err := safeSend(EmailNotifier{}, ""); !errors.Is(err, ErrEmptyMessage) {
		t.Errorf("ожидали ErrEmptyMessage, получили %v", err)
	}
	if err := safeSend(BrokenNotifier{}, "привет"); err == nil {
		t.Error("ожидали ошибку вместо паники")
	}
}

func TestNotifyAllPanicGoesUp(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("ожидали, что паника дойдёт до вызывающего")
		}
	}()

	order, _ := NewOrder(1)
	notifyAll([]Notifier{BrokenNotifier{}}, order)
}

func TestNotifyAllJoin(t *testing.T) {
	order, _ := NewOrder(1)
	err := notifyAll([]Notifier{EmailNotifier{}, EmailNotifier{}}, order)
	if err != nil {
		t.Errorf("ожидали nil, получили %v", err)
	}
}
