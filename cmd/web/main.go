package main

import (
	"database/sql"
	"flag"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go-todo/pkg/models/pg"
	"log"
	"net/http"
	"os"
)

func main() {
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	port, err := getEnvVariable("PORT")
	if err != nil {
		errLog.Fatal(err)
	}
	pgName, err := getEnvVariable("POSTGRES_NAME")
	if err != nil {
		errLog.Fatal(err)
	}
	pgPassword, err := getEnvVariable("POSTGRES_PASSWORD")
	if err != nil {
		errLog.Fatal(err)
	}
	pgDatabase, err := getEnvVariable("POSTGRES_DATABASE")
	if err != nil {
		errLog.Fatal(err)
	}
	localAddress := flag.String("localAddress", ":"+port, "Адрес веб-сервиса")
	connectionToDB := "postgres://" + pgName + ":" + pgPassword + "@localhost/" + pgDatabase + "?sslmode=disable"

	db, err := openDB(connectionToDB)

	if err != nil {
		errLog.Fatal(err)
		return
	}

	defer db.Close()

	app := application{
		errorLog: errLog,
		infoLog:  infoLog,
		TodoList: &pg.TodoModel{DB: db},
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
	TodoList *pg.TodoModel
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

func getEnvVariable(key string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}
