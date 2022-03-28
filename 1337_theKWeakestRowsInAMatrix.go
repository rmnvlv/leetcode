package main

import "sort"

/*
---------------------------------------------------------EASY---------------------------------------------------------
You are given an m x n binary matrix mat of 1's (representing soldiers) and 0's (representing civilians).
The soldiers are positioned in front of the civilians.
That is, all the 1's will appear to the left of all the 0's in each row.

A row i is weaker than a row j if one of the following is true:

The number of soldiers in row i is less than the number of soldiers in row j.
Both rows have the same number of soldiers and i < j.
Return the indices of the k weakest rows in the matrix ordered from weakest to strongest.

Example 1:
Input: mat =
[[1,1,0,0,0],
 [1,1,1,1,0],
 [1,0,0,0,0],
 [1,1,0,0,0],
 [1,1,1,1,1]],
k = 3
Output: [2,0,3]
Explanation:
The number of soldiers in each row is:
- Row 0: 2
- Row 1: 4
- Row 2: 1
- Row 3: 2
- Row 4: 5
The rows ordered from weakest to strongest are [2,0,3,1,4].

Example 2:
Input: mat =
[[1,0,0,0],
 [1,1,1,1],
 [1,0,0,0],
 [1,0,0,0]],
k = 2
Output: [0,2]
Explanation:
The number of soldiers in each row is:
- Row 0: 1
- Row 1: 4
- Row 2: 1
- Row 3: 1
The rows ordered from weakest to strongest are [0,2,3,1].

Constraints:
m == mat.length
n == mat[i].length
2 <= n, m <= 100
1 <= k <= m
matrix[i][j] is either 0 or 1.
*/

//First try Runtime: 4 ms, Memory Usage: 5 MB
type MyMap struct {
	Key   int
	Value int
}

type MyMapList []MyMap

//Funcs for quik sort
func (m MyMapList) Len() int { return len(m) }

func (m MyMapList) Less(i, j int) bool {
	if m[i].Value < m[j].Value {
		return true
	}
	if m[i].Value > m[j].Value {
		return false
	}
	return m[i].Key < m[j].Key
}

func (m MyMapList) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func kWeakestRows(mat [][]int, k int) []int {
	mapMat := make(map[int]int, len(mat))

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				mapMat[i] = j
				break
			}
			if mat[i][len(mat[i])-1] == 1 {
				mapMat[i] = len(mat[i]) + 1
				break
			}
		}
	}

	var m MyMapList
	for k, v := range mapMat {
		m = append(m, MyMap{
			Key:   k,
			Value: v,
		})
	}

	//quik sort
	sort.Sort(m)

	n := []int{}
	for i := 0; i < k; i++ {
		n = append(n, m[i].Value)
	}

	return n
}

func main() {

}
