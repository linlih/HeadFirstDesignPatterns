package machine

import (
	"fmt"
	"math/rand"
	"time"
)

type HasQuarterState struct {
	gumballMachine *GumballMachine
}

func NewHasQuarterState(gumballMachine *GumballMachine) *HasQuarterState {
	return &HasQuarterState{gumballMachine: gumballMachine}
}

func (s *HasQuarterState) InsertQuarter() {
	fmt.Println("You can't insert another quarter")
}

func (s *HasQuarterState) EjectQuarter() {
	fmt.Println("Quarter returned")
	s.gumballMachine.SetState(s.gumballMachine.GetNoQuarterState()) // 弹出硬币之后应该转变成没有硬币的状态
}

func (s *HasQuarterState) TurnCrank() {
	fmt.Println("You turned...")
	rand.Seed(time.Now().Unix())
	winner := rand.Intn(2) // 游戏：书中10%的概率太低了，改成50%
	fmt.Println(winner)
	if winner == 0 && s.gumballMachine.GetCount() > 1 {
		s.gumballMachine.SetState(s.gumballMachine.GetWinnerState())
	} else {
		s.gumballMachine.SetState(s.gumballMachine.GetSoldOutState())
	}
}

func (s *HasQuarterState) Dispense() {
	fmt.Println("No gumball dispensed")
}

func (s *HasQuarterState) String() string {
	return "waiting for turn of crank"
}
