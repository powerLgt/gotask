package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	common := make([]byte, 0)

	for longestIndex := range strs[0] {
		word := strs[0][longestIndex]
		for i := 1; i < len(strs); i++ {

			if longestIndex >= len(strs[i]) || strs[i][longestIndex] != word {
				return string(common)
			}
		}

		common = append(common, word)
	}
	return string(common)
}

func main() {
	inputs := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
	}

	for _, input := range inputs {
		fmt.Printf("the longest-common-prefix of %v is %v \n", input, longestCommonPrefix(input))
	}

}
