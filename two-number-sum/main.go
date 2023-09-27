package main

import (
	"fmt"
	"log"
	"time"
)

func TwoNumberSum(array []int, target int) []int {
	mapa := make(map[int]int)

	for i := 0; i < len(array); i++ {
		needed := target - array[i]
		if _, ok := mapa[needed]; ok {
			return []int{array[i], array[mapa[needed]]}
		} else {
			mapa[array[i]] = i
		}
	}

	return []int{}
}

func main() {
	start := time.Now()
	fmt.Println(TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
