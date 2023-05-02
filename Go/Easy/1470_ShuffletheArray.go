package Easy

//Runtime 3 ms Beats 87.89% Memory 4 MB Beats 5.92%
func shuffle1(nums []int, n int) []int {
	ans := []int{}

	for i := 0; i < n; i++ {
		ans = append(ans, nums[i], nums[i+n])
	}

	return ans
}

//Same - Two pointers
func shuffle2(nums []int, n int) []int {
	ans := make([]int, n*2)

	for i, j := 0, 0; i < n; j, i = j+2, i+1 {
		ans[j], ans[j+1] = nums[i], nums[i+n]
	}

	return ans
}

// Runtime 3 ms Beats 88.2% Memory 3.5 MB Beats 99.72%
//Too hard
func shuffle3(nums []int, n int) []int {
	mask := (1 << 10) - 1

	for i := 0; i < n; i++ {
		x1, x2 := nums[i], nums[i+n]
		tmpX1, tmpX2 := (x1&mask)<<10, (x2&mask)<<10
		nums[i*2] |= tmpX1
		nums[i*2+1] |= tmpX2
	}

	for i := 0; i < n*2; i++ {
		nums[i] >>= 10
	}

	return nums
}

/*
 nums = [2,3,5,4,1,7], n = 3 --> [2 4 3 1 5 7]
 i=0
 x1 = nums[0] = 2, x2 = nums[0+3] = 4
 tmpX1 = (2&1023)<<10 = 2048, tmpX2 = (4&1023)<<10 = 4096
 nums[0*2] = 2 | 2048 = 2050, nums[0*2+1]= 3 | 4096 = 4099
 nums = [2050,4099,5,4,1,7]

 i=1
 x1 = nums[1] = 4099, x2 = nums[1+3] = 1
 tmpX1 = (4099&1023)<<10 = 3072, tmpX2 = (1&1023)<<10 = 1024
 nums[1*2] = 5 | 3072 = 3077, nums[1*2+1]= 4 | 1024 = 1028
 nums = [2050,3072,3077,1028,1,7]

 i=2
 x1 = nums[2] = 3077, x2 = nums[2+3] = 7
 tmpX1 = (3077&1023)<<10 = 5120, tmpX2 = (7&1023)<<10 = 7168
 nums[2*2] = 1 | 5120 = 5121, nums[2*2+1]= 7 | 7168 = 7175
 nums = [2050,3072,3077,1028,5121,7165]
*/
