package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/todo/create", app.create)
	mux.HandleFunc("/todo/delete", app.delete)
	mux.HandleFunc("/todo/update", update)

	return mux
}
