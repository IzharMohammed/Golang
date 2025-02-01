package main

import (
	"fmt"
	"slices"
)

func main() {
	// uninitialized slices
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// slices with make
	// var nums = make([]int, 0, 5)
	// // capacity -> maximum number of elements that can fit
	// nums = append(nums, 1, 2, 3, 4, 5, 6)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))
	// fmt.Println(nums)

	// slices with shorthand
	// nums := []int{}
	// nums = append(nums, 1, 2, 3, 4, 5, 6)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))
	// fmt.Println(nums)

	// copying slices
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 1, 2, 3)
	// var nums2 = make([]int, len(nums))
	// // copy function
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// slice operator
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(nums[2:8])
	// fmt.Println(nums[:8])
	// fmt.Println(nums[4:])

	// slices
	var nums1 = []int{1, 2, 3}
	var nums2 = []int{1, 2, 4}
	fmt.Println(slices.Equal(nums1, nums2))

	// 2D slices
	nums3 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(nums3)

}
