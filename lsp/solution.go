package main

import "log"

type Bird struct {
	Type string
}

type BirdThatFlies struct {
	Bird
}

func (b *Bird) MoveWings() {
	log.Printf("%s is flapping its wings", b.Type)
}

func (b *BirdThatFlies) Fly() {
	log.Printf("%s is flying", b.Type)
}

type Owl struct {
	BirdThatFlies
}

type Penguin struct {
	Bird
}

func main() {
	owl := Owl{
		BirdThatFlies{
			Bird{
				Type: "Owl",
			},
		},
	}
	penguin := Penguin{
		Bird{
			Type: "Penguin",
		},
	}

	owl.MoveWings()     // ok - the owls move their wings
	penguin.MoveWings() // ok - the penguins move their wings

	owl.Fly() // ok - the owls can fly
}
