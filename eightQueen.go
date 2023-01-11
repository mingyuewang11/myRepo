package main

// import (
// 	"fmt"
// )

// var Queen [8]int
// var count = 0

// func queen() {
// 	eightQueen(0)

// }
// func eightQueen(row int) {
// 	if row == 8 {
// 		count++
// 		output(Queen)
// 		return
// 	}
// 	for column := 0; column < 8; column++ {
// 		if isSafe(row, column) {
// 			Queen[row] = column
// 			eightQueen(row + 1)
// 		}
// 	}
// }
// func output(nums [8]int) {
// 	for i := 0; i < 8; i++ {
// 		for j := 0; j < 8; j++ {
// 			if nums[i] == j {
// 				fmt.Print("Q")
// 			} else {
// 				fmt.Print("*")
// 			}

// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println(count)
// 	fmt.Println()
// }
// func isSafe(row, column int) bool {
// 	leftup := column - 1
// 	rightup := column + 1
// 	for i := row - 1; i >= 0; i-- {
// 		if Queen[i] == column {
// 			return false
// 		}
// 		if leftup >= 0 {
// 			if Queen[i] == leftup {
// 				return false
// 			}
// 		}
// 		if rightup < 8 {
// 			if Queen[i] == rightup {
// 				return false
// 			}
// 		}
// 		leftup--
// 		rightup++
// 	}
// 	return true
// }
