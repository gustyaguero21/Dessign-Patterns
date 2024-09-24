package main

import (
	"sync"
)

// Admin representa una estructura que contiene las credenciales de un administrador (por ejemplo, de una base de datos).
type Admin struct {
	Username string
	Password string
}

var database_admin *Admin // Variable que almacena la instancia única de Admin (singleton).
var once sync.Once        // Variable usada para asegurar que la inicialización solo se realice una vez.

func CreateNewAdmin(username, password string) *Admin {
	// CreateNewAdmin crea una única instancia de Admin si no existe (patrón Singleton).
	once.Do(func() {
		database_admin = &Admin{Username: username, Password: password}
	})

	return database_admin
}

//Funcion principal
// func main() {
// 	newDatabaseAdmin1 := CreateNewAdmin("Gustyaguero21", "1234")
// 	fmt.Println(newDatabaseAdmin1.Username, newDatabaseAdmin1.Password) // Imprime "Gustyaguero21 1234"

// 	newDatabaseAdmin2 := CreateNewAdmin("NuevoUsuario", "4321")
// 	fmt.Println(newDatabaseAdmin2.Username, newDatabaseAdmin2.Password) // También imprime "Gustyaguero21 1234" ya que la instancia es única.
// }
