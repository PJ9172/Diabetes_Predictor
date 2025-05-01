import pandas as pd
from sklearn.linear_model import LinearRegression
from sklearn.model_selection import train_test_split
import pickle

# load .csv data file
data = pd.read_csv("diabetes.csv")

# select features and target
features = ['Pregnancies', 'Glucose', 'BloodPressure', 'BMI', 'Age']
target = 'Outcome'

# drop rows with missing values
data = data[features + [target]].dropna()

# split data into input (inp) and target (tar)
inp = data[features]
tar = data[target]

# split into training and test sets
inp_train, inp_test, tar_train, tar_test = train_test_split(inp, tar, test_size=0.2, random_state=42)

# train the model
model = LinearRegression()
model.fit(inp_train, tar_train)

# save the model
with open("model.joblib",'wb') as m:
    pickle.dump(model, m)
    
print("Model trained & save to model.joblib")