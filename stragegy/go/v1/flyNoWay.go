package main

import "fmt"

type flyNoWay struct {
}

func (f *flyNoWay) fly() {
	fmt.Println("I can't fly")
}
