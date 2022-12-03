package main

import (
	"LeetCodeInGo/easy"
)

func main() {
	//fmt.Println(easy.IsPalindrome(121))

	//nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	//fmt.Println(easy.MissingNumber(nums))

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	easy.Merge(nums1, 3, nums2, 3)
}
