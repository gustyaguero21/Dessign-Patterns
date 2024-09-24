package main

import (
	"sync"
)

type Admin struct { // Estructura con campos del admin. Ej: Un administrador de Bases de datos.
	Username string
	Password string
}

var database_admin *Admin // Variable de acceso global.
var once sync.Once        // Metodo para ejecutar una linea de codigo solamente una vez.

func CreateNewAdmin(username, password string) *Admin { // Esta funcion se encarga de crear una unica instancia de nuestra variable global.
	once.Do(func() {
		database_admin = &Admin{Username: username, Password: password}
	})

	return database_admin
}

// func main() {
// 	newDatabaseAdmin1 := CreateNewAdmin("Gustyaguero21", "1234")
// 	fmt.Println(newDatabaseAdmin1.Username, newDatabaseAdmin1.Password) // Imprime "Gustyaguero21 1234"

// 	newDatabaseAdmin2 := CreateNewAdmin("NuevoUsuario", "4321")
// 	fmt.Println(newDatabaseAdmin2.Username, newDatabaseAdmin2.Password) // También imprime "Gustyaguero21 1234"
// }
