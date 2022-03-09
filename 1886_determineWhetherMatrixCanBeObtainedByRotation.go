package main

import "fmt"

/*
---------------------------------------------------------EASY---------------------------------------------------------
Given two n x n binary matrices mat and target,
return true if it is possible to make mat equal to target
by rotating mat in 90-degree increments, or false otherwise.

Example 1:
Input: mat = [[0,1],[1,0]], target = [[1,0],[0,1]]
Output: true
Explanation: We can rotate mat 90 degrees clockwise to make mat equal target.

Example 2:
Input: mat = [[0,1],[1,1]], target = [[1,0],[0,1]]
Output: false
Explanation: It is impossible to make mat equal to target by rotating mat.

Example 3:
Input: mat = [[0,0,0],[0,1,0],[1,1,1]], target = [[1,1,1],[0,1,0],[0,0,0]]
Output: true
Explanation: We can rotate mat 90 degrees clockwise two times to make mat equal target.

Constraints:
n == mat.length == target.length
n == mat[i].length == target[i].length
1 <= n <= 10
mat[i][j] and target[i][j] are either 0 or 1.
*/

//First try Runtime: 2 ms, Memory Usage: 2.4 MB
func findRotation(mat [][]int, target [][]int) bool {
	f1, f2, f3, f4 := true, true, true, true
	l := len(mat)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			if mat[i][j] != target[i][j] {
				f1 = false
			}
			if mat[i][j] != target[l-1-j][i] {
				f2 = false
			}
			if mat[i][j] != target[j][l-1-i] {
				f3 = false
			}
			if mat[i][j] != target[l-i-1][l-j-1] {
				f4 = false
			}
		}
	}
	return f1 || f2 || f3 || f4
}

func main() {
	matrix := [][]int{{0, 1}, {1, 0}}
	target := [][]int{{1, 0}, {0, 1}}
	fmt.Println(findRotation(matrix, target))

}
