package main

type Transport struct {
	Wheels int
	Fuel   string
	Doors  int
}

type Vehicles struct {
	transport Transport
}

func (v *Vehicles) SetDoors(doors int) *Vehicles {
	v.transport.Doors = doors

	return v
}

func (v *Vehicles) SetFuel(fuel string) *Vehicles {
	v.transport.Fuel = fuel

	return v
}

func (v *Vehicles) SetWheels(wheels int) *Vehicles {
	v.transport.Wheels = wheels

	return v
}

// func main() {
// 	car := Vehicles{}
// 	car.SetDoors(4)
// 	car.SetFuel("Diesel")
// 	car.SetWheels(4)

// }
