package main

import (
	"fmt"
	"log"
	"time"
)

func SortedSquaredArray(array []int) []int {
	response := make([]int, len(array))

	i, j := 0, len(array)-1

	counter := len(array) - 1

	for i != j {
		left := square(array[i])
		right := square(array[j])

		if left > right {
			response[counter] = left
			i++
		} else {
			response[counter] = right
			j--
		}

		counter--
	}
	response[counter] = square(array[i])
	return response
}

func square(x int) int {
	return x * x
}

func main() {
	start := time.Now()
	fmt.Println(SortedSquaredArray([]int{1, 2, 3, 5, 6, 8, 9}))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
