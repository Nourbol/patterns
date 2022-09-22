package main

import (
	"fmt"
	"patterns/strategy/ducks"
)

func main() {
	fmt.Println("Mallard Duck:")
	duck := ducks.NewMallardDuck()
	duck.PerformFly()
	duck.PerformQuack()
	duck.Display()

	fmt.Println("Changing quack behaviour")
	duck.QuackBehaviour = ducks.NewSqueak()
	duck.PerformQuack()
}
