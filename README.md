# 🩺 Diabetes Risk Predictor – Go + FastAPI + ML

A lightweight, full-stack machine learning project that predicts diabetes risk using user health inputs. Built with **Go** for frontend routing and **FastAPI** as a Python microservice for ML-based classification.

---

## 📌 Features

- 🧾 **Input Form :** Takes user health data like glucose, BMI, age, etc.
- 🤖 **ML Prediction :** Predicts if a person is at high or low risk of diabetes
- 🔁 **Microservice Architecture :** Go handles web UI, FastAPI handles model inference
- 🎨 **Clean UI :** Built with Bootstrap for responsiveness and clarity
- ⚡ **Fast & Lightweight :** Form → FastAPI → Prediction in milliseconds

---

## 🏗️ Tech Stack

| Layer          | Tech Used                   |
|----------------|-----------------------------|
| Frontend       | Go, Gorilla Mux, HTML       |
| Backend        | Python, FastAPI             |
| ML Model       | scikit-learn, pandas, numpy |
| UI Styling     | Bootstrap 5                 |
| Model Type     | Logistic Regression         |

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/diabetes-predictor.git
cd diabetes-predictor
```

### 2. Train the ML Model
```bash
cd ml-api
python train_model.py
```
This generates model.joblib from the diabetes dataset

### 3. Start the FastAPI Server
```bash
uvicorn main:app --reload
```
This runs the prediction microservice on `http://localhost:8000`

### 4. Start the Go Web Server
```bash
cd ../go-webapp
go run main.go
```
This runs your UI web server at `http://localhost:3000/form`


## 🔍 How It Works
1. User fills out a form (Pregnancies, Glucose, BMI, Age, etc.)
2. Form data is sent from Go to FastAPI (JSON)
3. FastAPI loads the trained model and returns prediction
4. Go renders the result: "Low Risk" or "High Risk"

## 🧠 ML Model Details
- Dataset : PIMA ndians Diabetes Dataset (Kaggle)
- Algorithm : Logistic Regression
- Target : Outcome (0 = No diabetes, 1 = At risk)
- Input Features : Glucose, BMI, BloodPressure, Age, Pregnancies

## 📁 Project Structure
```bash
diabetes-predictor/
├── Screenshots/
├── go-webapp/
│   ├── main.go
│   ├── templates/
│   │   ├── input_form.html
│   │   └── result.html
│   ├── routers/
│   │   └── routers.go
│   └── handlers/
│       └── handlers.go
├── ml-api/
│   ├── diabetes.csv
│   ├── main.py
│   ├── train_model.py
│   └── model.joblib
```

## 📸 Screenshots
- Input_form Page
    !['input_form'](/Screenshots/input_form.png)
- Result Page
    !['result1'](/Screenshots/result1.png)
- Result Page
    !['result2'](/Screenshots/result2.png)
## 🤝 Credits
- Kaggle: PIMA Diabetes Dataset
- GoLang: Gorilla Mux router
- FastAPI: High-speed Python microservice
- Bootstrap: For responsive UI
- scikit-learn: Model training and prediction