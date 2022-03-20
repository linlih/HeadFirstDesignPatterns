package duck

import "fmt"

type MallardDuck struct {
}

func (m *MallardDuck) Quack() {
	fmt.Println("Mallard duck Quack")
}

func (m *MallardDuck) Fly() {
	fmt.Println("Mallard duck: I'm flying")
}
