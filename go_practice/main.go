package main

import (
	"log"
	"net/http"
)

func main() {
	room := Newroom()

	http.Handle("/", room)
	go room.run()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
