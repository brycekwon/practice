package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if _, err := os.Stat("gooutput.txt"); err == nil {
		os.Remove("gooutput.txt")
	}

	startProcessingTime := time.Now()

	PreprocessImages("./no_aurora", "gooutput.txt", "-1:")
	PreprocessImages("./yes_aurora", "gooutput.txt", "1:")

	L, X, err := Parser("gooutput.txt")
	if err != nil {
		log.Fatal("failed parsing output file")
	}
	X_norm := Normalize(X)
	L, X_norm = Randomize(L, X_norm)

	fmt.Printf("Processing Time: %d\n", time.Since(startProcessingTime))

	partition := int(float32(len(X_norm)) * 0.8)
	X_norm_train := X_norm[:partition]
	X_norm_test := X_norm[partition:]
	L_train := L[:partition]
	L_test := L[partition:]

	startTrainingTime := time.Now()

	W, _, _ := TrainPerceptron(L_train, X_norm_train, 20, 0.01)

	fmt.Printf("Training Time: %d\n", time.Since(startTrainingTime))

	// fmt.Println(CA)
	// fmt.Println(TE)

	K := make([]float32, len(L_test))
	for i, e := range X_norm_test {
		if TestPerceptron(W, e) == L_test[i] {
			K[i] = 1
		} else {
			K[i] = 0
		}
	}
	correct := int(0)
	wrong := int(0)
	for _, val := range K {
		if val == 1 {
			correct++
		} else {
			wrong++
		}
	}
	fmt.Printf("%v correct\t%v wrong\n", correct, wrong)
}
