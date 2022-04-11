package main

import "strconv"

type MenuItem struct {
	Name        string
	Description string
	Vegetarian  bool
	Price       float32
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

func (m *MenuItem) GetPrice() float32 {
	return m.Price
}

func (m *MenuItem) IsVegetarian() bool {
	return m.Vegetarian
}

func (m *MenuItem) String() string {
	return m.Name + ", $" + strconv.FormatFloat(float64(m.Price), 'f', 2, 32) + "\n    " + m.Description
}
