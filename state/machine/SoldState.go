package machine

import (
	"fmt"
)

type SoldState struct {
	gumballMachine *GumballMachine
}

func NewSoldState(gumballMachine *GumballMachine) *SoldState {
	return &SoldState{gumballMachine: gumballMachine}
}

func (s *SoldState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

func (s *SoldState) EjectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}

func (s *SoldState) TurnCrank() {
	fmt.Println("You turned, but there are no gumballs")
}

func (s *SoldState) Dispense() {
	fmt.Println("No gumball dispensed")
}

func (s *SoldState) String() string {
	return "sold out"
}
