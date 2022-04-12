package main

import (
	"fmt"
)

type Waitress struct {
	allMenus IMenuComponent
}

func NewWaitress(allMenus IMenuComponent) *Waitress {
	return &Waitress{allMenus: allMenus}
}

func (w *Waitress) PrintMenu() {
	w.allMenus.Print()
}

func (w *Waitress) PrintVegetarianMenu() {
	iterator := w.allMenus.CreateIterator()
	fmt.Println("\nVEGETARIAN MENU\n-----")
	for {
		if iterator.HasNext() {
			it := iterator.Next()
			item := it.(IMenuComponent)
			if item.IsVegetarian() {
				item.Print()
			}
		} else {
			break
		}
	}
}
