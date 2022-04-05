package machine

import (
	"fmt"
)

type WinnerState struct {
	gumballMachine *GumballMachine
}

func NewWinnerState(gumballMachine *GumballMachine) *WinnerState {
	return &WinnerState{gumballMachine: gumballMachine}
}

func (w *WinnerState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you Gumball")
}

func (w *WinnerState) EjectQuarter() {
	fmt.Println("Please wait, we're already giving you Gumball")
}

func (w *WinnerState) TurnCrank() {
	fmt.Println("Turning again doesn't get you another gumball")
}

func (w *WinnerState) Dispense() {
	fmt.Println("YOU'RE A WINNER! you get two gumballs for you quarter")
	w.gumballMachine.ReleaseBall()
	if w.gumballMachine.GetCount() == 0 {
		w.gumballMachine.SetState(w.gumballMachine.GetSoldOutState())
	} else {
		w.gumballMachine.ReleaseBall()
		if w.gumballMachine.GetCount() > 0 {
			w.gumballMachine.SetState(w.gumballMachine.GetNoQuarterState())
		} else {
			fmt.Println("Oops, out of gumballs!")
			w.gumballMachine.SetState(w.gumballMachine.GetSoldOutState())
		}
	}
}

func (w *WinnerState) String() string {
	return "dispensing two gumballs for you quarter, because YOU'RE A WINNER!"
}
