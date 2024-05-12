package Medium

import "sort"

//First try Runtime 34ms
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var triplets [][]int
	for i := 0; i < len(nums)-2; i++ {

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				triplets = append(triplets, []int{nums[i], nums[l], nums[r]})

				//Avoid duplets
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}

		//for i duplets too
		for i < len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return triplets
}
