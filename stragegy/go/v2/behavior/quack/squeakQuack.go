package quack

import "fmt"

type SqueakQuack struct {
}

func (s *SqueakQuack) Quack() {
	fmt.Println("Squeak quack")
}
