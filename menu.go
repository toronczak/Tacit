package main

import (
	"fmt"
)

var A [3][3]byte
var hist [102][3]int
var diff int
var flag int

func history() {
	fmt.Printf("Session data against this ")
	if diff == 101 {
		fmt.Printf("player opponent:\n")
	} else {
		fmt.Printf("level %d opponent:\n", diff)
	}
	fmt.Printf("wins - %d\nties - %d\nlosses - %d\n\n", hist[diff][0], hist[diff][1], hist[diff][2])
}

func initFill() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			A[i][j] = ' '
		}
	}
}

func drawBoard() {
	fmt.Println("-------------")
	for i := 0; i < 3; i++ {
		fmt.Printf("| %c | %c | %c |\n", A[i][0], A[i][1], A[i][2])
		fmt.Println("-------------")
	}
}

func checkFull() int {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if A[i][j] == ' ' {
				return 0
			}
		}
	}
	return 1
}

func checkWin() int {
	for i := 0; i < 3; i++ {
		if A[i][0] == 'X' && A[i][1] == 'X' && A[i][2] == 'X' ||
			A[0][i] == 'X' && A[1][i] == 'X' && A[2][i] == 'X' {
			return 1
		}
	}
	if A[0][0] == 'X' && A[1][1] == 'X' && A[2][2] == 'X' ||
		A[2][0] == 'X' && A[1][1] == 'X' && A[0][2] == 'X' {
		return 1
	}

	for i := 0; i < 3; i++ {
		if A[i][0] == 'O' && A[i][1] == 'O' && A[i][2] == 'O' ||
			A[0][i] == 'O' && A[1][i] == 'O' && A[2][i] == 'O' {
			return -1
		}
	}
	if A[0][0] == 'O' && A[1][1] == 'O' && A[2][2] == 'O' ||
		A[2][0] == 'O' && A[1][1] == 'O' && A[0][2] == 'O' {
		return -1
	}
	return 0
}

func main() {
	if flag == 0 {
		for i := 0; i < 102; i++ {
			for j := 0; j < 3; j++ {
				hist[i][j] = 0
			}
		}
	} else {
		history()
	}
	flag = 1
	fmt.Println("      made by Toronczak")
	fmt.Println("[][][][]  [][][][]  [][][][]")
	fmt.Println("   []        []     []")
	fmt.Println("   []        []     []")
	fmt.Println("   []        []     []")
	fmt.Println("   []     [][][][]  [][][][]")
	fmt.Println("")
	fmt.Println("[][][][]  [][][][]  [][][][]")
	fmt.Println("   []     []    []  []")
	fmt.Println("   []     [][][][]  []")
	fmt.Println("   []     []    []  []")
	fmt.Println("   []     []    []  [][][][]")
	fmt.Println("1 - play vs. AI")
	fmt.Println("2 - local multiplayer")
	fmt.Println("anything else to exit")

	var n int8
	fmt.Scanln(&n)

	if n == 1 {
		fmt.Println("Enter difficulty level, from (random) 0 to (Minimax) 100:")
		fmt.Scanln(&diff)
		if diff < 0 {
			diff = 0
			fmt.Println("Level corrected to 0.")
		} else if diff > 100 {
			diff = 100
			fmt.Println("Level corrected to 100.")
		}
		vsCPU()
	} else if n == 2 {
		diff = 101
		vs2P()
	}
}
