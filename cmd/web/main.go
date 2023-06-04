package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"snippetbox.badrchoubai.dev/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

// Initialize a new instance of our application struct, containing dependencies.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "Http Network Address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySql Data Source Name")

	flag.Parse()

	// A big benefit of logging your messages to the standard streams (stdout and stderr) like we are is that your
	// application and logging are decoupled. Your application itself isn’t concerned with the routing or storage of
	// the logs, and that can make it easier to manage the logs differently depending on the environment
	// go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
	infoLog := log.New(os.Stdin, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDb(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// Initialize a new template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on http://127.0.0.1%s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
