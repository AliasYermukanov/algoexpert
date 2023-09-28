package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func IsValidSubsequence(array []int, sequence []int) bool {
	counter := 0
	for _, value := range array {
		if value == sequence[counter] {
			counter++
			if counter == len(sequence) {
				return true
			}
		}
	}

	return false
}
