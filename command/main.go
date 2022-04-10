package main

import "command/simpleremote"

func main() {
	controller := simpleremote.NewController()
	light := simpleremote.NewLight()
	garageDoor := simpleremote.NewGarageDoor()
	lightOnCommand := simpleremote.NewLightOnCommand(*light)
	garageOpen := simpleremote.NewGarageDoorOpenCommand(*garageDoor)

	controller.SetCommand(lightOnCommand)
	controller.ButtonWasPressed()
	controller.SetCommand(garageOpen)
	controller.ButtonWasPressed()
}
