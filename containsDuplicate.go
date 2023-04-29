/*
217. https://leetcode.com/problems/contains-duplicate/
*/

func containsDuplicate(nums []int) bool {
mp := make(map[int]bool)
for  _, x := range nums {
if mp[x] {
return true
} else {
mp[x] = true
}
}

return false
}