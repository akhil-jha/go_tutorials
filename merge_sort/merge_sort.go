package main

import (
	"fmt"
)

func MergeSort(leftArray []int, rightArray []int, Array []int) {
	leftSize := len(Array) / 2
	rightSize := len(Array) - leftSize

	i := 0     //array counter
	left := 0  //counter for left array
	right := 0 //counter for right array

	for left < leftSize && right < rightSize {
		if leftArray[left] < rightArray[right] {
			Array[i] = leftArray[left]
			left++
		} else {
			Array[i] = rightArray[right]
			right++
		}
		i++
	}

	for left < leftSize {
		Array[i] = leftArray[left]
		left++
		i++
	}

	for right < rightSize {
		Array[i] = rightArray[right]
		right++
		i++
	}

}

func SortMe(numArray []int) []int {
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

	leftArray = SortMe(leftArray)
	rightArray = SortMe(rightArray)
	MergeSort(leftArray, rightArray, numArray)

	return numArray
}

func main() {
	numArray := []int{234, 89, 19, 1000, 2873, 1773, 1, 0, 990}
	fmt.Println(SortMe(numArray))
}
