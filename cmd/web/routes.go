package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/todo/create", create)
	mux.HandleFunc("/todo/delete", delete)
	mux.HandleFunc("/todo/update", update)

	return mux
}
