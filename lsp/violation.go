//go:build violation
// +build violation

package main

import "log"

type Bird struct {
	Type string
}

func (b *Bird) MoveWings() {
	log.Printf("%s is flapping its wings", b.Type)
}

func (b *Bird) Fly() {
	log.Printf("%s is flying", b.Type)
}

type Owl struct {
	Bird
}

type Penguin struct {
	Bird
}

func main() {
	owl := Owl{
		Bird{
			Type: "Owl",
		},
	}
	penguin := Penguin{
		Bird{
			Type: "Penguin",
		},
	}

	owl.MoveWings()     // ok - the owls move their wings
	penguin.MoveWings() // ok - the penguins move their wings

	owl.Fly()     // ok - the owls can fly
	penguin.Fly() // fail - the penguins can't fly
}
