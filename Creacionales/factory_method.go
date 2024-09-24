package main

import "fmt"

// Transporte es una interfaz que define el método Mover(), común para distintos tipos de transporte.
type Transporte interface {
	Mover()
}

// Auto representa la estructura de un automóvil con atributos específicos como puertas y ruedas.
type Auto struct {
	Puertas int
	Ruedas  int
}

// Barco representa la estructura de un barco con atributos como camarotes y el nombre del capitán.
type Barco struct {
	Camarotes int
	Capitan   string
}

// Mover implementa el método de la interfaz Transporte para la estructura Auto.
// Define el comportamiento de movimiento de un automóvil.
func (a Auto) Mover() {
	fmt.Println("El auto se mueve en ruta.")
}

// Mover implementa el método de la interfaz Transporte para la estructura Barco.
// Define el comportamiento de movimiento de un barco.
func (b Barco) Mover() {
	fmt.Println("El barco se mueve en el mar.")
}

// CrearTransporte es una función que crea y devuelve un objeto que implementa la interfaz Transporte.
// Dependiendo del tipo de transporte ("ruta" o "mar"), devuelve un Auto o un Barco.
func CrearTransporte(tipo string) Transporte {
	if tipo == "ruta" {
		return Auto{
			Puertas: 4,
			Ruedas:  4,
		}
	} else if tipo == "mar" {
		return Barco{
			Camarotes: 40,
			Capitan:   "Gustavo",
		}
	}
	return nil
}

//Funcion principal
// func main() {
// 	transporte := CrearTransporte("ruta")
// 	transporte.Mover()
// }
