package main

import (
	"fmt"
	"strconv"
)

type Waitress struct {
	pancakeHouseMenu *PancakeHouseMenu
	dinerMenu        *DinerMenu
}

func NewWaitress(pancakeHouseMenu *PancakeHouseMenu, dinerMenu *DinerMenu) *Waitress {
	return &Waitress{pancakeHouseMenu: pancakeHouseMenu, dinerMenu: dinerMenu}
}

func (w *Waitress) PrintMenu() {
	pancakeIterator := w.pancakeHouseMenu.CreateIterator()
	dinerIterator := w.dinerMenu.CreateIterator()

	fmt.Println("MENU\n-----\nBREAKFAST")
	w.printMenu(pancakeIterator)
	fmt.Println("\nLUNCH")
	w.printMenu(dinerIterator)
}

func (w *Waitress) PrintVegetarianMenu() {
	pancakeIterator := w.pancakeHouseMenu.CreateIterator()
	dinerIterator := w.dinerMenu.CreateIterator()

	fmt.Println("Vegetarian MENU\n-----\nBREAKFAST")
	w.printVegetarianMenu(pancakeIterator)
	fmt.Println("\nDINER")
	w.printVegetarianMenu(dinerIterator)
}

func (w *Waitress) printVegetarianMenu(iterator IIterator) {
	for {
		item := iterator.Next()
		menuItem := item.(MenuItem)
		if menuItem.IsVegetarian() {
			fmt.Print(menuItem.GetName() + ",")
			fmt.Print(strconv.FormatFloat(float64(menuItem.GetPrice()), 'f', 2, 32) + " -- ")
			fmt.Println(menuItem.GetDescription())
		}
		if !iterator.HasNext() {
			break
		}
	}
}

func (w *Waitress) printMenu(iterator IIterator) {
	for {
		item := iterator.Next()
		menuItem := item.(MenuItem)
		fmt.Print(menuItem.GetName() + ",")
		fmt.Print(strconv.FormatFloat(float64(menuItem.GetPrice()), 'f', 2, 32) + " -- ")
		fmt.Println(menuItem.GetDescription())
		if !iterator.HasNext() {
			break
		}
	}
}
