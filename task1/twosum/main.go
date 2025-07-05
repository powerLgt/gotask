package main

import "fmt"

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{0, 0}
}

func main() {
	inputs := [][][]int{
		{{2, 7, 11, 15}, {9}},
		{{3, 2, 4}, {6}},
		{{3, 3}, {6}},
	}
	for _, input := range inputs {
		fmt.Printf("%v twoSum: %v \n", twoSum(input[0], input[1][0]), input[1][0])
	}
}
