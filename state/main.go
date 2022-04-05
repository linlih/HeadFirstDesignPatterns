package main

import (
	"fmt"
	"state/machine"
)

func main() {
	gumballMachine := machine.NewGumballMachine(10)

	fmt.Println(gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
}
