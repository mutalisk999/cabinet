package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	latestTag        string = ""
	latestCommitSHA1 string = ""
	releaseType      string = ""
)

func main() {
	endpoint := flag.String("e", "127.0.0.1:60070", "endpoint to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	log.Println("latestTag: ", latestTag)
	log.Println("releaseType: ", releaseType)
	log.Println("latestCommitSHA1: ", latestCommitSHA1)
	log.Printf("Serving [%s] on HTTP: %s\n", *directory, *endpoint)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(*directory)))
	server := http.Server{Addr: *endpoint, Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
	select {}
}
