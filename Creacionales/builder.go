package main

// Transport representa la estructura base de un medio de transporte con atributos como ruedas, tipo de combustible y número de puertas.
type Transport struct {
	Wheels int
	Fuel   string
	Doors  int
}

// Vehicles es una estructura que contiene un Transport como su composición.
// Provee métodos para configurar las propiedades del transporte.
type Vehicles struct {
	transport Transport
}

// SetDoors establece el número de puertas del vehículo y devuelve el mismo objeto Vehicles para permitir la llamada encadenada (method chaining).
func (v *Vehicles) SetDoors(doors int) *Vehicles {
	v.transport.Doors = doors
	return v
}

// SetFuel establece el tipo de combustible del vehículo y devuelve el mismo objeto Vehicles para permitir la llamada encadenada (method chaining).
func (v *Vehicles) SetFuel(fuel string) *Vehicles {
	v.transport.Fuel = fuel
	return v
}

// SetWheels establece el número de ruedas del vehículo y devuelve el mismo objeto Vehicles para permitir la llamada encadenada (method chaining).
func (v *Vehicles) SetWheels(wheels int) *Vehicles {
	v.transport.Wheels = wheels
	return v
}

//Funcion principal.
// func main() {
// 	car := Vehicles{}
// 	car.SetDoors(4)
// 	car.SetFuel("Diesel")
// 	car.SetWheels(4)
// }
