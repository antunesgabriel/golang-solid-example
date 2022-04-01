package main

import "errors"

type Transport interface {
	Move()
}

type Machine interface {
	Transport
	TurnOn()
}

type Car struct{}

func (c *Car) Move() {
	// get around with the car
}
func (c *Car) TurnOn() {
	// we need to start the car
}

type Bike struct{}

func (b *Bike) Move() {
	// get around with the bike
}

func (b *Bike) TurnOn() {
	errors.New("we wouldn't need to implement this because the bike doesn't turn on")
}

func MoveTransport(transport Transport) {
	transport.Move()
}

func MoveMachine(machine Machine) {
	machine.TurnOn()

	MoveTransport(machine)
}

func main() {
	car := Car{}
	bike := Bike{}

	MoveMachine(&car)
	MoveTransport(&bike)
}
