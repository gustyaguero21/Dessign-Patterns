package main

import "fmt"

type EuropeanSocket interface { // Enchufe europeo con dos clavijas
	PlugIntoEuropeanSocket()
}

type EuroPlug struct{} // Implementación de un enchufe europeo

func (e *EuroPlug) PlugIntoEuropeanSocket() {
	fmt.Println("Conectado a un enchufe europeo.")
}

type AmericanSocket interface { // Enchufe americano con tres clavijas
	PlugIntoAmericanSocket()
}

type USPlug struct{} // Implementación de un enchufe americano

func (u *USPlug) PlugIntoAmericanSocket() {
	fmt.Println("Conectado a un enchufe americano.")
}

type Adapter struct { // Adapter que permite conectar un enchufe americano a un enchufe europeo
	americanDevice AmericanSocket
}

func NewAdapter(device AmericanSocket) EuropeanSocket {
	return &Adapter{americanDevice: device}
}

func (a *Adapter) PlugIntoEuropeanSocket() { // Implementación del método del enchufe europeo adaptado
	fmt.Println("Adaptador en uso:")
	a.americanDevice.PlugIntoAmericanSocket()
}

func main() {

	americanPlug := &USPlug{} // Dispositivo americano con enchufe americano

	adapter := NewAdapter(americanPlug) // Adapter que permite conectar el enchufe americano a un enchufe europeo

	adapter.PlugIntoEuropeanSocket() // Conectamos el dispositivo americano a un enchufe europeo
}
