package main

import "fmt"

var blockStage = 3
var total = 0

func bai() {
	var Board = [9][9]int{
		0: {0, 9, 0, 1, 0, 0, 5, 0, 0},
		1: {0, 0, 0, 0, 7, 9, 8, 0, 1},
		2: {2, 0, 0, 0, 0, 0, 0, 0, 6},
		3: {0, 3, 0, 0, 0, 0, 0, 0, 0},
		4: {0, 0, 0, 7, 8, 1, 0, 0, 0},
		5: {0, 0, 4, 0, 0, 0, 0, 2, 0},
		6: {7, 0, 0, 6, 0, 0, 0, 0, 4},
		7: {6, 0, 1, 5, 3, 0, 0, 0, 0},
		8: {0, 0, 9, 0, 0, 7, 0, 6, 0},
	}
	goDown(0, Board)
	fmt.Println(total)
}

func goDown(step int, board [9][9]int) {
	if step == 9*blockStage*blockStage {
		fmt.Println("===================")
		for _, v := range board {
			fmt.Println(v)
		}
		total++
		return
	}
	x := step / (3 * blockStage)
	y := step % (3 * blockStage)
	if board[x][y] != 0 {
		goDown(step+1, board)
	} else {
		for i := 1; i <= 9; i++ {
			board[x][y] = i
			if isSafe(board, x, y, i) {
				goDown(step+1, board)
				board[x][y] = 0
			}
		}
	}
}

func isSafe(board [9][9]int, x, y, num int) bool {
	for i := 0; i < 3*blockStage; i++ {
		if board[x][i] == num && i != y {
			return false
		}
		if board[i][y] == num && i != x {
			return false
		}
	}
	sx := (x / 3) * 3
	sy := (y / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[sx+i][sy+j] == num && sx+i != x && sy+j != y {
				return false
			}
		}
	}
	return true
}

// var total = 0

// func hh() {
// 	q := make([]int, 8)
// 	puthh(q, 0)
// 	fmt.Println(total)
// }

// func puthh(q []int, x int) {
// 	if x == 8 {
// 		fmt.Println(q)
// 		total++
// 		return
// 	}
// 	for y := 1; y < 9; y++ {
// 		q[x] = y
// 		if isSafe(q, x, y) {
// 			puthh(q, x+1)
// 		}
// 	}
// }

// func isSafe(q []int, x, y int) bool {
// 	for k := 0; k < x; k++ {
// 		// fmt.Println(q[k], y)
// 		if q[k] == y {
// 			return false
// 		}
// 		if math.Abs(float64(x)-float64(k)) == math.Abs(float64(y)-float64(q[k])) {
// 			return false
// 		}
// 	}
// 	return true
// }
