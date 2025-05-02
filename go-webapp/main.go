package main

import (
	"Diabetes_Predictor/routers"
	"log"
	"net/http"
)

func main() {
	r := routers.SetRouter()

	log.Println("Server Starts at http://localhost:3000/form")
	log.Fatal(http.ListenAndServe(":3000", r))
}
