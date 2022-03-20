package main

import (
	"adapter/adapter"
	"adapter/duck"
	"adapter/turkey"
	"fmt"
)

func main() {
	mallardDuck := new(duck.MallardDuck)
	wildTurkey := new(turkey.WildTurkey)
	duck := adapter.NewTurkeyAdapter(wildTurkey)
	turkey := adapter.NewDuckAdapter(mallardDuck)

	fmt.Println("The Turkey says...")
	testTurkey(wildTurkey)

	fmt.Println("The DuckAdapter says...")
	testTurkey(turkey)

	fmt.Println("The Duck says...")
	testDuck(mallardDuck)

	fmt.Println("The TurkeyAdapter says...")
	testDuck(duck)
}

func testDuck(duck duck.IDuck) {
	duck.Quack()
	duck.Fly()
}

func testTurkey(turkey turkey.ITurkey) {
	turkey.Gobble()
	turkey.Fly()
}
