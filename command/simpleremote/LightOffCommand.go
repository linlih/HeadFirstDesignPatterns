package simpleremote

type LightOffCommand struct {
	light Light
}

func NewLightOffCommand(light Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (l *LightOffCommand) Execute() {
	l.light.Off()
}
