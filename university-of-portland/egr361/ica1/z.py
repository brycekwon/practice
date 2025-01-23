import pandas as pd


df = pd.read_csv("dataset.csv")

z = (df - df.mean()) / df.std()
print(z)

outliers = z.loc[(z.Data > 3) | (z.Data < -3)]
print(outliers)
