package main

import "testing"

func TestSlice(t *testing.T) {
	k := 1
	nums := make([][]int, 9)
	for i := 0; i < 9; i++ {
		nums[i] = make([]int, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			nums[i][j] = k
			k++
		}
	}
	t.Log(nums)
	row := 4
	column := 5
	temp := make([][]int, 3)
	// k:=row/3
	for i := 0; i < 3; i++ {
		temp[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			temp[i][j] = nums[row/3*3+i][column/3*3+j]
		}
		// copy(temp, nums[row/3*3 : row/3*3+3][column/3*3:column/3*3])
	}
	t.Log(temp)
	var a []int
	for _, value := range temp {
		a = append(a, value[0])
	}

	t.Log(a)
}

func TestNums(t *testing.T) {
	temp := []int{3, 0, 8, 9, 4, 6}
	var nums []int
	nums = base[0:]
	newNums := make([]int, 0)
	for _, v := range temp {
		repeat := false
		if len(newNums) > 0 {
			for _, value := range newNums {
				if value == v {
					repeat = true
					break
				}
			}
		}
		if repeat {
			continue
		}
		newNums = append(newNums, v)
	}
	t.Log(newNums)
	for key, value := range newNums {
		t.Log(key)
		for k, v := range nums {
			if value == v {
				t.Log(k, v)
				// if k == len(nums)+1 {
				// 	nums = nums[:k]
				// } else {
				nums = append(nums[:k], nums[k+1:]...)
				// }

			}

		}
		t.Log(nums)
		// nums = append(nums, value)
	}
	t.Log(nums)
}
func TestBox(t *testing.T) {
	soduku = [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9}}
	t.Log(getBox(1, 1))
}

func TestQueen(t *testing.T) {
	bai()
	// queen()
}
func TestStair(t *testing.T) {
	// stair()
	// var Board = [][]int{
	// 	0: {0, 9, 0, 1, 0, 0, 5, 0, 0},
	// 	1: {0, 0, 0, 0, 7, 9, 8, 0, 1},
	// 	2: {2, 0, 0, 0, 0, 0, 0, 0, 6},
	// 	3: {0, 3, 0, 0, 0, 0, 0, 0, 0},
	// 	4: {0, 0, 0, 7, 8, 1, 0, 0, 0},
	// 	5: {0, 0, 4, 0, 0, 0, 0, 2, 0},
	// 	6: {7, 0, 0, 6, 0, 0, 0, 0, 4},
	// 	7: {6, 0, 1, 5, 3, 0, 0, 0, 0},
	// 	8: {0, 0, 9, 0, 0, 7, 0, 6, 0},
	// }
	// Board[5][6] = 3
	// t.Log(isSafe(Board, 5, 6, 3))
}
