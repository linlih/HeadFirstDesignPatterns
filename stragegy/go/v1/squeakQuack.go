package main

import "fmt"

type squeakQuack struct {
}

func (s *squeakQuack) quack() {
	fmt.Println("Squeak quack")
}
