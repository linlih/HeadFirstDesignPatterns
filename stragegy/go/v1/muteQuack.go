package main

import "fmt"

type muteQuack struct {
}

func (m *muteQuack) quack() {
	fmt.Println("<< Silence >>")
}
