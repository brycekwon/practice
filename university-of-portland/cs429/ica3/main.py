# LINEAR REGRESSION


import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

########################################
# Preparation                          #
########################################

data = pd.read_csv("birthrate.dat")
# print(data)

rows = data.shape[0]
# print(rows)

cols = data.shape[1]
# print(cols)

# for consistency, make vectors all upper-case and scalars lowercase
data = data.values
X = data[:,2]
Y = data[:,1]
# print(X, Y)

# pseudorandom values
a1 = 0.2
a0 = 0.5

X_norm = np.true_divide(X, np.max(X))
# print(X_norm)

Y_norm = np.true_divide(Y, np.max(Y))
# print(Y_norm)


########################################
# Model                                #
########################################

Y_hat = a1 * X_norm + a0
# print(Y_hat)

# error
E = Y_hat - Y_norm
# print(E)

# mean square error
MSE = np.true_divide(np.sum(np.square(E)), rows)
print(MSE)



# LEARNING!!!!!
lr = 0.1 # learning rate

a0 = a0 - lr * np.true_divide(np.sum(E), rows)
# print(a0)

a1 = a1 - lr * np.true_divide(np.sum(E * X_norm), rows)
# print(a1)


Y_hat = a1 * X_norm + a0
# print(Y_hat)

# error
E = Y_hat - Y_norm
# print(E)

# mean square error
MSE = np.true_divide(np.sum(np.square(E)), rows)
print(MSE)


# plt.scatter(X, Y)
# plt.plot(X, Y_hat)
#
# plt.show()
