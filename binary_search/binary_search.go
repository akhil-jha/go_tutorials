package main

import "fmt"

func bsearch(nums []int, left int, right int, value int) int {
	if left > right {
		return -1 // Not found
	}

	mid := (left + right) / 2

	if nums[mid] == value {
		return mid        // Value found
	} else if nums[mid] < value {
		return bsearch(nums, mid+1, right, value) // Value less than nums[mid]
	} else {
		return bsearch(nums, left, mid-1, value) // Value more than nums[mid]
	}
}

func main() {
	nums := []int{5, 6, 8, 19, 55, 88, 111, 556, 789}
	fmt.Println(bsearch(nums, 0, len(nums)-1, 56))
	fmt.Println(bsearch(nums, 0, len(nums)-1, 556))
}
