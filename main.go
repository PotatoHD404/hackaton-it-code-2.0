package main

import (
	"hackaton-it-code-2.0/src/it_code"
	"log"
	"net/http"
)

func main() {
	handler := it_code.NewHttpHandler()
	log.Fatal(http.ListenAndServe(":3000", handler))
}
