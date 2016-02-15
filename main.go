package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	Success = iota
	Failed
)

func main() {
	inputFile, err := ioutil.ReadFile("data/input.txt")
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(Failed)
	}
	contents := string(inputFile)
	contents = strings.Trim(contents, "\n")
	size := strings.Count(contents, "\n") + 1
	if !isSqureMatrix(size) {
		fmt.Println("Matrix is not valid")
		os.Exit(Failed)
	}
	s := int(math.Sqrt(float64(size)))
	var matrix [10][10]int64 // max input matrix size is 10*10
	data := strings.Split(contents, "\n")
	for _, value := range data {
		lines := strings.Fields(value)
		row, _ := strconv.ParseInt(lines[0], 10, 64)
		col, _ := strconv.ParseInt(lines[1], 10, 64)
		val, _ := strconv.ParseInt(lines[2], 10, 64)
		matrix[row][col] = val
	}
	printMatrix(matrix, "A =>", s)
	tranMatrix := transpose(matrix, s)
	printMatrix(tranMatrix, "A`", s)
	result := multiply(matrix, tranMatrix, s)
	printMatrix(result, "A*A`", s)
}

func transpose(matrix [10][10]int64, size int) [10][10]int64 {
	var transpose [10][10]int64
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			transpose[i][j] = matrix[j][i]
		}
	}
	return transpose
}

func multiply(a, b [10][10]int64, size int) [10][10]int64 {
	var multiply [10][10]int64
	var sum int64
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				sum = sum + a[i][k]*b[k][j]
			}
			multiply[i][j] = sum
			sum = 0
		}

	}
	return multiply
}

func printMatrix(a [10][10]int64, b string, size int) {
	fmt.Println(b)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print("\t", a[i][j], " ")
		}
		fmt.Println()
	}

}
func isSqureMatrix(size int) bool {
	sqrt := int(math.Sqrt(float64(size)))
	return sqrt*sqrt == size
}
