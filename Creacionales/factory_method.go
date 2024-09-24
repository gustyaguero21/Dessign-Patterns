package main

import "fmt"

type Transporte interface { // Es la interfaz que contendra el metodo que actuara como fabrica de objetos.
	Mover()
}

type Auto struct { // Estructura de un Automovil con algunos campos asociados.
	Puertas int
	Ruedas  int
}

type Barco struct { // Estructura de un Barco con algunos campos asociados.
	Camarotes int
	Capitan   string
}

func (a Auto) Mover() { // Esta funcion posee la firma de la estructura Auto. Cualquier Transporte que creemos, accedera a las funciones de Auto.
	fmt.Println("El auto se mueve en ruta.")
}

func (b Barco) Mover() { // Esta funcion posee la firma de la estructura Barco. Cualquier Transporte que creemos, accedera a las funciones de Barco.
	fmt.Println("El barco se mueve en el mar.")
}

func CrearTransporte(tipo string) Transporte { // Funcion que se encarga de crear un nuevo objeto y que dependiendo de las caracteristicas del mismo, decidira el tipo de objeto que sera.
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

// func main() {
// 	transporte := CrearTransporte("ruta")
// 	transporte.Mover()
// }
