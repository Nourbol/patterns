package ducks

import "fmt"

type FlyWithWings struct {
}

func NewFlyWithWings() *FlyWithWings {
	return &FlyWithWings{}
}

func (f *FlyWithWings) Fly() {
	fmt.Println("I am flying with wings")
}

type FlyNoWay struct {
}

func NewFlyNoWay() *FlyNoWay {
	return &FlyNoWay{}
}

func (f *FlyNoWay) Fly() {
	fmt.Println("I cannot fly")
}
