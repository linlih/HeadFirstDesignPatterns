package simpleremote

type Controller struct {
	slot ICommand
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) SetCommand(command ICommand) {
	c.slot = command
}

func (c *Controller) ButtonWasPressed() {
	c.slot.Execute()
}
