package adapter

import "adapter/turkey"

type TurkeyAdapter struct {
	turkey turkey.ITurkey
}

func NewTurkeyAdapter(turkey turkey.ITurkey) *TurkeyAdapter {
	return &TurkeyAdapter{turkey: turkey}
}

func (t *TurkeyAdapter) Quack() {
	t.turkey.Gobble()
}

func (t *TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		t.turkey.Fly()
	}
}
