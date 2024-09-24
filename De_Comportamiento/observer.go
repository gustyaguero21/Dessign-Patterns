package main

import "fmt"

// Subscriber representa al observador (patrón Observer), quien recibirá notificaciones de actualizaciones del sujeto.
type Subscriber struct {
	name string
}

// Update es el método utilizado para notificar al suscriptor cuando se publica un nuevo video en el canal.
// Actualiza el estado del suscriptor con el título del video.
func (s *Subscriber) Update(videoTitle string) {
	fmt.Printf("%s: ¡Nuevo video publicado! Título: %s\n", s.name, videoTitle)
}

// YouTubeChannel representa el sujeto (observable) que notifica a los suscriptores cuando hay un nuevo video.
type YouTubeChannel struct {
	subscribers []Subscriber // Lista de suscriptores que recibirán las notificaciones.
}

// Register agrega un nuevo suscriptor a la lista de suscriptores del canal.
func (yt *YouTubeChannel) Register(subscriber Subscriber) {
	yt.subscribers = append(yt.subscribers, subscriber)
}

// UpdateNewVideo simula la publicación de un nuevo video y notifica a los suscriptores.
func (yt *YouTubeChannel) UpdateNewVideo(title string) {
	fmt.Println("Publicando video.")
	yt.Notify(title)
}

// Notify recorre la lista de suscriptores y llama al método Update para notificarles sobre el nuevo video.
func (yt *YouTubeChannel) Notify(videoTitle string) {
	for _, subscriber := range yt.subscribers {
		subscriber.Update(videoTitle)
	}
}

// Función principal
// func main() {
// 	subscriber1 := Subscriber{}
// 	subscriber1.name = "Gustavo"

// 	subscriber2 := Subscriber{}
// 	subscriber2.name = "Melisa"

// 	youtubeChannel := YouTubeChannel{}

// 	youtubeChannel.Register(subscriber1)
// 	youtubeChannel.Register(subscriber2)

// 	youtubeChannel.UpdateNewVideo("Programando en Go.")

// 	youtubeChannel.UpdateNewVideo("Programando en Python.")
// }
