package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	start := intervals[0][0]
	end := intervals[0][1]
	result := make([][]int, 0, len(intervals))
	for _, interval := range intervals {

		// 重叠
		if interval[0] <= end {
			if end < interval[1] {
				end = interval[1]
			}
		} else {
			result = append(result, []int{start, end})
			start = interval[0]
			end = interval[1]
		}
	}
	result = append(result, []int{start, end})
	return result
}

func main() {
	inputs := [][][]int{
		{{1, 4}, {0, 4}},
		{{1, 4}, {4, 5}},
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
	}
	for _, input := range inputs {
		fmt.Printf("%v after merge: %v \n", input, merge(input))
	}
}
