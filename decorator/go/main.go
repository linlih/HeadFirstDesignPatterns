package main

import (
	"fmt"
	"strconv"
)

func main() {
	beverage := NewEspresso()
	fmt.Println(beverage.GetDescription() + " $" + strconv.FormatFloat(float64(beverage.Cost()), 'f', 2, 32))

	beverage2 := NewDarkRoast()
	beverage2 = NewMocha(beverage2)
	beverage2 = NewMocha(beverage2)
	beverage2 = NewWhip(beverage2)
	fmt.Println(beverage2.GetDescription() + " $" + strconv.FormatFloat(float64(beverage2.Cost()), 'f', 2, 32))

	beverage3 := NewDecaf()
	beverage3 = NewSoy(beverage3)
	beverage3 = NewMocha(beverage3)
	beverage3 = NewMilk(beverage3)
	beverage3 = NewWhip(beverage3)
	fmt.Println(beverage3.GetDescription() + " $" + strconv.FormatFloat(float64(beverage3.Cost()), 'f', 2, 32))
}
