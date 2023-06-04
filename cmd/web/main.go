package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "Http Network Address")
	flag.Parse()

	// A big benefit of logging your messages to the standard streams (stdout and stderr) like we are is that your
	// application and logging are decoupled. Your application itself isn’t concerned with the routing or storage of
	// the logs, and that can make it easier to manage the logs differently depending on the environment
	// go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
	infoLog := log.New(os.Stdin, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	// Create file server to serve files out of "./ui/static"
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// Create handler from the file server that serves all requests to /static/,
	// stripping "/static" before a request reaches the file server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on http://127.0.0.1%s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
