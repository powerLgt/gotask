package main

import "fmt"

func findOnlyOne(arr []int) {
	intMap := make(map[int]int)
	for i := range arr {
		intMap[arr[i]] += 1
	}

	for k, v := range intMap {
		if v == 1 {
			fmt.Println(k)
		}
	}
}

func main() {
	findOnlyOne([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}
