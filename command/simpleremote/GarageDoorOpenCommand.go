package simpleremote

type GarageDoorOpenCommand struct {
	garageDoor GarageDoor
}

func NewGarageDoorOpenCommand(garageDoor GarageDoor) *GarageDoorOpenCommand {
	return &GarageDoorOpenCommand{garageDoor: garageDoor}
}

func (g *GarageDoorOpenCommand) Execute() {
	g.garageDoor.Up()
}
