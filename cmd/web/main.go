package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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

func main() {
	localAddress := flag.String("localAddress", ":8080", "Адрес веб-сервиса")
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		errorLog: errLog,
		infoLog:  infoLog,
	}

	server := http.Server{
		Addr:     *localAddress,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Сервер запущен")
	err := server.ListenAndServe()
	errLog.Fatal(err)
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}
