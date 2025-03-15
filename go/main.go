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
func calculateCollatz(n uint, printSteps bool) uint {
	if printSteps {
		fmt.Printf("\n%sCalculating Collatz sequence for %s%d%s...\n", Gray, Pink, n, Reset)
		fmt.Println(Gray + "Steps:" + Reset)
	}
	var steps uint = 0
	currentNum := n
	for currentNum != 1 {
		if (currentNum & 1) == 0 {
			currentNum = currentNum >> 1
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

func runProgram(mode string, printFreq uint) {
	showSteps := false
	fmt.Print("Do you want to see the steps? (y/n): ")
	var input string
	fmt.Scan(&input)
	if input == "y" {
		showSteps = true
	}

	// If printFreq wasn't set via flag, ask the user
	if printFreq == 0 {
		fmt.Print("Enter print frequency (how often to display results, e.g., 1 for every number, 10 for every 10th number): ")
		var freqStr string
		fmt.Scan(&freqStr)
		freq, err := strconv.ParseUint(freqStr, 10, 64)
		if err != nil || freq == 0 {
			fmt.Println("Invalid frequency. Using default of 1.")
			printFreq = 1
		} else {
			printFreq = uint(freq)
		}
	}

	if mode == "s" {
		var n uint
		fmt.Println("Would you like to start at 1 or a specific number? (1 or n): ")
		var input string
		fmt.Scan(&input)
		if input == "n" {
			fmt.Print("Enter a positive integer: ")
			fmt.Scan(&input)
			start, err := strconv.ParseUint(input, 10, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a positive integer.")
				main()
			}
			n = uint(start)
		} else {
			n = 1
		}
		startTime := time.Now()

		for n != 0 {
			steps := calculateCollatz(n, showSteps)

			// Only print if n is a multiple of printFreq or if n is 1 (to always show the first number)
			if n == 1 || n%printFreq == 0 {
				currentRate := float64(n) / time.Since(startTime).Seconds()
				fmt.Printf("%s%d%s took %s%d%s steps to reach 1. (%s%.2f %sMil/s)%s\n",
					Pink, n, Gray, Pink, steps, Gray, Green, currentRate/1000000, Gray, Reset)
			}
			n++
		}
	} else if mode == "c" {
		var n uint
		fmt.Print("Enter a positive integer: ")
		var inputStr string
		fmt.Scan(&inputStr)
		parsedN, err := strconv.ParseUint(inputStr, 10, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a positive integer.")
			main()
			return
		}
		n = uint(parsedN)

		startTime := time.Now()
		steps := calculateCollatz(n, showSteps)
		rate := float64(n) / time.Since(startTime).Seconds()
		fmt.Printf("%s%d%s took %s%d%s steps to reach 1. (%s%.2f %sMil/s)%s\n",
			Pink, n, Gray, Pink, steps, Gray, Green, rate/1000000, Gray, Reset)
	} else {
		fmt.Println("Invalid input. Please enter 's' or 'c'.")
		main()
	}
}

func runProgramDefault(printFreq uint) {
	// Default to sequential mode, no steps, start at 1
	startTime := time.Now()
	var n uint = 1
	for n != 0 {
		steps := calculateCollatz(n, false)

		// Only print if n is a multiple of printFreq or if n is 1 (to always show the first number)
		if n == 1 || n%printFreq == 0 {
			currentRate := float64(n) / time.Since(startTime).Seconds()
			fmt.Printf("%s%d%s took %s%d%s steps to reach 1. (%s%.2f %sMil/s)%s\n",
				Pink, n, Gray, Pink, steps, Gray, Green, currentRate/1000000, Gray, Reset)
		}
		n++
	}
}

func main() {
	// Define the flags
	defaultMode := flag.Bool("d", false, "Run in default mode without user input")
	printFreq := flag.Uint("f", 0, "Print frequency - only print every N numbers (default: 1)")

	// Parse the input flags
	flag.Parse()

	// If printFreq is 0, set it to 1 when using default mode
	if *defaultMode && *printFreq == 0 {
		*printFreq = 1
	}

	// Check if the -d flag is set
	if *defaultMode {
		runProgramDefault(*printFreq)
	} else {
		mode := "s"
		fmt.Print("Enter 's' to sequentially calculate every number, or 'c' to calculate the sequence of a single number: ")
		fmt.Scan(&mode)
		runProgram(mode, *printFreq)
	}
}
