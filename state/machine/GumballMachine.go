package machine

import (
	"fmt"
	"strconv"
	"strings"
)

type GumballMachine struct {
	state           IState
	soldOutState    IState
	noQuarterState  IState
	hasQuarterState IState
	soldState       IState
	winnerState     IState
	count           int // 记录当前还有多少糖果
}

func NewGumballMachine(count int) *GumballMachine {
	var gumballMachine GumballMachine
	gumballMachine.hasQuarterState = NewHasQuarterState(&gumballMachine)
	gumballMachine.soldState = NewSoldState(&gumballMachine)
	gumballMachine.noQuarterState = NewNoQuarterState(&gumballMachine)
	gumballMachine.winnerState = NewWinnerState(&gumballMachine)
	gumballMachine.soldOutState = NewSoldOutState(&gumballMachine)

	// 默认设置为没有硬币的状态
	gumballMachine.SetState(gumballMachine.GetNoQuarterState())
	gumballMachine.count = count
	return &gumballMachine
}

func (g *GumballMachine) InsertQuarter() {
	g.state.InsertQuarter()
}

func (g *GumballMachine) EjectQuarter() {
	g.state.EjectQuarter()
}

func (g *GumballMachine) TurnCrank() {
	g.state.TurnCrank()
	g.state.Dispense()
}

func (g *GumballMachine) Dispense() {
	g.state.Dispense()
}

func (g *GumballMachine) SetState(state IState) {
	g.state = state
}

func (g *GumballMachine) GetSoldOutState() IState {
	return g.soldOutState
}

func (g *GumballMachine) GetNoQuarterState() IState {
	return g.noQuarterState
}

func (g *GumballMachine) GetHasQuarterState() IState {
	return g.hasQuarterState
}

func (g *GumballMachine) GetWinnerState() IState {
	return g.winnerState
}

func (g *GumballMachine) GetCount() int {
	return g.count
}

func (g *GumballMachine) ReleaseBall() {
	fmt.Println("A gumball comes rolling out the slot ...")
	if g.count != 0 {
		g.count -= 1
	}
}

func (g *GumballMachine) Refill(count int) {
	g.count = count
	g.state = g.noQuarterState
}

func (g *GumballMachine) String() string {
	var str strings.Builder
	str.WriteString("\nMighty Gumball, Inc.")
	str.WriteString("\nGo-enabled Standing Gumball Model #2022")
	str.WriteString("\nInventory:" + strconv.FormatInt(int64(g.count), 10) + " gumball")
	if g.count != 1 {
		str.WriteString("s")
	}
	str.WriteString("\n")
	str.WriteString("Machine is " + g.state.String() + "\n")
	return str.String()
}
