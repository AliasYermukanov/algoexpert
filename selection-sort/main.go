package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(SelectionSort([]int{8, 5, 2, 9, 5, 6, 3}))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		ind := i
		for j := i; j < len(array); j++ {
			if array[j] < array[ind] {
				ind = j
			}
		}
		array[i], array[ind] = array[ind], array[i]
	}
	return array
}
