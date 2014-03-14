// Prime Factorization
// Have the user enter a number and find all Prime Factors (if there are any)
// and display them.

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var sequence []int
	fmt.Print("Enter an integer: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i := 2; i < n/2; i++ {
		if isPrime(i) && isFactor(n, i) {
			sequence = append(sequence, i)
		}
	}
	if len(sequence) == 0 {
		fmt.Println("No Prime Factors found.")
	} else {
		fmt.Println(sequence)
	}
}

func isFactor(num, factor int) bool {
	if num%factor == 0 {
		return true
	} else {
		return false
	}
}

func isPrime(num int) bool {
	max := int(math.Ceil(math.Sqrt(float64(num))))
	for i := 2; i <= max; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
