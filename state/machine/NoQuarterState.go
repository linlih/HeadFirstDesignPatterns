package machine

import (
	"fmt"
)

type NoQuarterState struct {
	gumballMachine *GumballMachine
}

func NewNoQuarterState(gumballMachine *GumballMachine) *NoQuarterState {
	return &NoQuarterState{gumballMachine: gumballMachine}
}

func (n *NoQuarterState) InsertQuarter() {
	fmt.Println("You inserted a quarter")
	n.gumballMachine.SetState(n.gumballMachine.GetHasQuarterState()) // 插入硬币了！！
}

func (n *NoQuarterState) EjectQuarter() {
	fmt.Println("You haven't inserted a quarter")
}

func (n *NoQuarterState) TurnCrank() {
	fmt.Println("You turned, but there's no quarter")
}

func (n *NoQuarterState) Dispense() {
	fmt.Println("You need to pay first")
}

func (n *NoQuarterState) String() string {
	return "waiting for quarter"
}
