package main

import "math/rand/v2"

func Normalize(X [][]float32) [][]float32 {
	for _, e := range X {
		max := float32(0)
		for _, val := range e {
			if val > max {
				max = val
			}
		}
		for j := range e {
			e[j] /= max
		}
	}

	return X
}

func Randomize(L []int, X [][]float32) ([]int, [][]float32) {
	idx := rand.Perm(len(X))
	for i, j := range idx {
		X[i] = X[j]
		L[i] = L[j]
	}

	return L, X
}

func TrainPerceptron(L []int, X [][]float32, N int, mu float32) ([]float32, []float32, []float32) {
	rows, cols := len(X), len(X[0])

	W := make([]float32, cols+1)
	for i := range W {
		W[i] = rand.Float32()
	}

	CA := make([]float32, N)
	TE := make([]float32, N)

	var acc float32
	for i := 0; i < N; i++ {
		acc = 0
		for j := 0; j < rows; j++ {
			change := W[0] + sum(W[1:], X[j])

			prediction := float32(1)
			if change < 0 {
				prediction = -1
			}

			if prediction == float32(L[j]) {
				acc++
			} else {
				E := prediction - float32(L[j])
				X_t := append([]float32{1}, X[j]...)
				W = subtract(W, multiply(E, X_t, mu))
			}
		}

		CA[i] = (acc / float32(rows)) * 100
		TE[i] = ((float32(rows) - acc) / float32(rows)) * 100

	}

	return W, CA, TE
}

func TestPerceptron(W []float32, hist []float32) int {
	max := float32(0)
	for _, val := range hist {
		if val > max {
			max = val
		}
	}
	for i := range hist {
		hist[i] /= max
	}

	P := W[0] + sum(W[1:], hist)
	if P >= 0 {
		return 1
	} else {
		return -1
	}
}

func sum(a []float32, b []float32) float32 {
	var result float32
	for i := 0; i < len(a); i++ {
		result += a[i] * b[i]
	}
	return result
}

func subtract(a, b []float32) []float32 {
	result := make([]float32, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = a[i] - b[i]
	}
	return result
}

func multiply(a float32, b []float32, mu float32) []float32 {
	result := make([]float32, len(b))
	for i := 0; i < len(b); i++ {
		result[i] = a * b[i] * mu
	}
	return result
}
