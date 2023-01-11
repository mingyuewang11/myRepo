package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	base   = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	board  [][]int
	soduku = [9][9]int{}
	// {0, 9, 0, 1, 0, 0, 5, 0, 0},
	// {0, 0, 0, 0, 7, 9, 8, 0, 1},
	// {2, 0, 0, 0, 0, 0, 0, 0, 6},
	// {0, 3, 0, 0, 0, 0, 0, 0, 0},
	// {0, 0, 0, 7, 8, 1, 0, 0, 0},
	// {0, 0, 4, 0, 0, 0, 0, 2, 0},
	// {7, 0, 0, 6, 0, 0, 0, 0, 4},
	// {6, 0, 1, 5, 3, 0, 0, 0, 0},
	// {0, 0, 9, 0, 0, 7, 0, 6, 0}}
	test []int
	c    = 0
	cou  = 0
)

func main() {
	for {
		board = make([][]int, 0)
		fmt.Println("请按行输入9*9的数独（空格用0代替，数字之间用空格分开）：")
		for i := 0; i < 9; i++ {
			reader := bufio.NewReader(os.Stdin)
			strBytes, _, _ := reader.ReadLine()
			str := strings.Fields(string(strBytes))
			var arr []int
			for j := range str {
				n, _ := strconv.Atoi(str[j])
				arr = append(arr, n)
			}
			if len(arr) != 9 {
				fmt.Println("格式错误，重新输入")
				break
			} else {
				board = append(board, arr)
			}
		}
		if len(board) == 9 {
			break
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			soduku[i][j] = board[i][j]
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if isOK(i, j, soduku[i][j]) || soduku[i][j] == 0 {
				continue
			} else {
				fmt.Println("有问题，重新输入")
				return
			}

		}
	}
	backtracking(0)
}

func backtracking(c int) {
	if c == 81 {
		cou++
		outPut(soduku)
		return
		// c = 0
	}
	row := c / 9
	column := c % 9
	// base = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if soduku[row][column] != 0 {
		backtracking(c + 1)
	} else {
		for i := 1; i <= 9; i++ {
			if isOK(row, column, i) {
				soduku[row][column] = i
				backtracking(c + 1)
				// if backtracking(c + 1) {
				// 	i++
				// 	// return true
				// } else {
				soduku[row][column] = 0
				// }

			}
		}
		// outPut(soduku)
		// fmt.Println()
		return //false
	}
}

func getArr(row, column int) (nums []int) {
	base = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	nums = make([]int, 0)
	//行，列，宫的全部元素
	temp := getList(getBox(row, column), row, column)
	nums = base[0:] //可填数字
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
	for _, value := range newNums {
		for k, v := range nums {
			if value == 0 {
				continue
			} else if value == v {
				nums = append(nums[:k], nums[k+1:]...)
			}
		}
	}

	return nums
}

func isOK(row, column, num int) bool {
	for k, v := range soduku[row] {
		if num == v && k != column {
			return false
		}
	}
	for _, v := range getBox(row, column) {
		if num == v && num != soduku[row][column] {
			return false
		}
	}
	for k := range base {
		if soduku[k][column] == num && k != row {
			return false
		}
	}
	return true
}
func outPut(nums [9][9]int) {
	fmt.Println("=====================")
	for i := 0; i < 9; i++ {
		fmt.Print("| ")
		for j := 0; j < 9; j++ {
			fmt.Print(nums[i][j], " ")
		}
		fmt.Println("|")
	}
	fmt.Println(cou)
	fmt.Println()
}

func getBox(row, column int) []int {
	temp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		temp[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			temp[i][j] = soduku[row/3*3+i][column/3*3+j]
		}
	}
	var t []int
	for _, value := range temp {
		t = append(t, value...)
	}
	return t
}

func getList(box []int, row, column int) []int {
	var list []int
	list = append(list, box...)
	for _, value := range soduku[row] {
		list = append(list, value)
	}
	for _, value := range soduku {
		list = append(list, value[column])
	}
	return list
}
