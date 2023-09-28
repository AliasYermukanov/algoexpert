package main

import (
	"fmt"
	"log"
	"time"
)

func GetNthFib(n int) int {
	if n == 1 {
		return 0
	}
	n1, n2 := 0, 1
	for i := 3; i <= n; i++ {
		n1, n2 = n2, n2+n1
	}
	return n2
}

func main() {
	start := time.Now()
	fmt.Println(GetNthFib(6))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
