package main

type PancakeHouseIterator struct {
	menuItem []MenuItem
	pos      int
}

func NewPancakeHouseIterator(menuItem []MenuItem) *PancakeHouseIterator {
	return &PancakeHouseIterator{menuItem: menuItem, pos: 0}
}

func (p *PancakeHouseIterator) Next() interface{} {
	obj := p.menuItem[p.pos]
	p.pos += 1
	return obj
}

func (p *PancakeHouseIterator) HasNext() bool {
	if p.pos >= len(p.menuItem) {
		return false
	}
	return true
}
