package main

// Vehicle es una interfaz que define métodos para elegir el color y la gama de confort de un vehículo.
type Vehicle interface {
	ChooseColor(color string) string
	SelectComfort(comfort string) string
}

// Supersport es una estructura que implementa la interfaz Vehicle,
// proporcionando su propia versión de los métodos ChooseColor y SelectComfort.
type Supersport struct{}

// ChooseColor implementa el método de la interfaz Vehicle para Supersport.
// Retorna un mensaje con el color seleccionado.
func (s *Supersport) ChooseColor(color string) string {
	return "El color elegido es " + color
}

// SelectComfort implementa el método de la interfaz Vehicle para Supersport.
// Retorna un mensaje con la gama de confort seleccionada.
func (s *Supersport) SelectComfort(comfort string) string {
	return "La gama elegida es la " + comfort
}

// RacingCar es una estructura que también implementa la interfaz Vehicle.
// Delegará las llamadas de los métodos en un objeto de tipo Vehicle, como Supersport.
type RacingCar struct {
	supersport Vehicle
}

// ChooseColor delega la implementación de la interfaz Vehicle al objeto supersport (que debe ser de tipo Vehicle).
func (r *RacingCar) ChooseColor(color string) string {
	return r.supersport.ChooseColor(color)
}

// SelectComfort delega la implementación de la interfaz Vehicle al objeto supersport (que debe ser de tipo Vehicle).
func (r *RacingCar) SelectComfort(comfort string) string {
	return r.supersport.SelectComfort(comfort)
}

//Funcion principal
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
