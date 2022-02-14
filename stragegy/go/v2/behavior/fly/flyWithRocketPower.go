package fly

import "fmt"

type FlyWithRocketPower struct {
}

func (f *FlyWithRocketPower) Fly() {
	fmt.Println("I'm flying with a rocket")
}
