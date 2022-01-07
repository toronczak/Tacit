package main

import (
	"fmt"
	"os"
)

func player1Move() {
	drawBoard()
	if checkWin() != 0 {
		fmt.Printf("Player 2 wins!\n\n")
		hist[diff][2]++
		main()
		os.Exit(0)
	}
	if checkFull() == 1 {
		fmt.Printf("D  R  A  W  !\n\n")
		hist[diff][1]++
		main()
		os.Exit(0)
	}
	var x, y int
	fmt.Println("Player 1's (x, y) input: ")
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	for x > 2 || x < 0 || y > 2 || y < 0 || A[y][x] != ' ' {
		fmt.Println("INCORRECT MOVE!\nPlayer 1's (x, y) input: ")
		fmt.Scanln(&x)
		fmt.Scanln(&y)
	}
	A[y][x] = 'X'
	player2Move()
}

func player2Move() {
	drawBoard()
	if checkWin() != 0 {
		fmt.Printf("Player 1 wins!\n\n")
		hist[diff][0]++
		main()
		os.Exit(0)
	}
	if checkFull() == 1 {
		fmt.Printf("D  R  A  W  !\n\n")
		hist[diff][1]++
		main()
		os.Exit(0)
	}
	var x, y int
	fmt.Println("Player 2's (x, y) input: ")
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	for x > 2 || x < 0 || y > 2 || y < 0 || A[y][x] != ' ' {
		fmt.Println("INCORRECT MOVE!\nPlayer 2's (x, y) input: ")
		fmt.Scanln(&x)
		fmt.Scanln(&y)
	}
	A[y][x] = 'O'
	player1Move()
}

func vs2P() {
	initFill()
	player1Move()
}
