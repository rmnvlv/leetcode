package Easy

import "fmt"

/*
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.



Constraints:
    1 <= strs.length <= 200
    0 <= strs[i].length <= 200
    strs[i] consists of only lower-case English letters.
*/

// First try Runtime: 7 ms, Memory Usage: 2.4 MB
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs[0] == "" {
		return ""
	}
	maxLenWord := 201
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < maxLenWord {
			maxLenWord = len(strs[i])
		}
	}
	letter := string(strs[0][0])
	answer := ""
	for k := 0; k < maxLenWord; k++ {
		var flag bool
		var i int
		var target int
		letter = string(strs[target][k])
		for i = 0; i < len(strs); i++ {
			if letter == string(strs[i][k]) {
				flag = true
				target = i
			} else {
				return answer
			}
		}
		if flag {
			answer += string(strs[target][k])
			fmt.Println(answer)
		}
	}
	return answer
}
