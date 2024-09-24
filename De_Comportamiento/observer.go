package main

import "fmt"

type Subscriber struct { //el suscriptor es el observador, quien va a recibir el nuevo estado del sujeto.
	name string
}

func (s *Subscriber) Update(videoTitle string) { //este metodo Update, se utilizara para actualizar el estado de los suscriptores mediante el sujeto.
	fmt.Printf("%s: ¡Nuevo video publicado! Título: %s\n", s.name, videoTitle)
}

type YouTubeChannel struct {
	subscribers []Subscriber //aqui se alojaran todos los nuevos subscriptores que se agregan para notificarles los estados.
}

func (yt *YouTubeChannel) Register(subscriber Subscriber) { //este metodo guarda los suscriptores en un array de suscriptores.
	yt.subscribers = append(yt.subscribers, subscriber)
}

func (yt *YouTubeChannel) UpdateNewVideo(title string) { //este metodo se encarga de crear un estado y notificar a los suscriptores.
	fmt.Println("Publicando video.")
	yt.Notify(title)
}

func (yt *YouTubeChannel) Notify(videoTitle string) { //esta funcion recorre el array de suscriptores y emplea el metodo Update para enviarle el nuevo estado del sujeto.
	for _, subscriber := range yt.subscribers {
		subscriber.Update(videoTitle)
	}
}

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
