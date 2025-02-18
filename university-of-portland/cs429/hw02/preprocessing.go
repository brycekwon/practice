package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"gocv.io/x/gocv"
)

func DeriveColorHistogram(filename string) ([]float32, error) {
	image := gocv.IMRead(filename, gocv.IMReadColor)
	if image.Empty() {
		return nil, fmt.Errorf("could not read image from file '%s'", filename)
	}
	defer image.Close()

	imgChannels := gocv.Split(image)
	if len(imgChannels) != 3 {
		return nil, fmt.Errorf("image does not have 3 channels")
	}

	histRe, histGr, histBl := gocv.NewMat(), gocv.NewMat(), gocv.NewMat()
	gocv.CalcHist([]gocv.Mat{imgChannels[2]}, []int{0}, gocv.NewMat(), &histRe, []int{256}, []float64{0, 256}, false)
	gocv.CalcHist([]gocv.Mat{imgChannels[1]}, []int{0}, gocv.NewMat(), &histGr, []int{256}, []float64{0, 256}, false)
	gocv.CalcHist([]gocv.Mat{imgChannels[0]}, []int{0}, gocv.NewMat(), &histBl, []int{256}, []float64{0, 256}, false)

	histAll := append(matToArray(histRe), matToArray(histGr)...)
	histAll = append(histAll, matToArray(histBl)...)

	return histAll, nil
}

func PreprocessImages(imageDir string, outputfile string, prepend string) error {
	files, err := os.ReadDir(imageDir)
	if err != nil {
		return fmt.Errorf("could not read image directory")
	}

	var wg sync.WaitGroup
	resultsCh := make(chan []float32, len(files))

	for _, file := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			hist, _ := DeriveColorHistogram(filepath.Join(imageDir, filename))
			resultsCh <- hist
		}(file.Name())
	}

	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	outfile, err := os.OpenFile(outputfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not create output file")
	}
	defer outfile.Close()

	for hist := range resultsCh {
		outfile.WriteString(prepend + strings.Join(arrayToString(hist), ",") + "\n")
	}

	return nil
}

func Parser(output string) ([]int, [][]float32, error) {
	file, err := os.Open(output)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open output file")
	}
	defer file.Close()

	var labels []int
	var histograms [][]float32

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)

		if len(parts) != 2 {
			continue
		}

		label, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("could not convert string to int")
		}
		labels = append(labels, label)

		data := strings.Split(parts[1], ",")
		histogram := make([]float32, len(data))
		for i, v := range data {
			value, err := strconv.ParseFloat(v, 32)
			if err != nil {
				return nil, nil, fmt.Errorf("could not convert string to float")
			}
			histogram[i] = float32(value)
		}
		histograms = append(histograms, histogram)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("could not read file")
	}

	return labels, histograms, nil
}

func matToArray(mat gocv.Mat) []float32 {
	rows, cols := mat.Rows(), mat.Cols()
	data, _ := mat.DataPtrFloat32()

	result := make([]float32, rows*cols)
	copy(result, data[:rows*cols])

	return result
}

func arrayToString(input []float32) []string {
	result := make([]string, len(input))
	for i, v := range input {
		result[i] = fmt.Sprintf("%.1f", v)
	}
	return result
}
