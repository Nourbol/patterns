package ducks

import "fmt"

type Quack struct {
}

func NewQuack() *Quack {
	return &Quack{}
}

func (q *Quack) Quack() {
	fmt.Println("I am quacking")
}

type Squeak struct {
}

func NewSqueak() *Squeak {
	return &Squeak{}
}

func (s *Squeak) Quack() {
	fmt.Println("I am squeaking")
}

type MutedQuack struct {
}

func NewMutedQuack() *MutedQuack {
	return &MutedQuack{}
}

func (m *MutedQuack) Quack() {
	fmt.Println("...")
}
