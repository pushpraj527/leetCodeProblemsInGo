/*
1. https://leetcode.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i, x := range nums {
		y := target - x
		ri, e := mp[y]
		if !e {
			mp[x] = i
		} else {
			return []int{ri, i}
		}
	}
	return []int{}
}
