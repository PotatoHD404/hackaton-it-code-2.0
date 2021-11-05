package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	function "hackaton-it-code-2.0/src/api/handler"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	handler := mux.NewRouter()
	handler.PathPrefix("/").HandlerFunc(function.Handler)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
