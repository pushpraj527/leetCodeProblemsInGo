/*
350. https://leetcode.com/problems/intersection-of-two-arrays-ii
 */

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var result []int
	for ptr1, ptr2 := 0, 0; ptr1 < len(nums1) && ptr2 < len(nums2); {
		if nums1[ptr1] == nums2[ptr2] {
			result = append(result, nums1[ptr1])
			ptr1++
			ptr2++
		} else if nums1[ptr1] < nums2[ptr2] {
			ptr1++
		} else {
			ptr2++
		}
	}
	return result

}
