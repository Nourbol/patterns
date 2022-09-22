package ducks

import (
	"fmt"
)

type WoodenDuck struct {
	FlyBehaviour   FlyBehaviour
	QuackBehaviour QuackBehaviour
}

func NewWoodenDuck() *WoodenDuck {
	return &WoodenDuck{
		FlyBehaviour:   NewFlyNoWay(),
		QuackBehaviour: NewMutedQuack(),
	}
}

func (w *WoodenDuck) PerformFly() {
	w.FlyBehaviour.Fly()
}

func (w *WoodenDuck) PerformQuack() {
	w.QuackBehaviour.Quack()
}

func (w *WoodenDuck) Display() {
	fmt.Println("I an a wooden duck")
}
