package main

import (
	"github.com/joho/godotenv"
	function "hackaton-it-code-2.0/src/api/handler"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	function.InitProject()
	handler := function.NewHttpHandler()
	log.Fatal(http.ListenAndServe(":3000", handler))
}
