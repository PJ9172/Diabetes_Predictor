import pandas as pd
from sklearn.linear_model import LogisticRegression
from sklearn.model_selection import train_test_split
import joblib

# Load dataset
df = pd.read_csv("diabetes.csv")

# Select relevant features
X = df[["Pregnancies", "Glucose", "BloodPressure", "BMI", "Age"]]
y = df["Outcome"]

# Train model
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2)
model = LogisticRegression()
model.fit(X_train, y_train)

# Save model
joblib.dump(model, "model.joblib")
print("✅ Model saved as model.joblib")
