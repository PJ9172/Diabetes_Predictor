from fastapi import FastAPI
from pydantic import BaseModel
import joblib
import numpy as np

app = FastAPI()
model = joblib.load("model.joblib")

class InputData(BaseModel):
    Pregnancies: int
    Glucose: int
    BloodPressure: int
    BMI: float
    Age: int

@app.post("/predict")
def predict(data: InputData):
    input_arr = np.array([[data.Pregnancies, data.Glucose, data.BloodPressure, data.BMI, data.Age]])
    prediction = model.predict(input_arr)[0]
    result = "😟⚠️ High Risk" if prediction == 1 else "😊 Low Risk"
    return {"result": result}
