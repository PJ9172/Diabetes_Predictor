package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Input struct {
	Pregnancies   int     `json:"Pregnancies"`
	Glucose       int     `json:"Glucose"`
	BloodPressure int     `json:"Blood_Pressure"`
	BMI           float64 `json:"BMI"`
	Age           int     `json:"Age"`
}

type PredictionResponse struct {
	Result string `json:"result"`
}

// convert string to int
func parseInt(val string) int {
	i, _ := strconv.Atoi(val)
	return i
}

// convert string to float64
func parseFloat(val string) float64 {
	f, _ := strconv.ParseFloat(val, 64)
	return f
}

func Prediction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// parse the form values
	err := r.ParseForm()
	if err != nil {
		log.Println("Form parse error:", err)
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	log.Println("pregnancies: ",r.FormValue("pregnancies"))

	// store form values in struct
	input := Input{
		Pregnancies:   parseInt(r.FormValue("pregnancies")),
		Glucose:       parseInt(r.FormValue("glucose")),
		BloodPressure: parseInt(r.FormValue("blood_pressure")),
		BMI:           parseFloat(r.FormValue("bmi")),
		Age:           parseInt(r.FormValue("age")),
	}
	// convert struct to json
	jsonData, err := json.Marshal(input)
	if err != nil {
		http.Error(w, "Error creating JSON", http.StatusInternalServerError)
		return
	}
	// send json data to fast api
	resp, err := http.Post("http://localhost:8000/predict", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Error to call fast api!!!", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result PredictionResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Error decoding response:", err)
		http.Error(w, "Error decoding prediction", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	err = tmpl.Execute(w, result)
	if err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}
