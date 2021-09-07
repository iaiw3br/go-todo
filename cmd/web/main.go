package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.NotFound(w, r)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Домашнаяя страница"))
}

func showAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Показать задачи"))
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Создать задачу"))
}

func delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Создать задачу"))
	w.Write([]byte("Удалить задачу"))
}

func update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Обновить задачу"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/todo", showAll)
	mux.HandleFunc("/todo/create", create)
	mux.HandleFunc("/todo/delete", delete)
	mux.HandleFunc("/todo/update", update)

	log.Println("Сервер запущен")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
