package main

type MenuIterator struct {
	menuComponents []IMenuComponent
	pos            int
}

func NewMenuIterator(menuComponents []IMenuComponent) *MenuIterator {
	return &MenuIterator{menuComponents: menuComponents, pos: 0}
}

func (m *MenuIterator) Next() interface{} {
	obj := (m.menuComponents)[m.pos]
	m.pos += 1
	return obj
}

func (m *MenuIterator) HasNext() bool {
	if m.pos >= len(m.menuComponents) {
		return false
	}
	return true
}
