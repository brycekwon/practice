import pandas as pd


df = pd.read_csv("./dataset.csv")

q3 = df.quantile(0.75).values[0]
q1 = df.quantile(0.25).values[0]
IQR = q3 - q1

upper = q3 + 1.5 * IQR
lower =  q1 - 1.5 * IQR

# print(upper)
# print(lower)

mean = df.mean().values[0]
outliers = df.loc[(df.Data > upper) | (df.Data < lower)]
print(outliers)
