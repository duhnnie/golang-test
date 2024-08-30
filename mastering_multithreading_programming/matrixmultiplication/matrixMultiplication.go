package main

import (
	"fmt"
	"math/rand"
	"time"
)

const matrixSize = 250

var matrixA = [matrixSize][matrixSize]int {}
var matrixB = [matrixSize][matrixSize]int {}
var result [matrixSize][matrixSize]int

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row ++ {
		for col := 0; col < matrixSize; col ++ {
			matrix[row][col] = rand.Intn(11) - 5
		}
	}
}

func rowWorkout(row int) {
	for col := 0; col <matrixSize; col ++ {
		for i := 0; i < matrixSize; i++ {
			result[row][col] += matrixA[row][i] * matrixB[i][col]
		}
	}
}

func main() {
	start := time.Now()

	for i := 0; i < 300; i++ {
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)

		for row := 0; row < matrixSize; row ++ {
			rowWorkout(row)
		}
	}

	elapsedTime := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsedTime)
}