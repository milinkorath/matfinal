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
	msize=10 // setting maximum array size
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("input file is missing.")
		os.Exit(Failed)
	}
	fname := os.Args[1]
	inputFile, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(Failed)
	}
	contents := string(inputFile)
	contents = strings.Trim(contents, "\n")
	size := strings.Count(contents, "\n") + 1  // actual size of the input matrix.Used for iteration
	if !isSqureMatrix(size) {
		fmt.Println("Matrix is not valid")
		os.Exit(Failed)
	}
	s := int(math.Sqrt(float64(size)))
	var matrix [msize][msize]int64 // max input matrix size is 10*10
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
// to find the transpose of the matrix
func transpose(matrix [msize][msize]int64, size int) [msize][msize]int64 {
	var transpose [msize][msize]int64
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			transpose[i][j] = matrix[j][i]
		}
	}
	return transpose
}
// to mutiply two matrix
func multiply(a, b [msize][msize]int64, size int) [msize][msize]int64 {
	var multiply [msize][msize]int64
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
// used to format the output
func printMatrix(a [msize][msize]int64, b string, size int) {
	fmt.Println(b)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print("\t", a[i][j], " ")
		}
		fmt.Println()
	}

}
// used to check square matrix
func isSqureMatrix(size int) bool {
	sqrt := int(math.Sqrt(float64(size)))
	return sqrt*sqrt == size
}
