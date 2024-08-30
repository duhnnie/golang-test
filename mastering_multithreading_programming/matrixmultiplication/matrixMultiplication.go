package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const matrixSize = 250

var matrixA = [matrixSize][matrixSize]int {}
var matrixB = [matrixSize][matrixSize]int {}
var result [matrixSize][matrixSize]int
var lock = sync.RWMutex{}
var condition = sync.NewCond(lock.RLocker())
var waitGroup = sync.WaitGroup{}

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row ++ {
		for col := 0; col < matrixSize; col ++ {
			matrix[row][col] = rand.Intn(11) - 5
		}
	}
}

func rowWorkout(row int) {
	for {
		lock.RLock()
		waitGroup.Done()
		condition.Wait()

		for col := 0; col <matrixSize; col ++ {
			for i := 0; i < matrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}

		lock.RUnlock()
	}
}

func main() {
	start := time.Now()

	for row := 0; row < matrixSize; row ++ {
		waitGroup.Add(1)
		go rowWorkout(row)
	}

	for i := 0; i < 100; i++ {
		waitGroup.Wait()
		lock.Lock()
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		waitGroup.Add(matrixSize)
		lock.Unlock()
		condition.Broadcast()
	}

	elapsedTime := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsedTime)
}