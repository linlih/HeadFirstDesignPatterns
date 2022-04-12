package main

import (
	"fmt"
	"strconv"
)

type MenuItem struct {
	Name        string
	Description string
	Vegetarian  bool
	Price       float32
}

func (m *MenuItem) Print() {
	fmt.Print("  " + m.GetName())
	if m.IsVegetarian() {
		fmt.Print("(v)")
	}
	fmt.Println(", " + strconv.FormatFloat(float64(m.GetPrice()), 'f', 2, 32))
	fmt.Println("     -- " + m.GetDescription())
}

func NewMenuItem(name string, description string, vegetarian bool, price float32) *MenuItem {
	return &MenuItem{Name: name, Description: description, Vegetarian: vegetarian, Price: price}
}

func (m *MenuItem) GetName() string {
	return m.Name
}

func (m *MenuItem) GetDescription() string {
	return m.Description
}

func (m *MenuItem) IsVegetarian() bool {
	return m.Vegetarian
}

func (m *MenuItem) CreateIterator() IIterator {
	return NewNullIterator()
}

func (m *MenuItem) GetIterator() IIterator {
	return NewNullIterator()
}

func (m *MenuItem) GetPrice() float32 {
	return m.Price
}

func (m *MenuItem) Add(component IMenuComponent) {
	//TODO implement me
	panic("implement me")
}

func (m *MenuItem) Remove(component IMenuComponent) {
	//TODO implement me
	panic("implement me")
}

func (m *MenuItem) GetChild() IMenuComponent {
	//TODO implement me
	panic("implement me")
}
