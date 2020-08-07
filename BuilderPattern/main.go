package main

import (
	"fmt"
)

/*
* Director
 */

type Dealer struct {
	VehicleBuilder VehicleBuilder
}

func (d *Dealer) construct(color, vehicleType string, numOfWheel int) {
	d.VehicleBuilder.setColor(color)
	d.VehicleBuilder.setType(vehicleType)
	d.VehicleBuilder.setNumOfWheel(numOfWheel)
}

func NewDealer(vb VehicleBuilder) *Dealer {
	d := &Dealer{
		VehicleBuilder: vb,
	}
	return d
}

/*
* Builder
 */

type VehicleBuilder interface {
	setColor(color string)
	setType(vehicleType string)
	setNumOfWheel(numOfWheel int)
}

/*
* ConcreateBuilder
 */

type Car struct {
	color       string
	vehicleType string
	numOfWheel  int
}

func NewCar() *Car {
	return &Car{}
}

type CarBuilder struct {
	*Car
}

func (cb CarBuilder) setColor(color string) {
	cb.Car.color = color
}

func (cb CarBuilder) setType(vehicleType string) {
	cb.Car.vehicleType = vehicleType
}

func (cb CarBuilder) setNumOfWheel(numOfWheel int) {
	cb.Car.numOfWheel = numOfWheel
}

func (cb CarBuilder) getResult() *Car {
	return cb.Car
}

func NewCarBuilder() CarBuilder {
	nc := NewCar()
	cb := CarBuilder{
		Car: nc,
	}
	return cb
}

func main() {
	carBuilder := NewCarBuilder()
	dealer := NewDealer(carBuilder)
	dealer.construct("silver", "sportsCar", 4)
	car := carBuilder.getResult()
	fmt.Print(*car)
}
