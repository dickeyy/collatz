package main

import (
	"flag"
	"fmt"
	"math/big"
	"time"
)

// ANSI color codes
const (
	Gray  = "\033[38;5;242m"
	Pink  = "\033[38;5;205m"
	Green = "\033[38;5;046m"
	Reset = "\033[0m"
)

var (
	// Define constants as big.Int values
	zero  = big.NewInt(0)
	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
)

// Main algorithm for calculating the Collatz sequence
func calculateCollatz(n *big.Int, printSteps bool) *big.Int {
	if printSteps {
		fmt.Printf("\n%sCalculating Collatz sequence for %s%v%s...\n", Gray, Pink, n, Reset)
		fmt.Println(Gray + "Steps:" + Reset)
	}

	steps := big.NewInt(0)
	currentNum := new(big.Int).Set(n)

	for currentNum.Cmp(one) != 0 {
		// Check if even: currentNum % 2 == 0
		remainder := new(big.Int)
		currentNum.DivMod(currentNum, two, remainder)

		if remainder.Cmp(zero) == 0 {
			// Already divided by 2 in the DivMod operation
		} else {
			// 3*currentNum + 1
			currentNum.Mul(currentNum, three)
			currentNum.Add(currentNum, one)
		}

		steps.Add(steps, one)

		if printSteps {
			fmt.Printf("%s%v%s\n", Gray, currentNum, Reset)
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
		fmt.Println("Would you like to start at 1 or a specific number? (1 or n): ")
		var input string
		fmt.Scan(&input)

		n := big.NewInt(1)

		if input == "n" {
			fmt.Print("Enter a positive integer: ")
			fmt.Scan(&input)

			// Try to parse the input as a big.Int
			_, success := n.SetString(input, 10)
			if !success || n.Cmp(zero) <= 0 {
				fmt.Println("Invalid input. Please enter a positive integer.")
				return
			}
		}

		startTime := time.Now()
		count := big.NewInt(0)

		for {
			if n.Cmp(zero) == 0 {
				break
			}

			steps := calculateCollatz(n, showSteps)
			count.Add(count, one)

			// Use floating point for rate calculation
			nFloat, _ := new(big.Float).SetInt(n).Float64()
			currentRate := nFloat / time.Since(startTime).Seconds()

			fmt.Printf("%s%v%s took %s%v%s steps to reach 1. (%s%.2f %sMil/s)%s\n",
				Pink, n, Gray, Pink, steps, Gray, Green, currentRate/1000000, Gray, Reset)

			n.Add(n, one)
		}
	} else if mode == "c" {
		fmt.Print("Enter a positive integer: ")
		var inputStr string
		fmt.Scan(&inputStr)

		n := new(big.Int)
		_, success := n.SetString(inputStr, 10)

		if !success || n.Cmp(zero) <= 0 {
			fmt.Println("Invalid input. Please enter a positive integer.")
			return
		}

		startTime := time.Now()
		steps := calculateCollatz(n, showSteps)
		elapsed := time.Since(startTime).Seconds()

		// For large numbers, calculation of rate might not be meaningful
		rate := 1.0 / elapsed

		fmt.Printf("%s%v%s took %s%v%s steps to reach 1. (%s%.2f %sMil/s)%s\n",
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
	n := big.NewInt(1)
	count := big.NewInt(0)

	for {
		if n.Cmp(zero) == 0 {
			break
		}

		steps := calculateCollatz(n, false)
		count.Add(count, one)

		// Use floating point for rate calculation
		nFloat, _ := new(big.Float).SetInt(n).Float64()
		currentRate := nFloat / time.Since(startTime).Seconds()

		fmt.Printf("%s%v%s took %s%v%s steps to reach 1. (%s%.2f %sMil/s)%s\n",
			Pink, n, Gray, Pink, steps, Gray, Green, currentRate/1000000, Gray, Reset)

		n.Add(n, one)
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
