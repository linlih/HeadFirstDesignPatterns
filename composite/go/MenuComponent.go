package main

type IMenuComponent interface {
	Add(component IMenuComponent)
	Remove(component IMenuComponent)
	GetChild() IMenuComponent
	GetName() string
	GetDescription() string
	GetPrice() float32
	IsVegetarian() bool
	CreateIterator() IIterator
	Print()
	GetIterator() IIterator
}
