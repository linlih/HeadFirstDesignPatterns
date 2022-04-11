package main

type PancakeHouseMenu struct {
	menuItem []MenuItem
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	p := &PancakeHouseMenu{}
	p.AddItem("K&B's Pancake Breakfast",
		"Pancakes with scrambled eggs, and toast",
		true,
		2.99)

	p.AddItem("Regular Pancake Breakfast",
		"Pancakes with fried eggs, and sausage",
		false,
		2.99)

	p.AddItem("Blueberry Pancake",
		"Pancakes with fresh blueberries",
		true,
		3.49)

	p.AddItem("Waffles",
		"Waffles, with your choice if blueberries or strawberries",
		true,
		3.59)
	return p
}

func (p *PancakeHouseMenu) AddItem(name, description string, isVegetarian bool, price float32) {
	p.menuItem = append(p.menuItem, MenuItem{
		Name:        name,
		Description: description,
		Vegetarian:  isVegetarian,
		Price:       price})
}

func (p *PancakeHouseMenu) GetMenuItems() []MenuItem {
	return p.menuItem
}

func (p *PancakeHouseMenu) CreateIterator() IIterator {
	return NewPancakeHouseIterator(p.menuItem)
}

func (p *PancakeHouseMenu) String() string {
	return "Object ville Pancake House Menu"
}
