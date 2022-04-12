package main

func main() {
	pancakeHouseMenu := NewMenu("PANCAKE HOUSE MENU", "Breakfast")
	dinerMenu := NewMenu("DINER MENU", "Lunch")
	cafeMenu := NewMenu("CAFE MENU", "Dinner")
	dessertMenu := NewMenu("DESSERT MENU", "Dessert of course!")

	allMenu := NewMenu("ALL MENUS", "ALL menus combined")
	allMenu.Add(pancakeHouseMenu)
	allMenu.Add(dinerMenu)
	allMenu.Add(cafeMenu)

	dinerMenu.Add(dessertMenu)

	pancakeHouseMenu.Add(NewMenuItem(
		"K&B's Pancake Breakfast",
		"Pancakes with scrambled eggs, and toast",
		true,
		2.99))
	pancakeHouseMenu.Add(NewMenuItem(
		"Regular Pancake Breakfast",
		"Pancakes with fried eggs, sausage",
		false,
		2.99))
	pancakeHouseMenu.Add(NewMenuItem(
		"Blueberry Pancakes",
		"Pancakes made with fresh blueberries, and blueberry syrup",
		true,
		3.49))
	pancakeHouseMenu.Add(NewMenuItem(
		"Waffles",
		"Waffles, with your choice of blueberries or strawberries",
		true,
		3.59))

	dinerMenu.Add(NewMenuItem(
		"Vegetarian BLT",
		"(Fakin') Bacon with lettuce & tomato on whole wheat",
		true,
		2.99))
	dinerMenu.Add(NewMenuItem(
		"BLT",
		"Bacon with lettuce & tomato on whole wheat",
		false,
		2.99))
	dinerMenu.Add(NewMenuItem(
		"Soup of the day",
		"A bowl of the soup of the day, with a side of potato salad",
		false,
		3.29))
	dinerMenu.Add(NewMenuItem(
		"Hotdog",
		"A hot dog, with saurkraut, relish, onions, topped with cheese",
		false,
		3.05))
	dinerMenu.Add(NewMenuItem(
		"Steamed Veggies and Brown Rice",
		"A medly of steamed vegetables over brown rice",
		true,
		3.99))

	dinerMenu.Add(NewMenuItem(
		"Pasta",
		"Spaghetti with Marinara Sauce, and a slice of sourdough bread",
		true,
		3.89))

	dessertMenu.Add(NewMenuItem(
		"Apple Pie",
		"Apple pie with a flakey crust, topped with vanilla icecream",
		true,
		1.59))
	dessertMenu.Add(NewMenuItem(
		"Cheesecake",
		"Creamy New York cheesecake, with a chocolate graham crust",
		true,
		1.99))
	dessertMenu.Add(NewMenuItem(
		"Sorbet",
		"A scoop of raspberry and a scoop of lime",
		true,
		1.89))

	cafeMenu.Add(NewMenuItem(
		"Veggie Burger and Air Fries",
		"Veggie burger on a whole wheat bun, lettuce, tomato, and fries",
		true,
		3.99))
	cafeMenu.Add(NewMenuItem(
		"Soup of the day",
		"A cup of the soup of the day, with a side salad",
		false,
		3.69))
	cafeMenu.Add(NewMenuItem(
		"Burrito",
		"A large burrito, with whole pinto beans, salsa, guacamole",
		true,
		4.29))

	waitress := NewWaitress(allMenu)
	waitress.PrintMenu()
	waitress.PrintVegetarianMenu()
}
