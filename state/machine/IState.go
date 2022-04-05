package machine

type IState interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
	String() string
}
