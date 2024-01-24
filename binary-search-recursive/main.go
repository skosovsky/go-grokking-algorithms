package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	foundItem := binarySearch(nums, 7)

	fmt.Println(foundItem)
}

func binarySearch(nums []int, item int) int {
	if len(nums) == 0 {
		return 0
	}

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0
}

//func binarySearchRecursive(nums []int, item int) {
//
//}
