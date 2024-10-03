# Collatz Conjecture Explorer

This repository contains a collection of programs designed to explore the fascinating Collatz
conjecture, also known as the 3n + 1 problem. It builds upon and extends the work from an earlier,
now archived repository [here](https://github.com/dickeyy/Collatz-Calculator).

## The Collatz Conjecture: An Enigmatic Mathematical Puzzle

The Collatz conjecture, named after mathematician
[Lothar Collatz](https://en.wikipedia.org/wiki/Lothar_Collatz), the most intriguing unsolved
problems in mathematics (in my opinion). It's deceptively simple to state, yet has resisted proof
for decades.

The conjecture works as follows:

1. Start with any positive integer n.
2. If n is even, divide it by 2.
3. If n is odd, multiply it by 3 and add 1.
4. Repeat steps 2 and 3 until you reach 1.

The conjecture posits that no matter what number you start with, you will always eventually reach 1,
entering the loop 4 → 2 → 1.

Mathematically, this can be expressed as:

```
f(n) = {
    n/2     if n is even
    3n + 1  if n is odd
}
```

While every number tested so far has been shown to reach the 4-2-1 loop, a rigorous mathematical
proof for all positive integers remains elusive.

## Project Overview

This repository houses a series of programs implemented in various programming languages. The
primary goals of this project are:

1. To provide educational resources for understanding the Collatz conjecture.
2. To offer efficient tools for computing Collatz sequences and analyzing their properties.
3. To serve as a platform for exploring algorithmic approaches to mathematical problems.

It's important to note that these programs are not aimed at proving the conjecture. Instead, they
allow users to experiment with the conjecture, visualize sequences, and potentially discover
interesting patterns or behaviors.

## Algorithms

[Collatz.py](python/src/collatz.py) is a Python program that implements the Collatz algorithm to
calculate the Collatz sequence for a given number. It also allows users to sequentially calculate
every number until they cancel.

[Collatz.ts](typescript/src/index.ts) is a TypeScript program that implements the Collatz algorithm
to calculate the Collatz sequence for a given number. It also allows users to sequentially calculate
every number until they cancel.

[Collatz.go](go/main.go) is a Go program that implements the Collatz algorithm to calculate the
Collatz sequence for a given number. It also allows users to sequentially calculate every number
until they cancel.
