package turkey

import "fmt"

type WildTurkey struct {
}

func (w *WildTurkey) Gobble() {
	fmt.Println("Wild Turkey gobble")
}

func (w *WildTurkey) Fly() {
	fmt.Println("Wild turkey: I'm flying a short distance")
}
