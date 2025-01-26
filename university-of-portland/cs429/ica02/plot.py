import numpy as np
import matplotlib.pyplot as plt


# fig = plt.figure()
# ax = plt.axes(projection='3d')
# data = np.loadtxt('data.csv', delimiter=',', skiprows=2, usecols=(0, 1, 2))
#
# ax.set_xlabel('a0')
# ax.set_ylabel('a1')
# ax.set_zlabel('error')
#
# ax.scatter(data[:, 0], data[:, 1], data[:, 2])
#
# plt.show()

###############################################################################
# EXAMPLE                                                                     #
###############################################################################

X = np.array([150, 175, 157, 355, 211, 115, 107, 203, 145, 137, 125, 122])
Y = np.array([8, 7, 8, 4, 6, 11, 12, 6, 5, 9, 9, 10])
R = np.array([])

fig = plt.figure()
ax = plt.axes(projection='3d')

for a0 in np.arange(-0.05, 0.05, 0.01):
    for a1 in np.arange(-20, 20, 1):
        Y_hat = a0*X+a1

        row = [a0, a1, np.sum(np.power((Y-Y_hat), 2), 0)]
        R = np.append(R, row)

R = R.reshape(int(R.size / 3), 3)
print(R)

print(R.shape)

ax.set_xlabel("a0")
ax.set_ylabel("a1")
ax.set_zlabel("error")

ax.scatter(R[:, 0], R[:, 1], R[:, 2])
plt.show()
