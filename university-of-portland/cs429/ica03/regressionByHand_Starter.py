import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

plt.xlabel('income')
plt.ylabel('mortality')
plt.title('Linear Regression: income vs. infant mortality')

data = pd.read_csv('birthrate.dat')

rows = data.shape [0]
cols = data.shape [1]

data = data.values
data = data[np.arange(0,rows),:]

X = data[:, 2]
Y = data[:, 1]
X_max = np.max(X)
Y_max = np.max(Y)
X = np.true_divide (X, X_max)
Y = np.true_divide (Y, Y_max)
plt.xlim(0,max(X))
plt.ylim(0,max(Y))

print(X)
print(Y)

b1 = -1.0
b0 =  0.5
mu = .01

plt.scatter(X, Y)
plt.plot (X, b1*X + b0)
print (" b0: " + str(b0) + " b1: " + str(b1) + " Error: " + str(0))
plt.pause(0.1)


# #batch 1 == iteration 1
# #United States,24.7,1723,0.12,27.2
# #Argentina,24.7,287,0.20,62.0
# #New Zealand,24.4,970,0.19,24.9
# Load X and Y vectors for the above data
# normalize the data
# calculate Error and update parameters 

# plt.plot (X, b1*X + b0)

# #batch 2 == iteration 2
# #Venezuela,46.4,392,0.40,68.5
# #Mexico,45.7,118,0.61,87.8
# #Ecuador,45.3,44,0.53,115.8
# Load X and Y vectors for the above data
# normalize the data
# calculate Error and update parameters 

# plt.scatter([X1, X2, X3],[Y1,Y2,Y3])

# plt.plot (X, b1*X + b0)

# #batch 3 == iteration 3
# #Greece,18.8,134,0.52,47.4
# #Norway,18.6,633,0.19,21.7
# #Italy,18.0,295,0.44,55.7
# Load X and Y vectors for the above data
# normalize the data
# calculate Error and update parameters 

# plt.scatter([X1, X2, X3],[Y1,Y2,Y3])

# plt.plot (X, b1*X + b0)

plt.show()
