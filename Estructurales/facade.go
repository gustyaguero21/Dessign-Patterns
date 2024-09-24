package main

// User representa la estructura de un usuario con campos como nombre de usuario, correo electrónico y contraseña.
type User struct {
	Username string
	Email    string
	Password string
}

// Car representa la estructura de un automóvil con campos como marca, modelo y color.
type Car struct {
	Brand string
	Model string
	Color string
}

// Location representa la estructura de una ubicación con campos como calle, número y provincia.
type Location struct {
	Street   string
	Number   int
	Province string
}

// NewRegister agrupa los registros de un usuario, un auto y una ubicación.
// Contiene como campos las estructuras User, Car y Location.
type NewRegister struct {
	User     User
	Car      Car
	Location Location
}

// CreateNewRegister se encarga de crear un nuevo registro asignando los valores de los parámetros
// a los campos correspondientes de las estructuras User, Car y Location anidadas en NewRegister.
func (nr *NewRegister) CreateNewRegister(name, email, password, brand, color, model, street, province string, number int) {
	nr.User = User{
		Username: name + "21", // Asigna un nombre de usuario basado en el nombre y agrega "21".
		Email:    email,
		Password: password,
	}

	nr.Car = Car{
		Brand: brand,
		Model: model,
		Color: color,
	}

	nr.Location = Location{
		Street:   street,
		Number:   number,
		Province: province,
	}
}

// Función principal
// func main() {
// 	newRegister := NewRegister{} // Crea una nueva instancia de NewRegister.

// 	newRegister.CreateNewRegister( // Utiliza la función CreateNewRegister para llenar los campos del registro.
// 		"John", "john@example.com", "password123",
// 		"Chevrolet", "Astra", "Plateado",
// 		"Main St", "NY", 123,
// 	)

// 	fmt.Println("User:", newRegister.User)
// 	fmt.Println("Car:", newRegister.Car)
// 	fmt.Println("Location:", newRegister.Location)
// }
