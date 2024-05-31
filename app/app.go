package app

import (
	"log"
	"net/http"
)

func Start() {
	addr := "localhost:8000"
	log.Fatal(http.ListenAndServe(addr, router()))
}
