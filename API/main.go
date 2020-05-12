package main

import (
	"log"
	"net/http"

	"data-eq1/API/router"
)

func main() {
	r := router.Init()
	log.Println("Serving on localhost:80")
	log.Fatal(http.ListenAndServe(":80", r))
}
