/*
88. https://leetcode.com/problems/merge-sorted-array
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	var tmp []int
	for _, x := range nums1 {
		tmp = append(tmp, x)
	}

	ptr1, ptr2 := 0, 0
	for i := 0; i < m+n; i++ {
		if ptr1 < m && ptr2 < n {
			if tmp[ptr1] < nums2[ptr2] {
				nums1[i] = tmp[ptr1]
				ptr1++
			} else {
				nums1[i] = nums2[ptr2]
				ptr2++
			}
		} else if ptr1 < m && ptr2 >= n {
			nums1[i] = tmp[ptr1]
			ptr1++
		} else if ptr1 >= m && ptr2 < n {
			nums1[i] = nums2[ptr2]
			ptr2++
		}
	}
}
