/*
Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist, return the maximum number.
Example 1:

Input: nums = [3,2,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2.
The third distinct maximum is 1.

Example 2:

Input: nums = [1,2]
Output: 2
Explanation:
The first distinct maximum is 2.
The second distinct maximum is 1.
The third distinct maximum does not exist, so the maximum (2) is returned instead.

Example 3:

Input: nums = [2,2,3,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2 (both 2's are counted together since they have the same value).
The third distinct maximum is 1.

Constraints:

	1 <= nums.length <= 104
	-231 <= nums[i] <= 231 - 1
*/
package main

import "fmt"

func MergeSort(leftArray []int, rightArray []int, numArray []int) {
	mid := len(numArray) / 2
	leftSize := mid
	rightSize := len(numArray) - leftSize

	i := 0
	left := 0
	right := 0

	for left < leftSize && right < rightSize {
		if leftArray[left] > rightArray[right] {
			numArray[i] = leftArray[left]
			left++
		} else {
			numArray[i] = rightArray[right]
			right++
		}
		i++
	}

	for left < leftSize {
		numArray[i] = leftArray[left]
		left++
		i++
	}

	for right < rightSize {
		numArray[i] = rightArray[right]
		right++
		i++
	}

}

func Sortme(numArray []int) []int {
	if len(numArray) <= 1 {
		return numArray
	}
	mid := len(numArray) / 2
	leftArray := make([]int, mid)
	rightArray := make([]int, len(numArray)-mid)

	for i := 0; i < len(numArray); i++ {
		if i < mid {
			leftArray[i] = numArray[i]
		} else {
			rightArray[i-mid] = numArray[i]
		}
	}
	rightArray = Sortme(rightArray)
	leftArray = Sortme(leftArray)
	MergeSort(leftArray, rightArray, numArray)

	return numArray
}

func thirdMax(nums []int) int {
	nums = Sortme(nums)
	temp := make([]int, 0, len(nums))
	temp = append(temp, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			temp = append(temp, nums[i])
		}
	}
	if len(temp) < 3 {
		return temp[0]
	}
	return temp[2]
}

func main() {
	nums := []int{2, 2, 3, 1, 5, 8}
	fmt.Println(thirdMax(nums))
}
