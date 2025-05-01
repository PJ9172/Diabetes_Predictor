package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server Starts on Port 8000!!!")
	http.ListenAndServe(":8000",nil)
}