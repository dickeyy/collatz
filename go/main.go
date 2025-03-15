package main

import (
	"flag"
	"fmt"
	"strconv"
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
func calculateCollatz(n int64, printSteps bool) int64 {
	if printSteps {
		fmt.Printf("\n%sCalculating Collatz sequence for %s%d%s...\n", Gray, Pink, n, Reset)
		fmt.Println(Gray + "Steps:" + Reset)
	}
	var steps int64 = 0
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
		var n int64
		fmt.Println("Would you like to start at 1 or a specific number? (1 or n): ")
		var input string
		fmt.Scan(&input)
		if input == "n" {
			fmt.Print("Enter a positive integer: ")
			fmt.Scan(&input)
			start, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a positive integer.")
				return
			}
			n = start
		} else {
			n = 1
		}
		startTime := time.Now()

		for n != 0 {
			steps := calculateCollatz(n, showSteps)
			currentRate := float64(n) / time.Since(startTime).Seconds()
			fmt.Printf("%s%d%s took %s%d%s steps to reach 1. (%s%.2f %sMil/s)%s\n",
				Pink, n, Gray, Pink, steps, Gray, Green, currentRate/1000000, Gray, Reset)
			n++
		}
	} else if mode == "c" {
		var n int64
		fmt.Print("Enter a positive integer: ")
		var inputStr string
		fmt.Scan(&inputStr)
		parsedN, err := strconv.ParseInt(inputStr, 10, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a positive integer.")
			return
		}
		n = parsedN
		startTime := time.Now()
		steps := calculateCollatz(n, showSteps)
		rate := 1.0 / time.Since(startTime).Seconds()
		fmt.Printf("%s%d%s took %s%d%s steps to reach 1. (%s%.2f %sMil/s)%s\n",
			Pink, n, Gray, Pink, steps, Gray, Green, rate/1000000, Gray, Reset)
	} else {
		fmt.Println("Invalid input. Please enter 's' or 'c'.")
		promptForMode()
	}
}

// Helper function to prompt for mode and run the program
func promptForMode() {
	mode := "s"
	fmt.Print("Enter 's' to sequentially calculate every number, or 'c' to calculate the sequence of a single number: ")
	fmt.Scan(&mode)
	runProgram(mode)
}

func runProgramDefault() {
	// here just default to sequential mode, no steps, start at 1
	startTime := time.Now()
	var n int64 = 1
	for n != 0 {
		steps := calculateCollatz(n, false)
		currentRate := float64(n) / time.Since(startTime).Seconds()
		fmt.Printf("%s%d%s took %s%d%s steps to reach 1. (%s%.2f %sMil/s)%s\n",
			Pink, n, Gray, Pink, steps, Gray, Green, currentRate/1000000, Gray, Reset)
		n++
	}
}

func main() {
	// Define the -d flag
	defaultMode := flag.Bool("d", false, "Run in default mode without user input")

	// Parse the input flags
	flag.Parse()

	// Check if the -d flag is set
	if *defaultMode {
		runProgramDefault()
	} else {
		promptForMode()
	}
}
