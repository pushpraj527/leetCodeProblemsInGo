/*
53. https://leetcode.com/problems/maximum-subarray/
*/
func maxSubArray(nums []int) int {
	max, cur := nums[0], 0
	for _, x := range nums {
		if cur < 0 {
			cur = 0
		}
		cur += x
		if cur > max {
			max = cur
		}
	}
	return max
}
