package main

import (
	"fmt"
	"time"
)

// Main algorithm for calculating the Collatz sequence
func calculateCollatz(n int, printSteps bool) int {
	currentNum := n
	steps := 0

	if printSteps {
		fmt.Printf("\033[s") // Save cursor position
	}

	for currentNum != 1 {
		if currentNum%2 == 0 {
			currentNum = currentNum / 2
		} else {
			currentNum = 3*currentNum + 1
		}
		steps += 1
		if printSteps {
			fmt.Printf("\033[u\033[K") // Restore cursor position and clear line
			fmt.Printf("Current: %d, Steps: %d", currentNum, steps)
		}
	}
	if printSteps {
		fmt.Println() // Move to next line after completion
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
		fmt.Printf("\033[s") // Save initial cursor position
		for n != 0 {
			steps := calculateCollatz(n, showSteps)
			fmt.Printf("\033[u\033[K") // Restore cursor and clear line
			fmt.Printf("The number %d took %d steps to reach 1.", n, steps)
			n += 1
		}
		fmt.Println() // Final newline
	} else if mode == "c" {
		var n int
		fmt.Print("Enter a positive integer: ")
		fmt.Scan(&n)

		startTime := time.Now()
		steps := calculateCollatz(n, showSteps)
		endTime := time.Now()
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
