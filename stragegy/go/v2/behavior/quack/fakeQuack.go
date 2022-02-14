package quack

import "fmt"

type FakeQuack struct {
}

func (f *FakeQuack) Quack() {
	fmt.Println("fake quack")
}
