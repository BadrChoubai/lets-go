package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Initialize a new instance of our application struct, containing dependencies.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "Http Network Address")
	flag.Parse()

	// A big benefit of logging your messages to the standard streams (stdout and stderr) like we are is that your
	// application and logging are decoupled. Your application itself isn’t concerned with the routing or storage of
	// the logs, and that can make it easier to manage the logs differently depending on the environment
	// go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
	infoLog := log.New(os.Stdin, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on http://127.0.0.1%s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
