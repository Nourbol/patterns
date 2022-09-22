package ducks

import (
	"fmt"
)

type MallardDuck struct {
	FlyBehaviour   FlyBehaviour
	QuackBehaviour QuackBehaviour
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{
		FlyBehaviour:   NewFlyWithWings(),
		QuackBehaviour: NewQuack(),
	}
}

func (m *MallardDuck) PerformQuack() {
	m.QuackBehaviour.Quack()
}

func (m *MallardDuck) PerformFly() {
	m.FlyBehaviour.Fly()
}

func (m *MallardDuck) Display() {
	fmt.Println("I am a mallard duck")
}
