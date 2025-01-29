package main

import (
	"fmt"
	"time"
)

// Main algorithm for calculating the Collatz sequence
func calculateCollatz(n int, printSteps bool) int {
	if printSteps {
		fmt.Printf("\nCalculating Collatz sequence for %d...\n", n)
		fmt.Println("Steps:")
	}

	steps := 0
	currentNum := n

	for currentNum != 1 {
		if currentNum%2 == 0 {
			currentNum = currentNum / 2
		} else {
			currentNum = 3*currentNum + 1
		}
		steps++
		if printSteps {
			fmt.Println(currentNum)
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
		// Ask user for report frequency
		fmt.Print("Enter how often you want status updates (int > 0; every n calculations): ")
		var reportFrequency int
		fmt.Scan(&reportFrequency)

		if reportFrequency <= 0 {
			fmt.Println("Report frequency must be greater than 0")
			return
		}

		n := 1
		for n != 0 {
			steps := calculateCollatz(n, showSteps && (n%reportFrequency == 0))
			if n%reportFrequency == 0 {
				fmt.Printf("The number %d took %d steps to reach 1.\n", n, steps)
			}
			n++
		}
	} else if mode == "c" {
		var n int
		fmt.Print("Enter a positive integer: ")
		fmt.Scan(&n)

		startTime := time.Now()
		steps := calculateCollatz(n, showSteps)
		elapsedTime := time.Now().Sub(startTime).Seconds()

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
