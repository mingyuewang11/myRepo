package main

import (
	"fmt"
	"testing"
)

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

func test() {
	fmt.Println("11111111111111")
}
