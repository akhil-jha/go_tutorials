/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]



Constraints:

    2 <= nums.length <= 104
    -109 <= nums[i] <= 109
    -109 <= target <= 109
    Only one valid answer exists.

*/

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	visited := make(map[int]int)

	for key, value := range nums {
		compliment := target - value
		visited[value] = key // because we want to return the key and track all the values visited.

		compliment_index, exists := visited[compliment]

		// if visited[compliment] exists, it will return the value and bool(True in this case),
		// else 0 and False.
		// 0 becasue in Go 0 is return if nothing is found.
		fmt.Println(visited)
		// map[11:0]
		// map[11:0 15:1]
		// map[7:2 11:0 15:1]
		// map[2:3 7:2 11:0 15:1]
		// [3 2]
		if exists {
			return []int{key, compliment_index} //
		}
	}
	return nil
}

func main() {
	nums := []int{11, 15, 7, 2}
	target := 9
	fmt.Println(twoSum(nums, target))
}
