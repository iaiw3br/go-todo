package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Домашнаяя страница"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Сервер запущен")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
