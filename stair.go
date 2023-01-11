package main

// import "fmt"

// var (
// 	steps [][]int
// )

// func stair() {
// 	fmt.Println(climb(5))
// 	fmt.Println(steps)
// }

// func climb(stair int) int {
// 	steps = make([][]int, stair)
// 	if stair == 1 {
// 		steps[stair] = []int{1}
// 		return 1
// 	} else if stair == 2 {
// 		steps = [][]int{{1, 1}, {2}}
// 		return 2
// 	} else {
// 		// steps[stair] = append(steps[stair], climb(stair-1))
// 		return climb(stair-1) + climb(stair-2)
// 	}
// }
