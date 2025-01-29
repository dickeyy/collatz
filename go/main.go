package main

import (
	"fmt"
	"time"
)

// ANSI color codes
const (
	Gray  = "\033[38;5;242m"
	Pink  = "\033[38;5;205m"
	Green = "\033[38;5;046m"
	Reset = "\033[0m"
)

// Main algorithm for calculating the Collatz sequence
func calculateCollatz(n int, printSteps bool) int {
	if printSteps {
		fmt.Printf("\n%sCalculating Collatz sequence for %s%d%s...\n", Gray, Pink, n, Reset)
		fmt.Println(Gray + "Steps:" + Reset)
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
			fmt.Printf("%s%d%s\n", Gray, currentNum, Reset)
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
		startTime := time.Now()

		for n != 0 {
			steps := calculateCollatz(n, showSteps)
			currentRate := float64(n) / time.Since(startTime).Seconds()
			fmt.Printf("%s%d%s took %s%d%s steps to reach 1. (%s%.2f %sMil/s)%s\n",
				Pink, n, Gray, Pink, steps, Gray, Green, currentRate/1000000, Gray, Reset)
			n++
		}
	} else if mode == "c" {
		var n int
		fmt.Print("Enter a positive integer: ")
		fmt.Scan(&n)
		startTime := time.Now()
		steps := calculateCollatz(n, showSteps)
		rate := 1.0 / time.Since(startTime).Seconds()
		fmt.Printf("%s%d%s took %s%d%s steps to reach 1. (%s%.2f %sMil/s)%s\n",
			Pink, n, Gray, Pink, steps, Gray, Green, rate/1000000, Gray, Reset)
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
