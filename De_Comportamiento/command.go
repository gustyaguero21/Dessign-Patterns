package main

import "fmt"

type Command interface { //Esta es la interfaz que define el método Execute, que todos los comandos deben implementar.
	Execute()
}

type Light struct{}

func (l *Light) On() { //Método para encender la luz.
	fmt.Println("Light is on.")
}

func (l *Light) Off() { //Método para apagar la luz.
	fmt.Println("Light is off.")
}

type OnCommand struct { // Contiene una referencia a la luz y llamará al método On.
	light Light
}

type OffCommand struct { //Esta es la implementación del comando para apagar la luz.
	light Light
}

func (on *OnCommand) Execute() { //Implementación del método Execute para el comando On.
	on.light.On()
}

func (off *OffCommand) Execute() { //Este método llamará al método Off del receptor (Light).
	off.light.Off()
}

type RemoteControl struct { //Esta es la clase que actúa como el "Invoker" en el patrón Command.
	command Command
}

func (r *RemoteControl) SetCommand(command Command) { //Permite cambiar el comando asociado al control remoto.
	r.command = command
}

func (r *RemoteControl) PressButton() { // Cuando se presiona el botón, el control remoto ejecuta
	r.command.Execute()
}

// // Función principal (main)
// func main() {
// 	// Creación de una luz (el receptor).
// 	lightOff := Light{}

// 	// Creación del comando Off para la luz.
// 	offCommand := OffCommand{light: lightOff}

// 	// Creación de un control remoto con el comando Off.
// 	newRemoteControl := RemoteControl{command: &offCommand}

// 	// Configuramos el comando actual del control remoto como "offCommand" (apagar la luz).
// 	newRemoteControl.SetCommand(&offCommand)

// 	// Simulación de presionar el botón, lo que ejecutará el comando de apagar la luz.
// 	newRemoteControl.PressButton()
// }
