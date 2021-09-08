package main

import (
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"go-todo/pkg/models"
	"log"
	"net/http"
	"os"
)

func main() {
	localAddress := flag.String("localAddress", ":8080", "Адрес веб-сервиса")
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	connectionToDB := "postgres://postgres:postgres@localhost/go-todolist?sslmode=disable"

	db, err := openDB(connectionToDB)

	if err != nil {
		errLog.Fatal(err)
		return
	}

	defer db.Close()

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
	err = server.ListenAndServe()
	errLog.Fatal(err)
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	TodoList *models.TodoList
}

func openDB(connectionToDB string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionToDB)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
