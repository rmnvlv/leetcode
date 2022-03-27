package main

import "fmt"

/*
---------------------------------------------------------EASY---------------------------------------------------------
You are given a string s consisting of lowercase English letters.
A duplicate removal consists of choosing two adjacent and equal letters and removing them.
We repeatedly make duplicate removals on s until we no longer can.
Return the final string after all such duplicate removals have been made.
It can be proven that the answer is unique.

Example 1:
Input: s = "abbaca"
Output: "ca"
Explanation:
For example, in "abbaca" we could remove "bb"
since the letters are adjacent and equal, and this is
the only possible move.  The result of this move is that the
string is "aaca", of which only "aa" is possible, so the final
string is "ca".

Example 2:
Input: s = "azxxzy"
Output: "ay"

Constraints:
1 <= s.length <= 105
s consists of lowercase English letters.
*/

//First try Runtime: 3 ms, Memory Usage: 6.6 MB
func removeDuplicates(s string) string {
	myStr := make([]byte, 0, len(s))

	for _, b := range []byte(s) {
		if len(myStr) > 0 {
			fmt.Println(string(b), " AND ", string(myStr[len(myStr)-1]))
		}
		if len(myStr) > 0 && b == myStr[len(myStr)-1] {
			myStr = myStr[:len(myStr)-1]
			continue
		}
		myStr = append(myStr, b)
		fmt.Println("Appended on: ", string(myStr))
	}
	return string(myStr)
}

func main() {
	fmt.Println(removeDuplicates("weewrtyyopw"))
}
