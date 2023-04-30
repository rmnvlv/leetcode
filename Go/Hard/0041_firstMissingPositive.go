package Hard

import "sort"

/*
Given an unsorted integer array nums, return the smallest missing positive integer.
You must implement an algorithm that runs in O(n) time and uses constant extra space.

Example 1:
Input: nums = [1,2,0]
Output: 3

Example 2:
Input: nums = [3,4,-1,1]
Output: 2

Example 3:
Input: nums = [7,8,9,11,12]
Output: 1

Constraints:
1 <= nums.length <= 5 * 10^5
-2^31 <= nums[i] <= 2^31 - 1
*/

// Thert solution Runtime: 157 ms, Memory Usage: 25.5 MB,
func firstMissingPositive3(nums []int) int {
	sort.Ints(nums)

	answer := 1

	for _, num := range nums {
		if num == answer {
			answer += 1
		} else if num > answer {
			break
		}
	}

	return answer
}

// Second solution Runtime: 425 ms, Memory Usage: 62.5 MB,
func firstMissingPositive2(nums []int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] > len(nums) {
			nums[i] = len(nums) + 1
		}
	}

	sort.Ints(nums)

	max := nums[len(nums)-1]

	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]] = i
	}

	for i := 1; i < max; i++ {
		if _, ok := hashmap[i]; !ok {
			return i
		}
	}

	return nums[len(nums)-1] + 1

}

// First solution Runtime: 251 ms, Memory Usage: 62.9 MB,
func firstMissingPositive(nums []int) int {

	sort.Ints(nums)

	max := nums[len(nums)-1]

	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]] = i
	}

	for i := 1; i < max; i++ {
		if _, ok := hashmap[i]; !ok {
			return i
		}
	}

	if len(nums) == 1 && nums[0] < 0 {
		return 1
	}

	if len(nums) < 3 && len(nums) > 1 && nums[1] < 0 {
		return 1
	}

	return nums[len(nums)-1] + 1

}

func main() {

}
