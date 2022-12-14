package main

import (
	"log"
	"net/http"
)

func main() {
	// create file server handler
	fs := http.FileServer(http.Dir("./files"))

	// start HTTP server with `fs` as the default handler
	log.Fatal(http.ListenAndServe(":80", fs))
}
