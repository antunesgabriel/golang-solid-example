//go:build violation
// +build violation

package main

import "errors"

type Transport interface {
	Move()
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

func main() {
	car := Car{}
	bike := Bike{}

	car.TurnOn()
	bike.TurnOn() // error
}
