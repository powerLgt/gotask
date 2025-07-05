package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i := range str {
		if len(str)/2 < i {
			return true
		}

		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

/**
 *  给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
 *  回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *  例如，121 是回文，而 123 不是。
 **/
func main() {
	inputs := []int{121, 123, -121, 10, 1}
	for _, input := range inputs {
		fmt.Printf("%d isPalindrome: %v \n", input, isPalindrome(input))
	}
}
