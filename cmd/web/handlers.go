package main

import (
	"fmt"
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

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.NotFound(w, r)
		return
	}
	id := r.URL.Query().Get("id")
	w.Write([]byte("Создать задачу"))
	fmt.Println("id", id)
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
