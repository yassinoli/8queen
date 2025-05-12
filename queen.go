package main

import "fmt"

const N = 8

var board [N]int

// Check if it's safe to place queen at (row, col)
func isSafe(row, col int) bool {
    for i := 0; i < row; i++ {
        if board[i] == col ||               // Same column
            board[i]-i == col-row ||        // Same \ diagonal
            board[i]+i == col+row {         // Same / diagonal
            return false
        }
    }
    return true
}

// Solve the puzzle with backtracking
func solve(row int) {
    if row == N {
        printBoard()
        return
    }

    for col := 0; col < N; col++ {
        if isSafe(row, col) {
            board[row] = col
            solve(row + 1)
        }
    }
}

// Print solution in format like: 15863724
func printBoard() {
    for i := 0; i < N; i++ {
        fmt.Print(board[i] + 1) // +1 to convert 0-based to 1-based
    }
    fmt.Println()
}

func main() {
    solve(0)
}
