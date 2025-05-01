package main

import (
	"Diabetes_Predictor/routers"
	"log"
	"net/http"
)

func main() {
	r := routers.SetRouter()

	log.Println("Server Starts at http://localhost:8000/form")
	log.Fatal(http.ListenAndServe(":8000", r))
}