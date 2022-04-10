package simpleremote

type LightOnCommand struct {
	light Light
}

func NewLightOnCommand(light Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (l *LightOnCommand) Execute() {
	l.light.On()
}
