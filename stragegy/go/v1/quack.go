package main

import "fmt"

type quack struct {
}

func (q *quack) quack() {
	fmt.Println("Quack")
}
