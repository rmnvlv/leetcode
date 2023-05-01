package Medium

//Runtime 9 ms Beats 74.85% Memory 6.6 MB Beats 54.49%
func minPartitions1(n string) int {
	max := 0
	for _, digit := range n {
		if max < int(digit-'0') {
			max = int(digit - '0')
		}
	}
	return max
}

//Runtime 4 ms Beats 94.61% Memory 6.4 MB Beats 95.81%
func minPartitions2(n string) int {
	var max byte = '0'
	for i := 0; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
		}
	}
	return int(max - '0')
}
