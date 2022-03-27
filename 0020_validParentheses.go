package main

import "fmt"

/*
---------------------------------------------------------EASY---------------------------------------------------------
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Constraints:
1 <= s.length <= 104
s consists of parentheses only '()[]{}'
*/

//First try Runtime: 0 ms, Memory Usage: 1.8 MB
func isValid(s string) bool {

	length := len(s)
	if length%2 != 0 || length < 2 {
		return false
	}

	stack := make([]byte, 0, length)

	for i := 0; i < length; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && s[i] == ')' && stack[len(stack)-1] == '(' ||
			len(stack) > 0 && s[i] == ']' && stack[len(stack)-1] == '[' ||
			len(stack) > 0 && s[i] == '}' && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}

	}

	return len(stack) == 0
}

func main() {
	s := "()"

	fmt.Println(isValid(s))
}
