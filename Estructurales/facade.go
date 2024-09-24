package main

type User struct { //Estructura de un Usuario.
	Username string
	Email    string
	Password string
}

type Car struct { //Estructura de un Auto.
	Brand string
	Model string
	Color string
}

type Location struct { //Estructura de una locacion.
	Street   string
	Number   int
	Province string
}

type NewRegister struct { //Estructura de registro que contienen como campos a las estructuras anteriores.
	User     User
	Car      Car
	Location Location
}

// Esta funcion se encarga de crear un nuevo registro utilizando la estructura NewRegister y asociando los campos de las estructuras anidadas, a los datos que llegan por parametros.

func (nr *NewRegister) CreateNewRegister(name, email, password, brand, color, model, street, province string, number int) {
	nr.User = User{
		Username: name + "21",
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
// 	newRegister := NewRegister{} //Creamos un nuevo registro.

// 	newRegister.CreateNewRegister(	// Utilizamos la funcion CreateNewRegister para crear un registro.
// 		"John", "john@example.com", "password123",
// 		"Chevrolet", "Astra", "Plateado",
// 		"Main St", "NY", 123,
// 	)

// 	fmt.Println("User:", newRegister.User)
// 	fmt.Println("Car:", newRegister.Car)
// 	fmt.Println("Location:", newRegister.Location)
// }
