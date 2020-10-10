package main

/*
* Director
 */

type Dealer struct {
	VehicleBuilder VehicleBuilder
}

func (d *Dealer) construct(parameters interface{}) {
	d.VehicleBuilder.Build(parameters)
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
	Build(parameters interface{})
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

func (cb CarBuilder) Build(parameters interface{}) {

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

type AirPlane struct {
	color       string
	vehicleType string
	numOfWheel  int
	brand       string
}

func NewAirPlan() *AirPlane {
	return &AirPlane{
		brand: "boeing",
	}
}

func main() {

}
