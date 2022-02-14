package main

import "fmt"

type fakeQuack struct {
}

func (f *fakeQuack) quack() {
	fmt.Println("fake quack")
}
