# Hit & Blow Solver
This is a command-line tool that solves the Hit & Blow game.   
It is a heuristic solver and does not provide a complete analytical solution.   
The solving method is based on the approach by B. Landy described on page 3 of the following paper (in Japanese):
https://www.tanaka.ecc.u-tokyo.ac.jp/ktanaka/papers/gpw96.pdf

## Installation
To install the CLI tool, run the following command:
```bash
go install https://github.com/ponyo877/hit-and-blow-solver
```

# Usage
The tool assumes a 3-digit Hit & Blow game using digits from 0-9.    
If you want to change the number of digits or the range of digits, you can edit the beginning of the `main` method in `main.go` directly.
```go
// main.go: main
disit := 3
numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
```

The tool will prompt you through the following steps:
1. [Automatic] A suitable guess sequence will be displayed.
2. [You] Submit the guess sequence from step 1 to your opponent.
3. [Automatic] A list of possible feedback options for the guess sequence will be displayed.
4. [You] Select the actual feedback from your opponent.

# Demo

Insert a GIF demo here.