package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrBadID         = errors.New("id должен быть положительным")
	ErrNilOrder      = errors.New("заказ nil")
	ErrEmptyName     = errors.New("пустое имя позиции")
	ErrNegativePrice = errors.New("отрицательная цена")
)

type ItemError struct {
	Name string
	Err  error
}

func (e *ItemError) Error() string {
	return fmt.Sprintf("позиция %q: %v", e.Name, e.Err)
}

func (e *ItemError) Unwrap() error {
	return e.Err
}

func (e *ItemError) Is(target error) bool {
	t, ok := target.(*ItemError)
	if !ok {
		return false
	}
	return e.Name == t.Name
}

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, "; ")
}

type Item struct {
	Name  string
	Price int
}

type Order struct {
	ID    int
	Items []Item
}

func NewOrder(id int) (*Order, error) {
	if id <= 0 {
		return nil, fmt.Errorf("создать заказ %d: %w", id, ErrBadID)
	}
	return &Order{ID: id}, nil
}

func checkItem(item Item) error {
	if item.Name == "" {
		return &ItemError{Name: item.Name, Err: ErrEmptyName}
	}
	if item.Price < 0 {
		return &ItemError{Name: item.Name, Err: ErrNegativePrice}
	}
	return nil
}

func AddItem(order *Order, item Item) error {
	if order == nil {
		return ErrNilOrder
	}
	if err := checkItem(item); err != nil {
		return fmt.Errorf("заказ %d: %w", order.ID, err)
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

func (o Order) ValidateBad() error {
	var m *MultiError
	for _, item := range o.Items {
		if err := checkItem(item); err != nil {
			if m == nil {
				m = &MultiError{}
			}
			m.Add(err)
		}
	}
	return m
}

func (o Order) Validate() error {
	var m *MultiError
	for _, item := range o.Items {
		if err := checkItem(item); err != nil {
			if m == nil {
				m = &MultiError{}
			}
			m.Add(err)
		}
	}
	if m != nil {
		return m
	}
	return nil
}
