package main

type Vehicle interface { // Definir la interfaz con parámetros adecuados
	ChooseColor(color string) string
	SelectComfort(comfort string) string
}

type Supersport struct{} // Supersport implementa la interfaz Vehicle

func (s *Supersport) ChooseColor(color string) string {
	return "El color elegido es " + color
}

func (s *Supersport) SelectComfort(comfort string) string {
	return "La gama elegida es la " + comfort
}

type RacingCar struct { // RacingCar también implementa la interfaz Vehicle y delega en Supersport
	supersport Vehicle
}

func (r *RacingCar) ChooseColor(color string) string {
	return r.supersport.ChooseColor(color)
}

func (r *RacingCar) SelectComfort(comfort string) string {
	return r.supersport.SelectComfort(comfort)
}

// func main() {
// 	// Crear una instancia de Supersport
// 	audi := &Supersport{}
// 	fmt.Println(audi.ChooseColor("azul"))
// 	fmt.Println(audi.SelectComfort("full"))

// 	// Crear una instancia de RacingCar, delegando en un objeto de tipo Vehicle (Supersport)
// 	audi_tt := RacingCar{supersport: audi}
// 	fmt.Println(audi_tt.ChooseColor("rojo"))
// 	fmt.Println(audi_tt.SelectComfort("Full"))
// }
