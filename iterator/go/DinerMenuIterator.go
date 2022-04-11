package main

import "container/list"

type DinerMenuIterator struct {
	menuItem list.List
	item     *list.Element
}

func NewDinerMenuIterator(menuItem list.List) *DinerMenuIterator {
	return &DinerMenuIterator{menuItem: menuItem}
}

func (d *DinerMenuIterator) Next() interface{} {
	if d.item == nil {
		d.item = d.menuItem.Front()
		return d.item.Value
	}
	d.item = d.item.Next()
	return d.item.Value
}

func (d *DinerMenuIterator) HasNext() bool {
	if d.item.Next() != nil {
		return true
	}
	return false
}
