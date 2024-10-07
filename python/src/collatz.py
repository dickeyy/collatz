import time

# Main algorithm for calculating the Collatz sequence
def calculate_collatz(n, print_steps=False):
    print(f"\nCalculating Collatz sequence for {n}...")
    if (print_steps):
        print("Steps:")

    steps = 0 

    while n != 1:
        if n % 2 == 0:
            n = n // 2
            steps += 1
            if print_steps:
                print(n)
        else:
            n = 3 * n + 1
            steps += 1
            if print_steps:
                print(n)
    return steps

# Helper function for running the program
def run_program(mode):
    show_steps = input("Do you want to see the steps? (y/n): ")

    if mode == "s":
        n = 1
        while n != 0:
            steps = calculate_collatz(n, print_steps=show_steps == "y")
            print(f"The number {n} took {steps} steps to reach 1.")
            n += 1

    elif mode == "c":
        n = int(input("Enter a positive integer: "))

        # Start timer
        start_time = time.time()
        steps = calculate_collatz(n, print_steps=show_steps == "y")
        end_time = time.time()

        # Calculate elapsed time
        elapsed_time = round(end_time - start_time, 4)

        print(f"\nThe number {n} took {steps} steps to reach 1 in {elapsed_time} seconds.")

    else:
        print("Invalid input. Please enter 's' or 'c'.")
        exit(1)

def main():
    mode = input("Enter 's' to sequentially calculate every number, or 'c' to calculate the sequence of a single number: ")
    run_program(mode)

if __name__ == "__main__":
    main()