package main

/*
---------------------------------------------------------EASY---------------------------------------------------------
There are n houses evenly lined up on the street, and each house is beautifully painted.
You are given a 0-indexed integer array colors of length n, where colors[i]
represents the color of the ith house.
Return the maximum distance between two houses with different colors.
The distance between the ith and jth houses is abs(i - j), where abs(x) is the absolute value of x.

Example 1:
Input: colors = [1,1,1,6,1,1,1]
Output: 3
Explanation: In the above image, color 1 is blue, and color 6 is red.
The furthest two houses with different colors are house 0 and house 3.
House 0 has color 1, and house 3 has color 6. The distance between them is abs(0 - 3) = 3.
Note that houses 3 and 6 can also produce the optimal answer.

Example 2:
Input: colors = [1,8,3,8,3]
Output: 4
Explanation: In the above image, color 1 is blue, color 8 is yellow, and color 3 is green.
The furthest two houses with different colors are house 0 and house 4.
House 0 has color 1, and house 4 has color 3. The distance between them is abs(0 - 4) = 4.

Example 3:
Input: colors = [0,1]
Output: 1
Explanation: The furthest two houses with different colors are house 0 and house 1.
House 0 has color 0, and house 1 has color 1. The distance between them is abs(0 - 1) = 1.

Constraints:
n == colors.length
2 <= n <= 100
0 <= colors[i] <= 100
Test data are generated such that at least two houses have different colors.
*/

//First try Runtime: 4 ms, Memory Usage: 2.3 MB buble sort
func maxDistance(colors []int) int {
	max := 0
	for i := 0; i < len(colors)-1; i++ {
		for j := i + 1; j < len(colors); j++ {
			if colors[i] != colors[j] && max < j-i {
				max = j - i
			}
		}
	}
	return max
}

//Second try Runtime: 0 ms,, Memory Usage: 2.2 MB
func maxDistance2(colors []int) int {
	n1, n2 := 0, 0
	i, j := 0, len(colors)-1

	for i < j {
		if colors[i] != colors[j] {
			n1 = j - i
			break
		} else {
			i++
		}
	}

	i, j = 0, len(colors)-1

	for i < j {
		if colors[i] != colors[j] {
			n2 = j - i
			break
		} else {
			j--
		}
	}

	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func main() {

}
