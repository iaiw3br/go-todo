package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

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
