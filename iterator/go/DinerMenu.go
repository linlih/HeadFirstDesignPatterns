package main

import "container/list"

type DinerMenu struct {
	menuItem list.List
}

func NewDinerMenu() *DinerMenu {
	d := &DinerMenu{}
	d.menuItem.Init()

	d.AddItem("Vegetarian BLT",
		"(Fakin') Bacon with lettuce & tomato on whole wheat",
		true,
		2.99)

	d.AddItem("BLT",
		"Bacon with lettuce & tomato on whole wheat",
		false,
		2.99)

	d.AddItem("Soup of the day",
		"Soup of the day, with a side of potato salad",
		false,
		3.29)

	d.AddItem("Hotdog",
		"A hot dog, with saurkraut, relish, onions, topped with cheese",
		false,
		3.05)

	d.AddItem("Steamed Veggies and Brown Rice",
		"Steamed vegetables over brown rice",
		true,
		3.99)

	d.AddItem("Pasta",
		"Spaghetti with Marinara Sauce, and a slice of sourdough bread",
		true,
		3.89)

	return d
}

func (d *DinerMenu) AddItem(name, description string, isVegetarian bool, price float32) {
	d.menuItem.PushFront(MenuItem{
		Name:        name,
		Description: description,
		Vegetarian:  isVegetarian,
		Price:       price,
	})
}

func (d *DinerMenu) GetMenuItems() list.List {
	return d.menuItem
}

func (d *DinerMenu) CreateIterator() IIterator {
	return NewDinerMenuIterator(d.menuItem)
}
