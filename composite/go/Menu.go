package main

import (
	"fmt"
)

type Menu struct {
	MenuItem
	menuComponents []IMenuComponent
}

func NewMenu(name, des string) *Menu {
	m := &Menu{}
	m.Name = name
	m.Description = des
	return m
}

func (m *Menu) Add(menuComponent IMenuComponent) {
	m.menuComponents = append(m.menuComponents, menuComponent)
}

func (m *Menu) Remove(menuComponent IMenuComponent) {
	idx := 0
	for {
		if m.menuComponents[idx] == menuComponent {
			break
		}
		idx += 1
	}
	if idx != len(m.menuComponents) {
		m.menuComponents = append(m.menuComponents[:idx], m.menuComponents[idx+1:]...)
	} else {
		m.menuComponents = m.menuComponents[:idx]
	}
}

// IsVegetarian 菜单不属于素菜
func (m *Menu) IsVegetarian() bool {
	return false
}

func (m *Menu) CreateIterator() IIterator {
	return NewCompositeIterator(NewMenuIterator(m.menuComponents))
}

func (m *Menu) GetIterator() IIterator {
	return NewMenuIterator(m.menuComponents)
}

func (m *Menu) Print() {
	fmt.Print("" + m.GetName())
	fmt.Println(", " + m.GetDescription())
	fmt.Println("---------------------")
	iterator := m.GetIterator()
	for {
		if !iterator.HasNext() {
			break
		}
		item := iterator.Next()
		menuItem := item.(IMenuComponent)
		menuItem.Print()
	}
}

func (m *Menu) PrintVegetarian() {
	fmt.Print("" + m.GetName())
	fmt.Println(", " + m.GetDescription())
	fmt.Println("---------------------")
	iterator := m.CreateIterator()
	for {
		if !iterator.HasNext() {
			break
		}
		item := iterator.Next()
		menuItem := item.(IMenuComponent)
		menuItem.Print()
	}
}
