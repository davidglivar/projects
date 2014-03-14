// Fibonacci Sequence
// Enter a number and have the program generate the Fibonacci sequence to that
// number or to the Nth number.

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter an integer: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			sequence[i] = 1
		} else {
			sequence[i] = sequence[i-1] + sequence[i-2]
		}
		fmt.Println(sequence[i])
		if sequence[i] == n {
			return
		}
	}
}
