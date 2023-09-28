package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(InsertionSort([]int{8, 5, 2, 9, 5, 6, 3}))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func InsertionSort(array []int) []int {
	for i := range array {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	return array
}
