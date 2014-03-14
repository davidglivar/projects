// Find PI to the Nth Digit
// Enter a number and have the program generate PI up to that many decimal
// places. Keep a limit to how far the program will go.

package main

import (
	"fmt"
	"math"
)

func main() {
	var place int
	fmt.Print("How many decimals? ")
	_, err := fmt.Scan(&place)
	if err != nil {
		fmt.Println(err)
		return
	}
	if place > 50 {
		fmt.Println("Try a smaller number")
		return
	}
	pi := (4 * (4*math.Atan(1.0/5.0) - math.Atan(1.0/239.0)))
	fmt.Printf("%."+fmt.Sprintf("%d", place)+"f\n", pi)
}
