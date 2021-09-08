package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.NotFound(w, r)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	todos, err := app.TodoList.GetAll()

	if err != nil {
		app.errorLog.Fatal(err)
		return
	}

	for _, todo := range todos {
		fmt.Fprintf(w, "%v\n", todo)
	}
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
