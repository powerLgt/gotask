package main

import "fmt"

/*
	给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
	这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
	将大整数加 1，并返回结果的数字数组。
*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] <= 9 {
			return digits
		}
		digits[i] = 0
	}

	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

func main() {
	inputs := [][]int{
		{1, 2, 3, 4, 5},
		{0},
		{9},
		{1, 0},
		{9, 9},
		{1, 2, 3},
		{4, 3, 2, 1},
	}
	for _, input := range inputs {
		fmt.Printf("%v", input)
		fmt.Printf(" plusOne: %v \n", plusOne(input))
	}
}
