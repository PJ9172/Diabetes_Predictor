from fastapi import FastAPI
from pydantic import BaseModel
import pickle
import numpy as np

# load the trained model
with open("model.joblib", "rb") as m:
    model = pickle.load(m)

# define fastapi app
app = FastAPI()

# define input structure 
class DiabetesFeatures(BaseModel):
    Pregnancies: int
    Glucose: int
    BloodPressure: int
    BMI: float #Body mass index (weight in kg/(height in m)^2)
    Age: int
    
# endpoint
@app.post("/predict")
def predict_diabetes(data: DiabetesFeatures):
    features = np.array([[data.Pregnancies, data.Glucose, data.BloodPressure, data.BMI, data.Age]])
    prediction = model.predict(features)
    return {
        "Outcome" : prediction[0]
    }