# Collatz.py

This program is written in Python and is designed to either calculate the Collatz sequence for a
given number, or sequentially calculate every number (until you cancel).

## How to Use

To use the program, simply run the `src/collatz.py` file. The program will prompt you to enter a
mode. Enter `s` to sequentially calculate every number until you cancel, or enter `c` to calculate
the Collatz sequence for a single number.

## Example Usage

### Sequentially Calculating Every Number

```
$ python src/collatz.py
Enter 's' to sequentially calculate every number, or 'c' to calculate the sequence of a single number: s

Calculating Collatz sequence for 1...
...
Calculating Collatz sequence for 100...
...
```

### Calculating the Collatz Sequence for a Single Number

```
$ python src/collatz.py
Enter 's' to sequentially calculate every number, or 'c' to calculate the sequence of a single number: c
Enter a positive integer: 100
Do you want to see the steps? (y/n): y

The number 100 took 10 steps to reach 1.
```

## Contributing

Contributions are welcome! Open an issue!
