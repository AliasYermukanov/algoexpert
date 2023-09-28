package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(BinarySearch([]int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}, 33))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func BinarySearch(array []int, target int) int {
	return find(array, target, 0, len(array)-1)
}

func find(array []int, target, start, end int) int {
	middle := (start + end) / 2
	if array[middle] == target {
		return middle
	} else if start >= end {
		return -1
	} else if target < array[middle] {
		return find(array, target, start, middle-1)
	} else if target > array[middle] {
		return find(array, target, middle+1, end)
	}
	return -1
}
