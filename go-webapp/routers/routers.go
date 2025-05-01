package routers

import (
	"Diabetes_Predictor/handlers"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func SetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		temp, _ := template.ParseFiles("templates/input_form.html")
		temp.Execute(w, nil)
	})

	r.HandleFunc("/predict", handlers.Prediction).Methods("POST")

	return r
}
