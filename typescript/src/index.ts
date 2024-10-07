function calculateCollatz(n: number, printSteps: boolean = false): number {
    console.log(`\nCalculating Collatz sequence for ${n}...`);
    if (printSteps) {
        console.log("Steps:");
    }

    let steps = 0;

    while (n !== 1) {
        if (n % 2 === 0) {
            n = n / 2;
            steps += 1;
            if (printSteps) {
                console.log(n);
            }
        } else {
            n = 3 * n + 1;
            steps += 1;
            if (printSteps) {
                console.log(n);
            }
        }
    }
    return steps;
}

function runProgram(mode: string): void {
    const showSteps = prompt("Do you want to see the steps? (y/n): ");

    if (mode === "s") {
        let n = 1;
        while (n !== 0) {
            const steps = calculateCollatz(n, showSteps === "y");
            console.log(`The number ${n} took ${steps} steps to reach 1.`);
            n += 1;
        }
    } else if (mode === "c") {
        const n = parseInt(prompt("Enter a positive integer: ") || "0");

        // Start timer
        const startTime = Date.now();
        const steps = calculateCollatz(n, showSteps === "y");
        const endTime = Date.now();

        // Calculate elapsed time
        const elapsedTime = ((endTime - startTime) / 1000).toFixed(4);

        console.log(`\nThe number ${n} took ${steps} steps to reach 1 in ${elapsedTime} seconds.`);
    } else {
        console.log("Invalid input. Please enter 's' or 'c'.");
        main();
    }
}

function main(): void {
    const mode = prompt(
        "Enter 's' to sequentially calculate every number, or 'c' to calculate the sequence of a single number: "
    );
    runProgram(mode || "s");
}

main();
