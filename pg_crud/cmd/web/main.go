package main

import (
	"flag"
	"log"
	"net/http"

	"os"

	"wallester.eu/snippetbox/pkg/models"
	"wallester.eu/snippetbox/pkg/models/crud"
)

type application struct {
	errorLog   *log.Logger
	infoLog    *log.Logger
	operations *crud.CrudMOdel
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	//dsn := flag.String("dsn", "postgres:admin@snippetbox?parseTime=true", "PostgresSQl data")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := models.InitDB("postgres://postgres:admin@localhost/customers_db")

	if err != nil {
		errorLog.Fatal(err)
	}

	// db, err := openDB(*dsn)
	// if err != nil {
	// 	errorLog.Fatal(err)
	// }

	defer db.Close()

	app := &application{
		errorLog:   errorLog,
		infoLog:    infoLog,
		operations: &crud.CrudMOdel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on ................. %s", *addr)
	infoLog.Printf("Connected to database Successfully ......! ")
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
