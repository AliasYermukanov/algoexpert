package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(IsPalindrome("abcdcba"))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func IsPalindrome(str string) bool {
	if len(str) <= 1 {
		return true
	}

	for i, j := 0, len(str)-1; i <= j; {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}
