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

//First try

func findRotation(mat [][]int, target [][]int) bool {
	rotate := make([][]int, 0, len(mat))
	rotate = mat
	t := 0
	flag := true

	fmt.Println(len(mat))

	fmt.Println(rotate, "Main value ________________")

	for t < 4 {
		rotate, t = rot(t, rotate, mat)
		mat = rotate
		flag = compare(rotate, target)
		if flag == true {
			return true
		}
	}

	return false
}

func compare(rotate [][]int, target [][]int) bool {
	fmt.Println(rotate, target)
	for i := 0; i < len(rotate); i++ {
		for j := 0; j < len(rotate); j++ {
			if rotate[i][j] != target[i][j] {
				return false
			}
		}
	}
	return true
}

func rot(t int, rotate [][]int, mat [][]int) ([][]int, int) {
	fmt.Println("---------------------START ROTATION--------------------------")
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			fmt.Println(rotate, " Cofig into this -> ", mat)
			rotate[j][len(mat)-i-1] = mat[i][j]
			fmt.Println("After: ....", rotate, mat)
		}
	}

	t++
	return rotate, t
}

func main() {
	matrix := [][]int{{0, 1}, {1, 0}}
	target := [][]int{{1, 0}, {0, 1}}
	fmt.Println(findRotation(matrix, target))

}
