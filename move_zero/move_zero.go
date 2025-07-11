/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:

Input: nums = [0]
Output: [0]



Constraints:

    1 <= nums.length <= 104
    -231 <= nums[i] <= 231 - 1


Follow up: Could you minimize the total number of operations done?
*/

package main

import "fmt"

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
		if nums[i] == 0 {
			continue // Do nothing. Writing here just to be explicit
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}

}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
