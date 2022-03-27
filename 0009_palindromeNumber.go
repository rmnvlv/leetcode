package main

import "fmt"

/*
---------------------------------------------------------EASY---------------------------------------------------------
Given an integer x, return true if x is palindrome integer.
An integer is a palindrome when it reads the same backward as forward.

    For example, 121 is a palindrome while 123 is not.


Example 1:
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Constraints:
    -231 <= x <= 231 - 1
*/

// First try Runtime: 24 ms, Memory Usage: 6.3 MB
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if 0 <= x && x < 10 {
		return true
	}

	var end int
	slice := make([]int, 0)
	for {
		end = x % 10
		slice = append(slice, end)
		x /= 10
		if x == 0 {
			break
		}
	}

	count := -1
	halfOfLen := len(slice) / 2

	for i := range slice {
		if slice[i] == slice[len(slice)-1-i] {
			count++
		}
		if i == halfOfLen {
			break
		}
	}
	if count == halfOfLen {
		return true
	}
	return false
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(isPalindrome(x))
}
