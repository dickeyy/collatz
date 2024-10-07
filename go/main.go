package main

import (
	"fmt"
	"time"
)

// Main algorithm for calculating the Collatz sequence
func calculateCollatz(n int, printSteps bool) int {
	fmt.Printf("\nCalculating Collatz sequence for %d...\n", n)
	if printSteps {
		fmt.Println("Steps:")
	}

	steps := 0

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			steps += 1
			if printSteps {
				fmt.Println(n)
			}
		} else {
			n = 3*n + 1
			steps += 1
			if printSteps {
				fmt.Println(n)
			}
		}
	}
	return steps
}

func runProgram(mode string) {
	showSteps := false
	fmt.Print("Do you want to see the steps? (y/n): ")

	var input string
	fmt.Scan(&input)
	if input == "y" {
		showSteps = true
	}

	if mode == "s" {
		n := 1
		for n != 0 {
			steps := calculateCollatz(n, showSteps)
			fmt.Printf("The number %d took %d steps to reach 1.\n", n, steps)
			n += 1
		}
	} else if mode == "c" {
		var n int
		fmt.Print("Enter a positive integer: ")
		fmt.Scan(&n)

		// Start timer
		startTime := time.Now()
		steps := calculateCollatz(n, showSteps)
		endTime := time.Now()

		// Calculate elapsed time
		elapsedTime := endTime.Sub(startTime).Seconds()

		fmt.Printf("\nThe number %d took %d steps to reach 1 in %.4f seconds.\n", n, steps, elapsedTime)
	} else {
		fmt.Println("Invalid input. Please enter 's' or 'c'.")
		main()
	}
}

func main() {
	mode := "s"
	fmt.Print("Enter 's' to sequentially calculate every number, or 'c' to calculate the sequence of a single number: ")
	fmt.Scan(&mode)
	runProgram(mode)
}
