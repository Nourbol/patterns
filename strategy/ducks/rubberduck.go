package ducks

import (
	"fmt"
)

type RubberDuck struct {
	FlyBehaviour   FlyBehaviour
	QuackBehaviour QuackBehaviour
}

func NewRubberDuck() *RubberDuck {
	return &RubberDuck{
		FlyBehaviour:   NewFlyNoWay(),
		QuackBehaviour: NewSqueak(),
	}
}

func (r *RubberDuck) PerformFly() {
	r.FlyBehaviour.Fly()
}

func (r *RubberDuck) PerformQuack() {
	r.QuackBehaviour.Quack()
}

func (r *RubberDuck) Display() {
	fmt.Println("I am a rubber duck")
}
