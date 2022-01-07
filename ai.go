package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func playerMove() {
	drawBoard()
	if checkWin() != 0 {
		fmt.Printf("The level %d computer wins!\n\n", diff)
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
	fmt.Println("Your (x, y) input: ")
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	for x > 2 || x < 0 || y > 2 || y < 0 || A[y][x] != ' ' {
		fmt.Println("INCORRECT MOVE!\nYour (x, y) input: ")
		fmt.Scanln(&x)
		fmt.Scanln(&y)
	}
	A[y][x] = 'X'
	CPUMove()
}

func CPUMove() {
	drawBoard()
	if checkWin() != 0 {
		hist[diff][0]++
		fmt.Printf("You win!\n\n")
		main()
		os.Exit(0)
	}
	if checkFull() == 1 {
		hist[diff][1]++
		fmt.Printf("D  R  A  W  !\n\n")
		main()
		os.Exit(0)
	}
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(3) > diff {
		weak()
	} else {
		strong()
	}
	playerMove()
}

func weak() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(3)
	y := rand.Intn(3)
	if A[x][y] == ' ' {
		A[x][y] = 'O'
	} else {
		weak()
	}
}

func strong() {
	var x, y int
	score := 1

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if A[i][j] == ' ' {
				A[i][j] = 'O'
				temp := searchMax()
				if score > temp {
					x = i
					y = j
					score = temp
				}
				A[i][j] = ' '
			}
		}
	}
	A[x][y] = 'O'
}

func searchMax() int {
	ret := -1
	if checkWin() != 0 {
		return checkWin()
	}

	if checkFull() == 1 {
		return 0
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if A[i][j] == ' ' {
				A[i][j] = 'X'
				ret = max(ret, searchMin())
				A[i][j] = ' '
			}
		}
	}
	return ret
}

func searchMin() int {
	ret := 1
	if checkWin() != 0 {
		return checkWin()
	}

	if checkFull() == 1 {
		return 0
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if A[i][j] == ' ' {
				A[i][j] = 'O'
				ret = min(ret, searchMax())
				A[i][j] = ' '
			}
		}
	}
	return ret
}

func vsCPU() {
	initFill()
	playerMove()
}
